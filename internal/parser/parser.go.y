%{
package parser

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/lexer"
	"github.com/raa0121/GoBCDice/internal/token"
	"strconv"
)

type Expression interface{
	
}

type IntNode struct{
	value int
	token token.Token
}

%}

%union{
	token token.Token
	expr Expression
}

%type<expr> command
%type<expr> expr

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

%%

command
	: expr
	{
		$$ = $1
		yylex.(*LexerWrapper).result = $$
	}

expr
	: INT
	{
		value, _ := strconv.Atoi($1.Literal)
		$$ = IntNode{
			value: value,
			token: $1,
		}
	}

%%

type LexerWrapper struct {
	Input string
	Column int
	result Expression
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
	lw.err = e
}

func Parse(input string) (Expression, error) {
	lw := newLexerWrapper(input)

	if yyParse(lw) != 0 {
		return nil, fmt.Errorf(lw.err)
	} else {
		return lw.result, nil
	}
}

/* vim: set filetype=goyacc: */
