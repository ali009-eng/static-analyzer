# Go Static Analyzer

A simple static analysis tool for Go source code, built from scratch using Go's standard library (`go/parser`, `go/ast`, `go/token`) — no external dependencies.

It currently supports two checks:

- **Unchecked errors** — flags `err := someFunc()` assignments that aren't immediately followed by an `if err != nil` check
- **Cyclomatic complexity** — reports a complexity score per function, based on the number of branching points (`if`, `for`, `range`, `case`)
