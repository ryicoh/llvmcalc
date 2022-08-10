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
	main := m.NewFunc("main", types.I64)
	entry := main.NewBlock("")

	res := codegen(entry, l.result)
	entry.NewRet(res)

	return m.String()
}

func codegen(entry *ir.Block, expression Expression) value.Value {
	switch expr := expression.(type) {
	case NumExpr:
		i, err := strconv.Atoi(expr.literal)
		if err != nil {
			panic(err)
		}

		a := entry.NewAlloca(types.I64)
		entry.NewStore(constant.NewInt(types.I64, int64(i)), a)
		return entry.NewLoad(types.I64, a)

	case BinOpExpr:
		left := codegen(entry, expr.left)
		right := codegen(entry, expr.right)
		switch expr.operator {
		case '+':
			return entry.NewAdd(left, right)
		case '-':
			return entry.NewSub(left, right)
		case '*':
			return entry.NewMul(left, right)
		case '/':
			return entry.NewUDiv(left, right)
		default:
			panic("???")
		}
	}
	panic("???")
}
