package llvmcalc

type Expression interface{}
type Token struct {
	token   int
	literal string
}

type NumberExpression struct {
	literal string
}
type BinaryOperatorExpression struct {
	left     Expression
	operator rune
	right    Expression
}

func Parse(yylex yyLexer) int {
	yyDebug = 1
	yyErrorVerbose = true
	return yyParse(yylex)
}
