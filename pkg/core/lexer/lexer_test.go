package lexer

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
	"testing"
)

// 字句解析の例。
// token.EOTが返るまでNextTokenを呼び続け、読み込まれたトークンを出力する。
func Example() {
	lexer := New("(2*3-4)d6-1d4+1")

	for tok := lexer.NextToken(); tok.Type != token.EOT; tok = lexer.NextToken() {
		fmt.Println(tok)
	}
	// Output:
	// <Token ( Literal:"(" Column:1>
	// <Token INT Literal:"2" Column:2>
	// <Token * Literal:"*" Column:3>
	// <Token INT Literal:"3" Column:4>
	// <Token - Literal:"-" Column:5>
	// <Token INT Literal:"4" Column:6>
	// <Token ) Literal:")" Column:7>
	// <Token D Literal:"d" Column:8>
	// <Token INT Literal:"6" Column:9>
	// <Token - Literal:"-" Column:10>
	// <Token INT Literal:"1" Column:11>
	// <Token D Literal:"d" Column:12>
	// <Token INT Literal:"4" Column:13>
	// <Token + Literal:"+" Column:14>
	// <Token INT Literal:"1" Column:15>
}

type tokenExpectation struct {
	expectedType    token.TokenType
	expectedLiteral string
	expectedColumn  int
}

type tokenTestCase struct {
	input        string
	expectations []tokenExpectation
}

func TestNextToken(t *testing.T) {
	testcases := []tokenTestCase{
		// 空文字列
		{
			input: "",
			expectations: []tokenExpectation{
				{token.EOT, "", 1},
			},
		},
		{
			input: "3d10/4+2*1D6-5",
			expectations: []tokenExpectation{
				{token.INT, "3", 1},
				{token.D, "d", 2},
				{token.INT, "10", 3},
				{token.SLASH, "/", 5},
				{token.INT, "4", 6},
				{token.PLUS, "+", 7},
				{token.INT, "2", 8},
				{token.ASTERISK, "*", 9},
				{token.INT, "1", 10},
				{token.D, "D", 11},
				{token.INT, "6", 12},
				{token.MINUS, "-", 13},
				{token.INT, "5", 14},
				{token.EOT, "", 15},
			},
		},
		{
			input: "((2+3)*4/3)d6*2+5",
			expectations: []tokenExpectation{
				{token.L_PAREN, "(", 1},
				{token.L_PAREN, "(", 2},
				{token.INT, "2", 3},
				{token.PLUS, "+", 4},
				{token.INT, "3", 5},
				{token.R_PAREN, ")", 6},
				{token.ASTERISK, "*", 7},
				{token.INT, "4", 8},
				{token.SLASH, "/", 9},
				{token.INT, "3", 10},
				{token.R_PAREN, ")", 11},
				{token.D, "d", 12},
				{token.INT, "6", 13},
				{token.ASTERISK, "*", 14},
				{token.INT, "2", 15},
				{token.PLUS, "+", 16},
				{token.INT, "5", 17},
				{token.EOT, "", 18},
			},
		},
		{
			input: "[1...5]D6",
			expectations: []tokenExpectation{
				{token.L_BRACKET, "[", 1},
				{token.INT, "1", 2},
				{token.DOTS, "...", 3},
				{token.INT, "5", 6},
				{token.R_BRACKET, "]", 7},
				{token.D, "D", 8},
				{token.INT, "6", 9},
				{token.EOT, "", 10},
			},
		},
		{
			input: "[1..5]D6",
			expectations: []tokenExpectation{
				{token.L_BRACKET, "[", 1},
				{token.INT, "1", 2},
				{token.ILLEGAL, ".", 3},
				{token.ILLEGAL, ".", 4},
				{token.INT, "5", 5},
				{token.R_BRACKET, "]", 6},
				{token.D, "D", 7},
				{token.INT, "6", 8},
				{token.EOT, "", 9},
			},
		},
		{
			input: "2d6/3u",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.SLASH, "/", 4},
				{token.INT, "3", 5},
				{token.U, "u", 6},
				{token.EOT, "", 7},
			},
		},
		{
			input: "2d6/3r",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.SLASH, "/", 4},
				{token.INT, "3", 5},
				{token.R, "r", 6},
				{token.EOT, "", 7},
			},
		},
		{
			input: "2d6=7",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.EQ, "=", 4},
				{token.INT, "7", 5},
			},
		},
		{
			input: "2d6>7",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.GT, ">", 4},
				{token.INT, "7", 5},
			},
		},
		{
			input: "2d6<7",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.LT, "<", 4},
				{token.INT, "7", 5},
			},
		},
		{
			input: "2d6>=7",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.GTEQ, ">=", 4},
				{token.INT, "7", 6},
			},
		},
		{
			input: "2d6<=7",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.LTEQ, "<=", 4},
				{token.INT, "7", 6},
			},
		},
		{
			input: "2d6<>7",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D, "d", 2},
				{token.INT, "6", 3},
				{token.DIAMOND, "<>", 4},
				{token.INT, "7", 6},
			},
		},
		{
			input: "2b6+4b10",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.B, "b", 2},
				{token.INT, "6", 3},
				{token.PLUS, "+", 4},
				{token.INT, "4", 5},
				{token.B, "b", 6},
				{token.INT, "10", 7},
			},
		},
		{
			input: "2b6+4b10>3",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.B, "b", 2},
				{token.INT, "6", 3},
				{token.PLUS, "+", 4},
				{token.INT, "4", 5},
				{token.B, "b", 6},
				{token.INT, "10", 7},
				{token.GT, ">", 9},
				{token.INT, "3", 10},
			},
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			l := New(test.input)

			for _, e := range test.expectations {
				var name string

				if e.expectedType == token.EOT {
					name = "EOT"
				} else {
					name = fmt.Sprintf("%q", e.expectedLiteral)
				}

				t.Run(name, func(t *testing.T) {
					tok := l.NextToken()

					if tok.Type != e.expectedType {
						t.Errorf("Type: got: %q, want: %q",
							tok.Type, e.expectedType)
					}

					if tok.Literal != e.expectedLiteral {
						t.Errorf("Literal: got: %q, want: %q",
							tok.Literal, e.expectedLiteral)
					}

					if tok.Column != e.expectedColumn {
						t.Errorf("Column: got: %d, want: %d",
							tok.Column, e.expectedColumn)
					}
				})
			}
		})
	}
}
