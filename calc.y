%{
package llvmcalc
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
      $$ = BinaryOperatorExpression{left: $1, operator: '+', right: $3}
    }
  |
    expr '-' expr {
      $$ = BinaryOperatorExpression{left: $1, operator: '-', right: $3}
    }
  |
    expr '*' expr {
      $$ = BinaryOperatorExpression{left: $1, operator: '*', right: $3}
    }
  |
    expr '/' expr {
      $$ = BinaryOperatorExpression{left: $1, operator: '/', right: $3}
    }

primary:
    NUMBER {
      $$ = NumberExpression{literal: $1.literal}
    }
  |
    '(' expr ')' {
      $$ = $2
    }
%%
