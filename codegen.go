package llvmcalc

import (
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func Codegen(l *Lexer) string {
	m := ir.NewModule()
	main := m.NewFunc("main", types.I8)
	entry := main.NewBlock("")

	res := codegen(entry, l.result)
	entry.NewRet(entry.NewFPToSI(res, types.I8))

	return m.String()
}

func codegen(entry *ir.Block, expression Expression) value.Value {
	switch expr := expression.(type) {
	case NumExpr:
		f, err := strconv.ParseFloat(expr.literal, 64)
		if err != nil {
			panic(err)
		}

		return constant.NewFloat(types.Double, float64(f))

	case BinOpExpr:
		left := codegen(entry, expr.left)
		right := codegen(entry, expr.right)
		switch expr.operator {
		case '+':
			return entry.NewFAdd(left, right)
		case '-':
			return entry.NewFSub(left, right)
		case '*':
			return entry.NewFMul(left, right)
		case '/':
			return entry.NewFDiv(left, right)
		default:
			panic("???")
		}
	}
	panic("???")
}
