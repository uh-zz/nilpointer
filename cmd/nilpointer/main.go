package main

import (
	"github.com/uh-zz/nilpointer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nilpointer.Analyzer) }
