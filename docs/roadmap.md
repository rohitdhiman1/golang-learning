# Roadmap — Week by Week

Each session: **Read concept file → Do exercises → Update cheatsheet → Log gaps → Update next-session.md**

---

## Phase 1 — Foundations (Week 1–2)

### Week 1

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | Setup & Tooling | [`go-compilation-pipeline.md`](go-compilation-pipeline.md) | Install Go, VS Code + gopls, run `go tour` |
| 2 | Syntax Basics | [`concepts/syntax-basics.md`](../concepts/syntax-basics.md) | [`exercises/01-basics/`](../exercises/01-basics/) — Part 1: Variables |
| 3 | Control Flow | [`concepts/syntax-basics.md`](../concepts/syntax-basics.md) (control flow section) | [`exercises/01-basics/`](../exercises/01-basics/) — Part 2: Control Flow |
| 4 | Functions | [`concepts/functions.md`](../concepts/functions.md) | [`exercises/01-basics/`](../exercises/01-basics/) — Part 3: Functions |
| 5 | Review + Catch-up | — | Revisit anything from days 1–4 |

### Week 2

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | Arrays & Slices | [`concepts/slices-maps.md`](../concepts/slices-maps.md) | [`exercises/02-data-structures/`](../exercises/02-data-structures/) — Part 1: Slices |
| 2 | Maps | [`concepts/slices-maps.md`](../concepts/slices-maps.md) (maps section) | [`exercises/02-data-structures/`](../exercises/02-data-structures/) — Part 2: Maps |
| 3 | Structs | [`concepts/slices-maps.md`](../concepts/slices-maps.md) (structs section) | [`exercises/02-data-structures/`](../exercises/02-data-structures/) — Part 3: Structs |
| 4 | Review + cheatsheets | — | Fill [`cheatsheets/syntax.md`](../cheatsheets/syntax.md), [`cheatsheets/slices.md`](../cheatsheets/slices.md) |
| 5 | Phase 1 wrap-up | — | Log gaps, update [`next-session.md`](next-session.md) |

---

## Phase 2 — Type System & Composition (Week 3–4)

### Week 3

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | Structs & Methods (deep) | [`concepts/structs-methods.md`](../concepts/structs-methods.md) | [`exercises/03-interfaces/`](../exercises/03-interfaces/) — Part 1 |
| 2 | Pointers | [`concepts/pointers.md`](../concepts/pointers.md) | [`exercises/03-interfaces/`](../exercises/03-interfaces/) — Part 2 |
| 3 | Interfaces | [`concepts/interfaces.md`](../concepts/interfaces.md) | [`exercises/03-interfaces/`](../exercises/03-interfaces/) — Part 3 |
| 4 | Type Assertions & Switches | [`concepts/interfaces.md`](../concepts/interfaces.md) | [`exercises/03-interfaces/`](../exercises/03-interfaces/) — Part 4 |
| 5 | Review + Catch-up | — | — |

### Week 4

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | Error Handling | [`concepts/error-handling.md`](../concepts/error-handling.md) | [`exercises/03-interfaces/`](../exercises/03-interfaces/) — Part 5 |
| 2 | Generics | [`concepts/generics.md`](../concepts/generics.md) | [`exercises/03-interfaces/`](../exercises/03-interfaces/) — Part 6 |
| 3 | Review | — | — |
| 4 | Fill cheatsheets | — | [`cheatsheets/idioms.md`](../cheatsheets/idioms.md) (partial) |
| 5 | Phase 2 wrap-up | — | Log gaps, update [`next-session.md`](next-session.md) |

---

## Phase 3 — Concurrency (Week 5–6)

### Week 5

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | Goroutines | [`concepts/goroutines-channels.md`](../concepts/goroutines-channels.md) | [`exercises/04-concurrency/`](../exercises/04-concurrency/) — Part 1 |
| 2 | Channels (unbuffered) | [`concepts/goroutines-channels.md`](../concepts/goroutines-channels.md) | [`exercises/04-concurrency/`](../exercises/04-concurrency/) — Part 2 |
| 3 | Channels (buffered, select) | [`concepts/goroutines-channels.md`](../concepts/goroutines-channels.md) | [`exercises/04-concurrency/`](../exercises/04-concurrency/) — Part 3 |
| 4 | WaitGroup & Mutex | [`concepts/sync-primitives.md`](../concepts/sync-primitives.md) | [`exercises/04-concurrency/`](../exercises/04-concurrency/) — Part 4 |
| 5 | Review | — | — |

