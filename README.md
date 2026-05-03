# golang-learning

Personal Go study repository — targeting job-ready, production-quality Go code.

**Background:** Java · Python · JS/Node  
**Go version:** 1.26.1 (1.22+ features available)  
**Module:** `github.com/rohitdhiman/golang-learning`

---

## Phases

| Phase | Topic | Status |
|---|---|---|
| 1 | Foundations — syntax, control flow, functions, slices/maps | 🟡 In Progress |
| 2 | Type System — structs, interfaces, pointers, generics, errors | ⬜ Not started |
| 3 | Concurrency — goroutines, channels, sync, patterns | ⬜ Not started |
| 4 | Standard Library — I/O, HTTP, JSON, testing, context | ⬜ Not started |
| 5 | Production — project structure, databases, CLI, REST API | ⬜ Not started |
| 6 | Advanced & Interview — performance, internals, interview questions | ⬜ Not started |

---

## Repo Layout

```
golang-learning/
├── docs/           # roadmap.md · next-session.md · gaps.md
├── concepts/       # one markdown file per topic (theory + Java/Python comparisons)
├── exercises/      # hands-on coding tasks, one folder per phase
├── cheatsheets/    # condensed quick-reference cards
├── projects/       # capstone: cli-tool/ · rest-api/
└── interview/      # questions.md · code-review.md
```

## Session Workflow

1. Open `docs/roadmap.md` — find the current week
2. Read `concepts/<topic>.md`
3. Do exercises in `exercises/<phase>/`
4. Fill cheatsheet at end of phase
5. Log gaps in `docs/gaps.md`
6. Update `docs/next-session.md`

## Quick Commands

```bash
go run main.go          # run a file
go test ./...           # run all tests
go test -race ./...     # run with race detector
go vet ./...            # static analysis
go mod tidy             # sync dependencies
gofmt -w .              # format all files
```