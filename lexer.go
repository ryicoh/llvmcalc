package llvmcalc

import (
	"text/scanner"
)

type Lexer struct {
	Scanner scanner.Scanner
	result  Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scanner.Scan())
	switch token {
	case scanner.Int, scanner.Float:
		token = NUMBER
	}
	lval.token = Token{token: token, literal: l.Scanner.TokenText()}

	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}
