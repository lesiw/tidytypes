package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/tidytypes"
)

func main() { singlechecker.Main(tidytypes.Analyzer) }
