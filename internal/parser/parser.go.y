%{
package parser

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/lexer"
	"github.com/raa0121/GoBCDice/internal/token"
	"strconv"
)

%}

%union{
	token token.Token
	expr ast.Node
}

%token<token> ILLEGAL

%token<token> IDENT
%token<token> INT

%token<token> PLUS
%token<token> MINUS
%token<token> ASTERISK
%token<token> SLASH

%token<token> EQ
%token<token> LT
%token<token> GT

%token<token> L_PAREN
%token<token> R_PAREN
%token<token> L_BRACKET
%token<token> R_BRACKET

%token<token> D_ROLL
%token<token> B_ROLL
%token<token> R_ROLL
%token<token> U_ROLL
%token<token> SECRET

%token<token> CALC
%token<token> CHOICE

%type<expr> command
%type<expr> d_roll
%type<expr> calc
%type<expr> int_expr
%type<expr> int

%left PLUS, MINUS
%left ASTERISK, SLASH

%%

command
	: d_roll
	{
		$$ = $1
		yylex.(*LexerWrapper).ast = $$
	}
	| calc
	{
		$$ = $1
		yylex.(*LexerWrapper).ast = $$
	}

d_roll
	: int D_ROLL int
	{
		$$ = ast.NewDRoll($1, $2, $3)
	}

calc
	: CALC L_PAREN int_expr R_PAREN
	{
		$$ = &ast.Calc{
			Tok: $1,
			Expression: $3,
		}
	}

int_expr
	: int
	{
		$$ = $1
	}
	| L_PAREN int_expr R_PAREN
	{
		$$ = $2
	}
	| int_expr PLUS int_expr
	{
		$$ = ast.NewInfixExpression($1, $2, $3)
	}
	| int_expr MINUS int_expr
	{
		$$ = ast.NewInfixExpression($1, $2, $3)
	}
	| int_expr ASTERISK int_expr
	{
		$$ = ast.NewInfixExpression($1, $2, $3)
	}
	| int_expr SLASH int_expr
	{
		$$ = ast.NewInfixExpression($1, $2, $3)
	}

int
	: INT
	{
		// TODO: 整数が大きすぎるときなどのエラー処理が必要
		value, _ := strconv.Atoi($1.Literal)

		$$ = &ast.Int{
			Tok: $1,
			Value: value,
		}
	}

%%

type LexerWrapper struct {
	Input string
	Column int
	ast ast.Node
	lexer *lexer.Lexer
	err string
}

var tokenTypeToYYTokenType = map[token.TokenType]int {
	token.ILLEGAL: ILLEGAL,

	token.IDENT: IDENT,
	token.INT: INT,

	token.PLUS: PLUS,
	token.MINUS: MINUS,
	token.ASTERISK: ASTERISK,
	token.SLASH: SLASH,

	token.EQ: EQ,
	token.LT: LT,
	token.GT: GT,

	token.L_PAREN: L_PAREN,
	token.R_PAREN: R_PAREN,
	token.L_BRACKET: L_BRACKET,
	token.R_BRACKET: R_BRACKET,

	token.D_ROLL: D_ROLL,
	token.B_ROLL: B_ROLL,
	token.R_ROLL: R_ROLL,
	token.U_ROLL: U_ROLL,
	token.SECRET: SECRET,

	token.CALC: CALC,
	token.CHOICE: CHOICE,
}

func newLexerWrapper(input string) *LexerWrapper {
	lw := &LexerWrapper{
		Input: input,
		lexer: lexer.New(input),
	}

	return lw
}

func (lw *LexerWrapper) Lex(lval *yySymType) int {
	tok := lw.lexer.NextToken()
	lw.Column = tok.Column

	if tok.Type == token.EOT {
		return 0
	}

	lval.token = tok

	return tokenTypeToYYTokenType[tok.Type]
}

func (lw *LexerWrapper) Error(e string) {
	lw.err = fmt.Sprintf("column %d: %s", lw.Column, e)
}

func Parse(input string) (ast.Node, error) {
	lw := newLexerWrapper(input)

	if yyParse(lw) != 0 {
		return nil, fmt.Errorf(lw.err)
	} else {
		return lw.ast, nil
	}
}

/* vim: set filetype=goyacc: */
