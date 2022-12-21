package main

import (
	"nilpointer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nilpointer.Analyzer) }
