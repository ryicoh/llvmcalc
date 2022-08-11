%{
package llvmcalc

type Expression interface{}
type Token struct {
    token   int
    literal string
}

type NumExpr struct {
    literal string
}
type BinOpExpr struct {
    left     Expression
    operator rune
    right    Expression
}
%}

%union{
  token Token
  expr Expression
}

%token<token> NUMBER
%type<expr> program
%type<expr> expr
%type<expr> primary

%left '+' '-'
%left '*' '/'

%%

program:
  expr {
    $$ = $1
    yylex.(*Lexer).result = $$
  }

expr:
    primary
  |
    expr '+' expr {
      $$ = BinOpExpr{left: $1, operator: '+', right: $3}
    }
  |
    expr '-' expr {
      $$ = BinOpExpr{left: $1, operator: '-', right: $3}
    }
  |
    expr '*' expr {
      $$ = BinOpExpr{left: $1, operator: '*', right: $3}
    }
  |
    expr '/' expr {
      $$ = BinOpExpr{left: $1, operator: '/', right: $3}
    }

primary:
    NUMBER {
      $$ = NumExpr{literal: $1.literal}
    }
  |
    '(' expr ')' {
      $$ = $2
    }
%%

func Parse(yylex yyLexer) int {
  yyDebug = 1
  yyErrorVerbose = true
  return yyParse(yylex);
}
