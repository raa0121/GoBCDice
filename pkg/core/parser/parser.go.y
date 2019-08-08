%{
/*
BCDiceコマンドの構文解析処理のパッケージ。

BCDiceのコマンドはLALR(1)文法で表現できる。
BCDiceコマンドのLALR(1)構文解析器は、goyaccを使用して生成する。
*/
package parser

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/lexer"
	"github.com/raa0121/GoBCDice/pkg/core/token"
	"strconv"
)

%}

%union{
	token token.Token
	node ast.Node
	bRoll *ast.BRoll
	bRollList *ast.BRollList
	choiceItemList *ast.Choice
	str *ast.String
}

%token<token> ILLEGAL

%token<token> IDENT
%token<token> INT
%token<token> STRING

%token<token> PLUS
%token<token> MINUS
%token<token> ASTERISK
%token<token> SLASH

%token<token> EQ
%token<token> LT
%token<token> GT
%token<token> LTEQ
%token<token> GTEQ
%token<token> DIAMOND

%token<token> L_PAREN
%token<token> R_PAREN
%token<token> L_BRACKET
%token<token> R_BRACKET
%token<token> COMMA

%token<token> D
%token<token> B
%token<token> R
%token<token> U
%token<token> SECRET
%token<token> DOTS

%token<token> CALC

%token<token> CHOICE_BEGIN
%token<token> CHOICE_END

%type<node> command
%type<choiceItemList> choice_item_list
%type<node> int_expr
%type<node> int_rand_expr
%type<node> d_roll_expr
%type<node> d_roll_comp
%type<bRollList> b_roll_list
%type<node> b_roll_comp
%type<node> roll_operand
%type<node> d_roll
%type<bRoll> b_roll
%type<node> rand_operand
%type<node> rand
%type<node> int
%type<str> string

%nonassoc EQ, LT, GT, LTEQ, GTEQ, DIAMOND
%left PLUS, MINUS
%left ASTERISK, SLASH
%nonassoc D, B
%nonassoc DOTS
%nonassoc UPLUS, UMINUS

%%

command
	: d_roll_expr
	{
		$$ = ast.NewDRollExpr($1.Token(), $1)
		yylex.(*LexerWrapper).ast = $$
	}
	| d_roll_comp
	{
		$$ = ast.NewDRollComp($1.Token(), $1)
		yylex.(*LexerWrapper).ast = $$
	}
	| b_roll_list
	{
		$$ = $1
		yylex.(*LexerWrapper).ast = $$
	}
	| b_roll_comp
	{
		$$ = ast.NewBRollComp($1.Token(), $1)
		yylex.(*LexerWrapper).ast = $$
	}
	| CALC L_PAREN int_expr R_PAREN
	{
		$$ = ast.NewCalc($1, $3)
		yylex.(*LexerWrapper).ast = $$
	}
	| CHOICE_BEGIN choice_item_list comma_opt CHOICE_END
	{
		$$ = $2
		yylex.(*LexerWrapper).ast = $$
	}


int_expr
	: int
	| L_PAREN int_expr R_PAREN
	{
		$$ = $2
	}
	| MINUS int_expr %prec UMINUS
	{
		$$ = ast.NewUnaryMinus($1, $2)
	}
	| PLUS int_expr %prec UPLUS
	{
		$$ = $2
	}
	| int_expr PLUS int_expr
	{
		$$ = ast.NewAdd($1, $2, $3)
	}
	| int_expr MINUS int_expr
	{
		$$ = ast.NewSubtract($1, $2, $3)
	}
	| int_expr ASTERISK int_expr
	{
		$$ = ast.NewMultiply($1, $2, $3)
	}
	| int_expr SLASH int_expr
	{
		$$ = ast.NewDivideWithRoundingDown($1, $2, $3)
	}
	| int_expr SLASH int_expr U
	{
		$$ = ast.NewDivideWithRoundingUp($1, $2, $3)
	}
	| int_expr SLASH int_expr R
	{
		$$ = ast.NewDivideWithRounding($1, $2, $3)
	}

