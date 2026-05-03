# Go (Golang) — Learning Prompt

Personal study repository for **Go proficiency** — targeting job readiness and production-quality Go code.

---

## Goal

| Field | Detail |
|---|---|
| Language | Go (Golang) |
| Purpose | Work / job prep — write production Go confidently |
| Background | Java (statically typed, OOP), Python (scripting, dynamic), JS/Node (async, web) |
| Timeline | Open-ended, self-paced |
| Focus | General proficiency — not limited to web or CLI |

---

## What I Already Know (leverage this)

- **From Java:** static typing, interfaces, structs vs classes, compiled language, strong stdlib, error handling patterns, build tools
- **From Python:** scripting mindset, slices ≈ lists, maps ≈ dicts, first-class functions, clean syntax preference
- **From JS/Node:** async patterns (though Go's concurrency model is very different), JSON handling, HTTP servers

When explaining Go concepts, relate them to Java/Python equivalents where helpful — but flag where Go deliberately breaks from those patterns (e.g., no inheritance, no exceptions, no generics-heavy design).

---

## Topic Breakdown

### Phase 1 — Foundations (Week 1–2)
- **Setup & Tooling** — `go install`, `go mod`, `go run/build/test`, VS Code + gopls, `gofmt`/`goimports`
- **Syntax Basics** — variables, constants, types, zero values, type inference (`:=`), basic I/O
- **Control Flow** — if/else, for (Go's only loop), switch (no fallthrough by default), defer
- **Functions** — multiple returns, named returns, variadic, first-class functions, closures
- **Data Structures** — arrays, slices (length vs capacity, append, copy), maps, structs

### Phase 2 — Type System & Composition (Week 3–4)
- **Structs & Methods** — receiver functions (value vs pointer), constructor patterns, embedding (composition over inheritance)
- **Interfaces** — implicit satisfaction (duck typing), empty interface (`any`), type assertions, type switches
- **Pointers** — `&` and `*`, pointer receivers, nil pointers, no pointer arithmetic
- **Generics** — type parameters, constraints, `comparable`, when to use vs when not to
- **Error Handling** — `error` interface, `errors.New`, `fmt.Errorf` with `%w`, `errors.Is`/`errors.As`, sentinel errors, custom error types, why no exceptions

### Phase 3 — Concurrency (Week 5–6)
- **Goroutines** — lightweight threads, `go` keyword, scheduling model (GMP)
- **Channels** — unbuffered vs buffered, directional channels, `range` over channel, `select`
- **Sync Primitives** — `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Map`
- **Patterns** — fan-in, fan-out, worker pools, pipeline, context-based cancellation, errgroup
- **Common Pitfalls** — goroutine leaks, race conditions, `go vet -race`, channel deadlocks

### Phase 4 — Standard Library Deep Dive (Week 7–8)
- **I/O** — `io.Reader`/`io.Writer` interfaces, `bufio`, `os`, file handling
- **Strings & Bytes** — `strings`, `strconv`, `bytes`, `unicode/utf8`, runes vs bytes
- **Encoding** — `encoding/json` (struct tags, `Marshal`/`Unmarshal`, custom marshalers), `encoding/csv`, `encoding/xml`
- **Net/HTTP** — `http.ListenAndServe`, `http.Handler`/`HandlerFunc`, `ServeMux`, middleware pattern, `http.Client`, timeouts
- **Testing** — `testing` package, table-driven tests, `t.Run` subtests, benchmarks, `httptest`, test fixtures
- **Time** — `time.Time`, `time.Duration`, `time.Ticker`, `time.After`, reference time format (`Mon Jan 2 15:04:05 MST 2006`)
- **Context** — `context.Background()`, `context.WithCancel`, `context.WithTimeout`, `context.WithValue`, propagation rules

### Phase 5 — Project Structure & Production Patterns (Week 9–10)
- **Project Layout** — `cmd/`, `internal/`, `pkg/` conventions, when to split packages, circular dependency avoidance
- **Dependency Management** — `go.mod`, `go.sum`, versioning, `go get`, `go mod tidy`, private modules
- **Configuration** — env vars, flags (`flag` package), config files, 12-factor patterns
- **Logging** — `log/slog` (structured logging, Go 1.21+), log levels, JSON output
- **Database** — `database/sql`, connection pooling, prepared statements, `sqlx` (third-party), migrations
- **CLI Tools** — `cobra` (industry standard), subcommands, flags, arg parsing
- **API Development** — RESTful patterns, router libraries (chi, gorilla/mux, stdlib ServeMux in Go 1.22+), middleware chains, request validation, graceful shutdown

### Phase 6 — Advanced & Interview Prep (Week 11–12)
- **Reflection** — `reflect` package, when to use (almost never), struct tags
- **Code Generation** — `go generate`, `stringer`, mock generation
- **Performance** — profiling (`pprof`), benchmarks, escape analysis, memory allocation, `sync.Pool`
- **Build & Deploy** — cross-compilation (`GOOS`/`GOARCH`), static binaries, Docker multi-stage builds, `-ldflags` for version injection
- **Common Interview Topics** — slice internals (header, backing array, gotchas with append), map internals (hash buckets), interface internals (type + value), goroutine scheduling, garbage collector (tri-color mark & sweep), `init()` function ordering

---

## Repo Structure

```
golang-learning/
├── README.md                     # This file — project overview & guide
├── docs/
│   ├── roadmap.md                # Week-by-week schedule with file references
│   ├── next-session.md           # What to study in the next session
│   └── gaps.md                   # Knowledge gaps and weak areas tracker
├── concepts/                     # Core theory — one file per topic
│   ├── syntax-basics.md
│   ├── functions.md
│   ├── slices-maps.md
│   ├── structs-methods.md
│   ├── interfaces.md
│   ├── pointers.md
│   ├── generics.md
│   ├── error-handling.md
│   ├── goroutines-channels.md
│   ├── sync-primitives.md
│   ├── concurrency-patterns.md
│   ├── stdlib-io.md
│   ├── stdlib-http.md
│   ├── testing.md
│   ├── context.md
│   ├── project-structure.md
│   └── performance.md
├── exercises/                    # Hands-on coding exercises
│   ├── 01-basics/                # Variables, control flow, functions
│   ├── 02-data-structures/       # Slices, maps, structs
│   ├── 03-interfaces/            # Interface design, type assertions
│   ├── 04-concurrency/           # Goroutines, channels, worker pools
│   ├── 05-http/                  # HTTP server, client, middleware
│   ├── 06-testing/               # Table-driven tests, benchmarks, mocks
│   └── 07-project/               # Mini-project (CLI tool or REST API)
├── cheatsheets/                  # Quick-reference cards
│   ├── syntax.md                 # Go vs Java vs Python side-by-side
│   ├── slices.md                 # Slice operations, gotchas, internals
│   ├── concurrency.md            # Goroutine + channel patterns
│   ├── stdlib.md                 # Most-used stdlib packages
│   ├── testing.md                # Test patterns and commands
│   └── idioms.md                 # Go proverbs, naming conventions, style
├── projects/                     # Capstone projects (build from scratch)
│   ├── cli-tool/                 # Phase 5 — CLI app with cobra
│   └── rest-api/                 # Phase 5 — REST API with chi/stdlib
└── interview/                    # Interview prep
    ├── questions.md              # Common Go interview questions + answers
    └── code-review.md            # "What's wrong with this code?" exercises
```

### Folder purposes

| Folder | Purpose | When to use |
|---|---|---|
| `docs/` | Planning, tracking, progress | Open `roadmap.md` at the start of every session |
| `concepts/` | Core theory — structured, with Java/Python comparisons | Primary reading for each topic |
| `exercises/` | Small, focused coding tasks (30–60 min each) | After reading the concept file |
| `cheatsheets/` | Condensed reference cards | Fill after each phase; review before interviews |
| `projects/` | Build real things end-to-end | Phase 5+ capstone work |
| `interview/` | Interview-specific prep | Phase 6 and ongoing |

### Concept file template (every `concepts/*.md` follows this structure)

```
## What It Is         → brief definition
## Java/Python Comparison → how it maps to what you already know
## How It Works       → mechanics, syntax, key rules
## Key Gotchas        → subtle bugs, common mistakes
## Idiomatic Go       → the "Go way" to do it (not the Java/Python way)
## Exercises          → links to exercises/ folder
## Related Files      → links to cheatsheet, project usage
```

---

## How to Use This Repo (Session Workflow)

1. **Open `docs/roadmap.md`** — find the current phase/week. It tells you: Read → Code → Summarise.
2. **Read the concept file** — `concepts/<topic>.md` for the session's topic.
3. **Do the exercises** — `exercises/<phase>/` — write real Go code, run it, test it.
4. **Fill the cheatsheet** — `cheatsheets/<topic>.md` at end of each phase.
5. **Log gaps** — anything confusing goes in `docs/gaps.md` for targeted review.
6. **Update next session** — `docs/next-session.md` at end of each session.
7. **Build projects** — Phase 5+: apply everything in `projects/`.

---

## Study Resources

| Resource | Notes |
|---|---|
| [A Tour of Go](https://go.dev/tour/) | Official interactive tutorial — start here |
| [Effective Go](https://go.dev/doc/effective_go) | Official style and idioms guide — read after basics |
| [Go by Example](https://gobyexample.com) | Annotated code examples for every concept |
| [Go Standard Library Docs](https://pkg.go.dev/std) | Ground truth for any stdlib package |
| [Go Proverbs](https://go-proverbs.github.io/) | Rob Pike's design philosophy — internalise these |
| [100 Go Mistakes](https://100go.co/) | Book: common pitfalls (great for interview prep) |
| [Ardan Labs — Ultimate Go](https://www.ardanlabs.com/training/) | Deep-dive course (concurrency, performance) |
| [Go Wiki — Code Review Comments](https://go.dev/wiki/CodeReviewComments) | Style guide used by Go team |
| [Go Playground](https://go.dev/play/) | Browser-based Go sandbox for quick experiments |

---

## Progress Tracker

Update as phases are completed.

| Phase | Topic Focus | Status |
|---|---|---|
| 1 | Foundations — syntax, control flow, functions, slices/maps | Not started |
| 2 | Type System — structs, interfaces, pointers, generics, errors | Not started |
| 3 | Concurrency — goroutines, channels, sync, patterns | Not started |
| 4 | Standard Library — I/O, HTTP, JSON, testing, context | Not started |
| 5 | Production — project structure, databases, CLI, REST API | Not started |
| 6 | Advanced & Interview — performance, internals, interview questions | Not started |

---

## Key Go Mental Model Shifts (from Java/Python)

These are the things that will feel "wrong" at first. Lean into them — they're deliberate design choices.

| You're used to (Java/Python) | Go does this instead | Why |
|---|---|---|
| Classes + inheritance | Structs + composition (embedding) | Flat hierarchies, no fragile base class |
| `try/catch` exceptions | Multiple return values (`val, err`) | Errors are values, not control flow |
| Implements keyword | Implicit interface satisfaction | Decoupled packages, no import cycles |
| Thread pools + locks | Goroutines + channels | CSP model — "share memory by communicating" |
| Generics everywhere | Generics sparingly (added in 1.18) | Simplicity over abstraction |
| Package managers (pip, maven) | `go mod` (built-in, minimal) | No central registry needed |
| Getter/setter patterns | Exported fields (uppercase) | If it doesn't need protection, don't wrap it |
| Null/None | Zero values (every type has one) | No null pointer surprise — but still check `nil` for pointers/interfaces |

---

## Useful Commands

```bash
# Run a file
go run main.go

# Build a binary
go build -o myapp ./cmd/myapp

# Run all tests
go test ./...

# Run tests with race detector
go test -race ./...

# Run benchmarks
go test -bench=. ./...

# Format code (auto-applied on save in VS Code)
gofmt -w .

# Vet code for suspicious constructs
go vet ./...

# Download dependencies
go mod tidy

# View docs for a package
go doc fmt.Println

# Profile CPU
go test -cpuprofile cpu.prof -bench=. && go tool pprof cpu.prof
```
