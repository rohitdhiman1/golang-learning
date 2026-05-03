# Go Compilation Pipeline

How Go turns source code into a running binary — and how it compares to Java.

---

## The Full Pipeline

```
┌─────────────────────────────────────────────────────┐
│                   YOU WRITE CODE                     │
│              main.go, utils.go, etc.                 │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  1. PARSE                                            │
│     Source code → Tokens → AST (Abstract Syntax Tree)│
│     • Catches syntax errors here                     │
│     • Each .go file becomes a tree of nodes          │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  2. TYPE CHECK                                       │
│     • Resolves all types, confirms assignments match │
│     • Checks interface satisfaction                  │
│     • Catches: wrong types, unused vars, missing     │
│       imports, undeclared names                       │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  3. SSA (Static Single Assignment) IR                │
│     • Converts AST → intermediate representation     │
│     • Each variable assigned exactly once             │
│     • Makes optimisation easier                      │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  4. OPTIMISE                                         │
│     • Escape analysis (stack vs heap allocation)     │
│     • Dead code elimination                          │
│     • Inline small functions                         │
│     • Bounds check elimination                       │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  5. MACHINE CODE GENERATION                          │
│     • SSA → native assembly for target platform      │
│     • Target set by GOOS + GOARCH                    │
│       (e.g. linux/amd64, darwin/arm64)               │
│     • Produces .o object files                       │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  6. LINK                                             │
│     • Combines all .o files + dependencies           │
│     • Statically links the Go runtime (GC, scheduler,│
│       goroutine stack management) INTO the binary    │
│     • No external runtime needed                     │
└──────────────────────┬──────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────┐
│  OUTPUT: Single static binary                        │
│     • Contains YOUR code + Go runtime + all deps     │
│     • Runs directly on the OS — no VM, no runtime    │
│     • Just copy and execute: ./myapp                 │
└─────────────────────────────────────────────────────┘
```

---

## Java vs Go — Side by Side

```
        JAVA                              GO
  ┌────────────┐                   ┌────────────┐
  │  .java     │                   │  .go       │
  └─────┬──────┘                   └─────┬──────┘
        ▼                                ▼
  ┌────────────┐                   ┌────────────┐
  │   javac    │                   │  go build  │
  │ (compiler) │                   │ (compiler  │
  └─────┬──────┘                   │  + linker) │
        ▼                          └─────┬──────┘
  ┌────────────┐                         ▼
  │  .class    │                   ┌────────────┐
  │ (bytecode) │                   │  binary    │
  └─────┬──────┘                   │ (native)   │
        ▼                          └─────┬──────┘
  ┌────────────┐                         ▼
  │    JVM     │                   Runs directly
  │ (interprets│                   on the OS.
  │  + JIT     │                   Nothing else
  │  compiles) │                   needed.
  └────────────┘
```

---

## Comparison Table

| | Java | Go |
|---|---|---|
| Need to install to **compile** | JDK (`javac`) | Go toolchain (`go`) |
| Need to install to **run** | JRE/JVM | **Nothing** — binary is self-contained |
| What the binary contains | Bytecode (instructions for JVM) | Native machine code + Go runtime |
| Cross-compile | Write once, run anywhere (via JVM) | `GOOS=linux GOARCH=amd64 go build` — build anywhere, run on target |

---

## `go run` vs `go build`

| Command | What it does | Use case |
|---|---|---|
| `go run main.go` | Compiles to a temp dir + runs immediately, discards binary | Development |
| `go build -o myapp .` | Compiles and saves the binary | Deployment |

Both go through the full pipeline above. `go run` just deletes the binary after execution.

---

## Cross-Compilation

From your Mac, build a Linux binary without a Linux machine:

```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux ./cmd/myapp
```

Common targets:

| GOOS | GOARCH | Platform |
|---|---|---|
| `darwin` | `arm64` | macOS (Apple Silicon) |
| `darwin` | `amd64` | macOS (Intel) |
| `linux` | `amd64` | Linux (x86_64) |
| `linux` | `arm64` | Linux (ARM — AWS Graviton, Raspberry Pi) |
| `windows` | `amd64` | Windows (x86_64) |
