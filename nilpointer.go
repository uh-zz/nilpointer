package nilpointer

import (
	"go/types"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "nilpointer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "nilpointer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		commentmap.Analyzer,
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs

	reportFail := func(v ssa.Value, ret *ssa.Return, format string) {
		pos := ret.Pos()
		line := getNodeLineNumber(pass, ret)
		if !cmaps.IgnoreLine(pass.Fset, line, "nilpointer") {
			pass.Reportf(pos, format, line)
		}
	}

	for i := range funcs {
		for _, b := range funcs[i].Blocks {
			ret := getReturn(b)
			if ret == nil {
				continue
			}
			if ret := isReturnNil(ret); ret != nil {
				v := ret.Results[len(ret.Results)-1]
				reportFail(v, ret, "return nil pointer: line:%d")
			}
		}
	}
	return nil, nil
}

func getReturn(b *ssa.BasicBlock) *ssa.Return {
	if len(b.Instrs) == 0 {
		return nil
	}

	if ret, ok := b.Instrs[len(b.Instrs)-1].(*ssa.Return); ok {
		return ret
	}

	return nil
}

func isReturnNil(ret *ssa.Return) *ssa.Return {
	if len(ret.Results) == 1 {
		return nil
	}

	nilReturnValues := 0
	for _, res := range ret.Results {
		if _, ok := res.Type().(*types.Pointer); !ok {
			continue
		}

		v, ok := res.(*ssa.Const)
		if !ok {
			return nil
		}

		if v.IsNil() {
			nilReturnValues++
		}
	}

	if nilReturnValues == 0 {
		return nil
	}

	return ret
}

func getNodeLineNumber(pass *analysis.Pass, node ssa.Node) int {
	pos := node.Pos()
	return pass.Fset.File(pos).Line(pos)
}
