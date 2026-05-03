# Standard Library — I/O, Strings, Encoding

> **Phase 4 · Week 7**  
> Covers: `io.Reader`/`io.Writer`, `bufio`, `os`, `strings`, `strconv`, `bytes`, `encoding/json`, `encoding/csv`

---

## What It Is

_Coming in Phase 4. Stub created — fill after completing Phase 3._

Topics to cover:
- `io.Reader` and `io.Writer` — the two most important interfaces in the stdlib
- Composing readers/writers: `io.TeeReader`, `io.MultiWriter`, `io.LimitReader`
- `bufio.Scanner` — line-by-line file/stdin reading
- `os.Open`, `os.Create`, `os.ReadFile`, `os.WriteFile`
- `strings` package — `Contains`, `HasPrefix`, `Split`, `Join`, `Builder`, `Replacer`
- `strconv` — `Atoi`, `Itoa`, `ParseFloat`, `FormatInt`
- Runes vs bytes — `len(s)` returns bytes, `utf8.RuneCountInString(s)` returns chars
- `encoding/json` — `Marshal`/`Unmarshal`, struct tags `json:"name,omitempty"`, custom marshalers
- `encoding/csv` — `csv.Reader`, `csv.Writer`

---

## Java/Python Comparison

_Fill when reading._

---

## How It Works

_Fill when reading._

---

## Key Gotchas

_Log gotchas here as you encounter them._

---

## Idiomatic Go

_Fill when reading._

---

## Exercises

See [`exercises/05-http/`](../exercises/05-http/README.md) — Parts 1–3

---

## Related Files

- Next: [`concepts/stdlib-http.md`](stdlib-http.md)
- Cheatsheet: [`cheatsheets/stdlib.md`](../cheatsheets/stdlib.md)
