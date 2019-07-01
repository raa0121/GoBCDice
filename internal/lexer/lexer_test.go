package lexer

import (
	"github.com/raa0121/GoBCDice/internal/token"
	"testing"
)

type tokenExpectation struct {
	expectedType    token.TokenType
	expectedLiteral string
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
				{token.EOT, ""},
			},
		},
		{
			input: "3d10/4+2*1D6-5",
			expectations: []tokenExpectation{
				{token.INT, "3"},
				{token.D_ROLL, "d"},
				{token.INT, "10"},
				{token.SLASH, "/"},
				{token.INT, "4"},
				{token.PLUS, "+"},
				{token.INT, "2"},
				{token.ASTERISK, "*"},
				{token.INT, "1"},
				{token.D_ROLL, "D"},
				{token.INT, "6"},
				{token.MINUS, "-"},
				{token.INT, "5"},
				{token.EOT, ""},
			},
		},
		{
			input: "((2+3)*4/3)d6*2+5",
			expectations: []tokenExpectation{
				{token.L_PAREN, "("},
				{token.L_PAREN, "("},
				{token.INT, "2"},
				{token.PLUS, "+"},
				{token.INT, "3"},
				{token.R_PAREN, ")"},
				{token.ASTERISK, "*"},
				{token.INT, "4"},
				{token.SLASH, "/"},
				{token.INT, "3"},
				{token.R_PAREN, ")"},
				{token.D_ROLL, "d"},
				{token.INT, "6"},
				{token.ASTERISK, "*"},
				{token.INT, "2"},
				{token.PLUS, "+"},
				{token.INT, "5"},
				{token.EOT, ""},
			},
		},
	}

	for i, testcase := range testcases {
		l := New(testcase.input)

		for j, e := range testcase.expectations {
			tok := l.NextToken()

			if tok.Type != e.expectedType {
				t.Errorf("#%d-%d: Type: got: %q, want: %q",
					i, j, tok.Type, e.expectedType)
			}

			if tok.Literal != e.expectedLiteral {
				t.Errorf("#%d-%d: Literal: got: %q, want: %q",
					i, j, tok.Literal, e.expectedLiteral)
			}
		}
	}
}
