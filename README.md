# Analyzers

## Use

Install with `go install github.com/mutility/analyzers/cmd/____@latest`,
replacing ____ with your chosen analyzer, such as `mutility-analyzers` or
`varfmt`. This should install the analyzer into your `GOPATH/bin`.

Then invoke like any other `golang.org/x/tools/go/analysis` checker:
`varfmt ./...`.

Currently `varfmt` is the only analyzer in this repository, so there is no
immediate advantage to using mutility-analyzers. However once others are added,
it will be an easier way to use them all.

## VarFmt

`varfmt` is an opinionated go analyzer intended as a diagnostic tool. When you
have a large project and something somewhere is writing messages with
`%!x(MISSING)` strewn through them, this analyzer can help you find the source.
It reports any format strings that are non-const, non-literal strings. Except
for ones used as pass through format strings in a wrapper.

### Example VarFmt messages

> ```errors
> ...: variable `os.Args[0]` used for fmt.Fprintf format parameter
> ```

### VarFmt usage

Optionally provide `-funcs` to specify some printf-like functions that were not
automatically detected.
