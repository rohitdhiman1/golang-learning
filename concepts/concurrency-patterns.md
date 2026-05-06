# Concurrency Patterns

> **Phase 3 · Week 6**  
> Covers: fan-in, fan-out, worker pools, pipeline, context-based cancellation, errgroup, common pitfalls

---

## What It Is

Go's concurrency patterns are composable building blocks using goroutines and channels. These patterns solve real-world problems: processing streams of data, distributing work, and managing goroutine lifetimes.

### Pipeline

A pipeline is a series of stages connected by channels. Each stage is a group of goroutines that:
1. Receive values from an upstream channel
2. Perform some computation
3. Send results to a downstream channel

```go
// Stage 1: generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: square each number
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage: generate → square → consume
for v := range square(generate(2, 3, 4)) {
    fmt.Println(v)   // 4, 9, 16
}
```

### Fan-Out

Spread work from one channel across multiple goroutines to parallelise CPU-bound or I/O-bound work:

```go
// Multiple workers read from the same channel
for i := 0; i < numWorkers; i++ {
    go func() {
        for job := range jobs {
            process(job)
        }
    }()
}
```

### Fan-In (Merge)

Combine multiple channels into one:

```go
func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    merged := make(chan int)

    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                merged <- v
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(merged)
    }()

    return merged
}
```

### Worker Pool

A fixed number of goroutines processing jobs from a shared channel — the most common Go concurrency pattern:

```go
func workerPool(numWorkers int, jobs <-chan Job, results chan<- Result) {
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                results <- process(id, job)
            }
        }(i)
    }
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### errgroup (golang.org/x/sync/errgroup)

Run goroutines that can return errors, and collect the first error:

```go
g, ctx := errgroup.WithContext(context.Background())

for _, url := range urls {
    g.Go(func() error {
        return fetch(ctx, url)
    })
}

if err := g.Wait(); err != nil {
    log.Fatal(err)   // first error from any goroutine
}
```

`errgroup` handles WaitGroup + error collection + context cancellation in one package.

---

## Java/Python Comparison

| Pattern | Go | Java | Python |
|---|---|---|---|
| Pipeline | Channels chained through stages | `Stream.map().filter()` (lazy), `CompletableFuture.thenApply()` | Generator chains, `asyncio` pipelines |
| Worker pool | Goroutines + buffered channel | `ExecutorService.submit()` with thread pool | `concurrent.futures.ThreadPoolExecutor` |
| Fan-in/out | Explicit channel wiring | Fork/Join framework | `asyncio.gather()` |
| Error group | `errgroup.Group` | `CompletableFuture.allOf()` + exception handling | `asyncio.TaskGroup` (3.11+) |
| Cancellation | Context propagation | `Future.cancel()` | `asyncio.Task.cancel()` |

**Key difference**: In Go, you wire these patterns explicitly with channels. Java and Python use higher-level abstractions (executors, futures, async/await) that hide the plumbing.

---

## How It Works

### Pattern Selection Guide

```
Do I need to process a stream of data through stages?
  → Pipeline

Do I need to parallelise independent work?
  → Fan-out (multiple goroutines reading one channel)

Do I need to combine results from multiple sources?
  → Fan-in (merge channels)

Do I need a fixed number of workers processing jobs?
  → Worker Pool (fan-out + fan-in combined)

Do I need goroutines that can fail?
  → errgroup
```

### Goroutine Lifecycle Management

Every goroutine you start must have a clear exit path. The three common approaches:

1. **Done channel**: goroutine selects on a `done` channel
   ```go
   go func() {
       for {
           select {
           case <-done:
               return
           case v := <-in:
               process(v)
           }
       }
   }()
   ```

2. **Context cancellation**: pass `ctx` and check `ctx.Done()`
3. **Channel close**: goroutine ranges over input — exits when input is closed

---

## Key Gotchas

### 1. Goroutine leaks from unclosed channels
```go
func gen() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; ; i++ {
            ch <- i   // blocks forever if consumer stops reading
        }
    }()
    return ch
}
```
Fix: accept a `done` channel or context for cancellation.

### 2. Sending on a closed channel
Only the **sender** (or a coordinating goroutine) should close a channel. Closing from the receiver side causes panics if the sender is still sending.

### 3. Forgetting to close the results channel
In a worker pool, if you forget `close(results)`, any `range results` loop blocks forever.

### 4. Race conditions in fan-out
Multiple goroutines writing to the same channel is fine (channels are thread-safe). Multiple goroutines writing to the same slice/map without synchronisation is a race.

---

## Idiomatic Go

1. **Pipeline stages return `<-chan T`** — the stage owns the channel and closes it when done.
2. **The sender closes the channel** — never the receiver.
3. **Use `defer wg.Done()` as the first line in worker goroutines** — guarantees the WaitGroup decrements even on panic.
4. **Prefer `errgroup` over manual WaitGroup + error handling** — less boilerplate, built-in context cancellation.
5. **Always provide an exit path for goroutines** — context, done channel, or channel close.
6. **Size worker pools based on the workload** — CPU-bound: `runtime.NumCPU()`. I/O-bound: higher, determined by benchmarking.

---

## Exercises

See [`exercises/04-concurrency/`](../exercises/04-concurrency/README.md) — Part 5

---

## Related Files

- Previous: [`concepts/sync-primitives.md`](sync-primitives.md)
- Next: [`concepts/context.md`](context.md)
- Cheatsheet: [`cheatsheets/concurrency.md`](../cheatsheets/concurrency.md)