int_rand_expr
	: int
	| rand
	| L_PAREN int_rand_expr R_PAREN
	{
		$$ = $2
	}
	| MINUS int_rand_expr %prec UMINUS
	{
		$$ = ast.NewUnaryMinus($1, $2)
	}
	| PLUS int_rand_expr %prec UPLUS
	{
		$$ = $2
	}
	| int_rand_expr PLUS int_rand_expr
	{
		$$ = ast.NewAdd($1, $2, $3)
	}
	| int_rand_expr MINUS int_rand_expr
	{
		$$ = ast.NewSubtract($1, $2, $3)
	}
	| int_rand_expr ASTERISK int_rand_expr
	{
		$$ = ast.NewMultiply($1, $2, $3)
	}
	| int_rand_expr SLASH int_rand_expr
	{
		$$ = ast.NewDivideWithRoundingDown($1, $2, $3)
	}
	| int_rand_expr SLASH int_rand_expr U
	{
		$$ = ast.NewDivideWithRoundingUp($1, $2, $3)
	}
	| int_rand_expr SLASH int_rand_expr R
	{
		$$ = ast.NewDivideWithRounding($1, $2, $3)
	}

d_roll_expr
	: d_roll
	| L_PAREN d_roll_expr R_PAREN
	{
		$$ = $2
	}
	| MINUS d_roll_expr %prec UMINUS
	{
		$$ = ast.NewUnaryMinus($1, $2)
	}
	| PLUS d_roll_expr %prec UPLUS
	{
		$$ = $2
	}
	| d_roll_expr PLUS d_roll_expr
	{
		$$ = ast.NewAdd($1, $2, $3)
	}
	| d_roll_expr MINUS d_roll_expr
	{
		$$ = ast.NewSubtract($1, $2, $3)
	}
	| d_roll_expr ASTERISK d_roll_expr
	{
		$$ = ast.NewMultiply($1, $2, $3)
	}
	| d_roll_expr SLASH d_roll_expr
	{
		$$ = ast.NewDivideWithRoundingDown($1, $2, $3)
	}
	| d_roll_expr SLASH d_roll_expr U
	{
		$$ = ast.NewDivideWithRoundingUp($1, $2, $3)
	}
	| d_roll_expr SLASH d_roll_expr R
	{
		$$ = ast.NewDivideWithRounding($1, $2, $3)
	}
	| int_rand_expr PLUS d_roll_expr
	{
		$$ = ast.NewAdd($1, $2, $3)
	}
	| int_rand_expr MINUS d_roll_expr
	{
		$$ = ast.NewSubtract($1, $2, $3)
	}
	| int_rand_expr ASTERISK d_roll_expr
	{
		$$ = ast.NewMultiply($1, $2, $3)
	}
	| int_rand_expr SLASH d_roll_expr
	{
		$$ = ast.NewDivideWithRoundingDown($1, $2, $3)
	}
	| int_rand_expr SLASH d_roll_expr U
	{
		$$ = ast.NewDivideWithRoundingUp($1, $2, $3)
	}
	| int_rand_expr SLASH d_roll_expr R
	{
		$$ = ast.NewDivideWithRounding($1, $2, $3)
	}
	| d_roll_expr PLUS int_rand_expr
	{
		$$ = ast.NewAdd($1, $2, $3)
	}
	| d_roll_expr MINUS int_rand_expr
	{
		$$ = ast.NewSubtract($1, $2, $3)
	}
	| d_roll_expr ASTERISK int_rand_expr
	{
		$$ = ast.NewMultiply($1, $2, $3)
	}
	| d_roll_expr SLASH int_rand_expr
	{
		$$ = ast.NewDivideWithRoundingDown($1, $2, $3)
	}
	| d_roll_expr SLASH int_rand_expr U
	{
		$$ = ast.NewDivideWithRoundingUp($1, $2, $3)
	}
	| d_roll_expr SLASH int_rand_expr R
	{
		$$ = ast.NewDivideWithRounding($1, $2, $3)
	}

d_roll_comp
	: d_roll_expr EQ int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| d_roll_expr DIAMOND int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| d_roll_expr LT int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| d_roll_expr GT int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| d_roll_expr LTEQ int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| d_roll_expr GTEQ int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}

b_roll_list
	: b_roll
	{
		$$ = ast.NewBRollList($1)
	}
	| b_roll_list PLUS b_roll
	{
		$$.Append($3)
	}

