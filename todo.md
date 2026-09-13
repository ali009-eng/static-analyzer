# TODO

## Done
- [x] Core AST parsing pipeline (go/parser, go/ast, go/token)
- [x] Unchecked-errors rule (checkUncheckedErrors)
- [x] Cyclomatic complexity rule (checkCyclomaticComplexity)
- [x] Multi-file / directory support (filepath.WalkDir)
- [x] Standardized file:line:col output format
- [x] --rule flag to selectively run individual checks
- [x] README

## Next up
- [ ] Refactor rules to RETURN results instead of printing directly
      (needed before tests are possible in a clean way)
- [ ] Write tests (go test ./...) covering:
  - [ ] unchecked error correctly flagged
  - [ ] properly checked error NOT flagged
  - [ ] wrong-variable check (e.g. `if result != nil`) still flags err as unchecked
  - [ ] complexity baseline (no branches) == 1
  - [ ] complexity with if/for/range/case increments correctly
  - [ ] multi-file directory scanning finds all .go files
  - [ ] --rule flag correctly filters which checks run

## Later / stretch goals
- [ ] Broaden unchecked-error detection to match by TYPE (error) not just variable name "err"
- [ ] Count && / || in cyclomatic complexity (true McCabe complexity)
- [ ] Look further than just the NEXT statement for an err check
      (currently only checks stmts[i+1], misses later checks)
- [ ] Consider migrating to golang.org/x/tools/go/analysis framework
- [ ] Add unused-imports or os.Exit-outside-main as a third rule
- [ ] CI: run `go vet` +
