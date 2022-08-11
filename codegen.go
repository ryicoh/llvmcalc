package llvmcalc

import (
	"fmt"
	"strconv"
	"strings"

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
	if types.IsFloat(res.Type()) {
		res = entry.NewFPToSI(res, types.I8)
	} else {
		res = entry.NewTrunc(res, types.I8)
	}

	entry.NewRet(res)

	return m.String()
}

func codegen(entry *ir.Block, expression Expression) value.Value {
	switch expr := expression.(type) {
	case NumberExpression:
		if strings.ContainsRune(expr.literal, '.') {
			f, err := strconv.ParseFloat(expr.literal, 64)
			if err != nil {
				panic(err)
			}
			return constant.NewFloat(types.Double, float64(f))
		}

		i, err := strconv.ParseInt(expr.literal, 10, 64)
		if err != nil {
			panic(err)
		}
		return constant.NewInt(types.I64, i)

	case BinaryOperatorExpression:
		left := codegen(entry, expr.left)
		right := codegen(entry, expr.right)
		if !left.Type().Equal(right.Type()) {
			panic(fmt.Sprintf("unmatched type: %s, %s", left.String(), right.String()))
		}

		switch expr.operator {
		case '+':
			if types.IsInt(left.Type()) {
				return entry.NewAdd(left, right)
			}
			return entry.NewFAdd(left, right)
		case '-':
			if types.IsInt(left.Type()) {
				return entry.NewSub(left, right)
			}
			return entry.NewFSub(left, right)
		case '*':
			if types.IsInt(left.Type()) {
				return entry.NewMul(left, right)
			}
			return entry.NewFMul(left, right)
		case '/':
			if types.IsInt(left.Type()) {
				return entry.NewSDiv(left, right)
			}
			return entry.NewFDiv(left, right)
		default:
			panic("???")
		}
	}
	panic("???")
}