b_roll_comp
	: b_roll_list EQ int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| b_roll_list DIAMOND int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| b_roll_list LT int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| b_roll_list GT int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| b_roll_list LTEQ int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}
	| b_roll_list GTEQ int_expr
	{
		$$ = ast.NewCompare($1, $2, $3)
	}

choice_item_list
	: string
	{
		$$ = ast.NewChoice($1)
	}
	| choice_item_list COMMA string
	{
		$$.Append($3)
	}

comma_opt
	:
	| COMMA

roll_operand
	: int
	| rand
	| L_PAREN int_rand_expr R_PAREN
	{
		$$ = $2
	}

d_roll
	: roll_operand D roll_operand
	{
		$$ = ast.NewDRoll($1, $2, $3)
	}

b_roll
	: roll_operand B roll_operand
	{
		$$ = ast.NewBRoll($1, $2, $3)
	}

rand_operand
	: int
	| L_PAREN int_expr R_PAREN
	{
		$$ = $2
	}

rand
	: L_BRACKET rand_operand DOTS rand_operand R_BRACKET
	{
		$$ = ast.NewRandomNumber($2, $3, $4)
	}

int
	: INT
	{
		// TODO: 整数が大きすぎるときなどのエラー処理が必要
		value, _ := strconv.Atoi($1.Literal)

		$$ = ast.NewInt(value, $1)
	}

string
	: STRING
	{
		$$ = ast.NewString($1.Literal, $1)
	}

%%

// 字句解析器をyyParseで使用できるようにするためのラッパー。
type LexerWrapper struct {
	// 入力文字列
	Input string
	// 現在の桁
	Column int
	// 現在のルートノード
	ast ast.Node
	// 字句解析器
	lexer *lexer.Lexer
	// エラーの内容
	err string
}

// トークンの種類とyyParseで使用する定数との対応
var tokenTypeToYYTokenType = map[token.TokenType]int {
	token.ILLEGAL: ILLEGAL,

	token.IDENT: IDENT,
	token.INT: INT,
	token.STRING: STRING,

	token.PLUS: PLUS,
	token.MINUS: MINUS,
	token.ASTERISK: ASTERISK,
	token.SLASH: SLASH,

	token.EQ: EQ,
	token.LT: LT,
	token.GT: GT,
	token.LTEQ: LTEQ,
	token.GTEQ: GTEQ,
	token.DIAMOND: DIAMOND,

	token.L_PAREN: L_PAREN,
	token.R_PAREN: R_PAREN,
	token.L_BRACKET: L_BRACKET,
	token.R_BRACKET: R_BRACKET,
	token.COMMA: COMMA,

	token.D: D,
	token.B: B,
	token.R: R,
	token.U: U,
	token.SECRET: SECRET,
	token.DOTS: DOTS,

	token.CALC: CALC,

	token.CHOICE_BEGIN: CHOICE_BEGIN,
	token.CHOICE_END: CHOICE_END,
}

// newLexerWrapper は新しい字句解析器ラッパーを作る。
//
// input: 入力文字列
func newLexerWrapper(input string) *LexerWrapper {
	lw := &LexerWrapper{
		Input: input,
		lexer: lexer.New(input),
	}

	return lw
}

// Lex は次のトークンを読み込み、対応する定数を返す。
//
// 文字列の終端に達した場合は0を返す。
func (lw *LexerWrapper) Lex(lval *yySymType) int {
	tok := lw.lexer.NextToken()
	lw.Column = tok.Column

	if tok.Type == token.EOT {
		return 0
	}

	lval.token = tok

	yyTokenType, ok := tokenTypeToYYTokenType[tok.Type]
	if !ok {
		return ILLEGAL
	}

	return yyTokenType
}

// Error は発生したエラーを記録する。
func (lw *LexerWrapper) Error(e string) {
	lw.err = fmt.Sprintf("column %d: %s", lw.Column, e)
}

// Parse は入力文字列をBCDiceコマンドとして構文解析する。
// 構文解析に成功した場合は、抽象構文木のルートノードを返す。
// 構文解析に失敗した場合は、エラーを返す。
func Parse(input string) (ast.Node, error) {
	lw := newLexerWrapper(input)

	if yyParse(lw) != 0 {
		return nil, fmt.Errorf(lw.err)
	} else {
		return lw.ast, nil
	}
}

/* vim: set filetype=goyacc: */
