# lesiw.io/tidytypes [![Go Reference](https://pkg.go.dev/badge/lesiw.io/tidytypes.svg)](https://pkg.go.dev/lesiw.io/tidytypes)

An `analysis.Analyzer` that reports redundant type declarations in function
parameters and results.

The analyzer identifies consecutive parameters or results with the same type
that can be grouped together using Go's type grouping syntax, making code more
concise and readable.

## Usage (test)

```go
package main

import (
	"testing"

	"lesiw.io/checker"
	"lesiw.io/tidytypes"
)

func TestCheck(t *testing.T) { checker.Run(t, tidytypes.Analyzer) }
```

## Usage (CLI)

```sh
go get -tool lesiw.io/tidytypes/cmd/tidytypes
go tool tidytypes ./...
```

## Features

- Detects redundant type declarations in function parameters and results.
- Works with both named and unnamed parameters/results
- Provides automatic fixes via suggested edits
- Supports function literals and method declarations
- Preserves existing grouped parameters that are already correct or cannot be
  reduced

## Fixes

- Consecutive parameters with the same type: `func(a int, b int)` →
  `func(a, b int)`
- Consecutive return values with the same type: `func() (int, int)` →
  `func() (_, _ int)`
- Multiple redundant groups: `func(a int, b int, c string, d string)` →
  `func(a, b int, c, d string)`
- Complex types: `func(a []int, b []int)` → `func(a, b []int)`
- Function literals: `var fn = func(x int, y int) {}` →
  `var fn = func(x, y int) {}`