### Week 6

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | Concurrency Patterns | [`concepts/concurrency-patterns.md`](../concepts/concurrency-patterns.md) | [`exercises/04-concurrency/`](../exercises/04-concurrency/) — Part 5 |
| 2 | Context | [`concepts/context.md`](../concepts/context.md) | [`exercises/04-concurrency/`](../exercises/04-concurrency/) — Part 6 |
| 3 | Race Conditions & go vet | — | Run `go test -race` on previous exercises |
| 4 | Fill cheatsheets | — | [`cheatsheets/concurrency.md`](../cheatsheets/concurrency.md) |
| 5 | Phase 3 wrap-up | — | Log gaps, update [`next-session.md`](next-session.md) |

---

## Phase 4 — Standard Library (Week 7–8)

### Week 7

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | I/O (Reader/Writer) | [`concepts/stdlib-io.md`](../concepts/stdlib-io.md) | [`exercises/05-http/`](../exercises/05-http/) — Part 1 |
| 2 | Strings, Bytes, Runes | [`concepts/stdlib-io.md`](../concepts/stdlib-io.md) | [`exercises/05-http/`](../exercises/05-http/) — Part 2 |
| 3 | JSON encoding | [`concepts/stdlib-io.md`](../concepts/stdlib-io.md) | [`exercises/05-http/`](../exercises/05-http/) — Part 3 |
| 4 | Time & Duration | — | [`exercises/05-http/`](../exercises/05-http/) — Part 4 |
| 5 | Review | — | — |

### Week 8

| Day | Topic | Concept File | Exercises |
|---|---|---|---|
| 1 | HTTP server | [`concepts/stdlib-http.md`](../concepts/stdlib-http.md) | [`exercises/05-http/`](../exercises/05-http/) — Part 5 |
| 2 | HTTP client | [`concepts/stdlib-http.md`](../concepts/stdlib-http.md) | [`exercises/05-http/`](../exercises/05-http/) — Part 6 |
| 3 | Testing | [`concepts/testing.md`](../concepts/testing.md) | [`exercises/06-testing/`](../exercises/06-testing/) |
| 4 | Fill cheatsheets | — | [`cheatsheets/stdlib.md`](../cheatsheets/stdlib.md), [`cheatsheets/testing.md`](../cheatsheets/testing.md) |
| 5 | Phase 4 wrap-up | — | Log gaps, update [`next-session.md`](next-session.md) |

---

## Phase 5 — Production Patterns (Week 9–10)

| Topic | Concept File | Deliverable |
|---|---|---|
| Project layout | [`concepts/project-structure.md`](../concepts/project-structure.md) | Scaffold [`projects/cli-tool/`](../projects/cli-tool/) |
| Dependency management | [`concepts/project-structure.md`](../concepts/project-structure.md) | — |
| Configuration & logging | — | Add config/logging to cli-tool |
| Database (`database/sql`) | — | Add DB layer |
| CLI with cobra | — | Working CLI with subcommands |
| REST API | — | [`projects/rest-api/`](../projects/rest-api/) with chi router |
| Graceful shutdown | — | Apply to rest-api |

---

## Phase 6 — Advanced & Interview Prep (Week 11–12)

| Topic | File |
|---|---|
| Reflection & struct tags | [`concepts/performance.md`](../concepts/performance.md) |
| Profiling with pprof | [`concepts/performance.md`](../concepts/performance.md) |
| Build & deploy (cross-compile, Docker) | — |
| Interview Q&A | [`interview/questions.md`](../interview/questions.md) |
| Code review exercises | [`interview/code-review.md`](../interview/code-review.md) |
| Fill final cheatsheets | [`cheatsheets/idioms.md`](../cheatsheets/idioms.md) |

---

## Current Position

> **Current →** Phase 3 complete. Next: Phase 4, Week 7, Day 1
>
> See [`next-session.md`](next-session.md) for the immediate next task.
