# Project — CLI Tool

> **Phase 5 · Week 9–10**  
> Build a non-trivial CLI application using `cobra`.

---

## Goal

Build a CLI tool from scratch that demonstrates production Go patterns:
- Multi-level subcommands
- Persistent flags (e.g., `--config`, `--verbose`)
- Structured logging with `log/slog`
- Configuration via env vars + flags
- Graceful error handling and exit codes

---

## Suggested Project: `gostat` — a local file statistics tool

```
gostat files count ./src          # count files by extension
gostat files largest ./src -n 10  # top 10 largest files
gostat lines ./src --ext .go      # total lines in .go files
gostat help
```

Feel free to choose your own domain — the important thing is having subcommands and real I/O.

---

## Minimum Requirements

- [ ] `cobra` for CLI framework (`go get github.com/spf13/cobra`)
- [ ] At least 2 subcommands
- [ ] Persistent `--verbose` flag at root level
- [ ] At least 1 config value from env var
- [ ] Uses `log/slog` for structured logging
- [ ] Returns non-zero exit code on error (`os.Exit(1)`)
- [ ] Has at least one table-driven test
- [ ] Code passes `go vet ./...` and `go test ./...`

---

## Getting Started

```bash
cd projects/cli-tool
go mod init github.com/rohitdhiman/golang-learning/projects/cli-tool
go get github.com/spf13/cobra@latest

# Suggested structure:
# cmd/
#   root.go     — root command, persistent flags
#   files.go    — 'files' subcommand
#   lines.go    — 'lines' subcommand
# internal/
#   scanner/    — file system logic
# main.go
```

---

_Placeholder — start when you reach Phase 5._
