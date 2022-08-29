# Analyzers

`mutility-analyzers` is a collection of the best analyzers in the mutility project.

## Use

Run from source with
`go run github.com/mutility/analyzers/cmd/mutility-analyzers@latest` or install
with `go install github.com/mutility/analyzers/cmd/mutility-analyzers@latest`
and run mutility-analyzers from GOPATH/bin.

Most of the included analyzers offer some level of configuration through flags.

## Available analyzers and commands

- [Unused Results](https://github.com/mutility/unused-results): report results that appear to always be ignored
- [Variable Format](https://github.com/mutility/variable-format): report non-const, non-literal strings passed to printf-like functions

## Backwards Compatibility

Previously mutility-analyzers included varfmt directly, so for now it still
includes stubs. Please update your references to use
[variable-format](https://github.com/mutility/variable-format) as the stubs will
be removed from here in the future.
