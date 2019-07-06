package lexer

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
	"testing"
)

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
				{token.D_ROLL, "d", 2},
				{token.INT, "10", 3},
				{token.SLASH, "/", 5},
				{token.INT, "4", 6},
				{token.PLUS, "+", 7},
				{token.INT, "2", 8},
				{token.ASTERISK, "*", 9},
				{token.INT, "1", 10},
				{token.D_ROLL, "D", 11},
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
				{token.D_ROLL, "d", 12},
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
				{token.D_ROLL, "D", 8},
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
				{token.D_ROLL, "D", 7},
				{token.INT, "6", 8},
				{token.EOT, "", 9},
			},
		},
		{
			input: "2d6/3u",
			expectations: []tokenExpectation{
				{token.INT, "2", 1},
				{token.D_ROLL, "d", 2},
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
				{token.D_ROLL, "d", 2},
				{token.INT, "6", 3},
				{token.SLASH, "/", 4},
				{token.INT, "3", 5},
				{token.R, "r", 6},
				{token.EOT, "", 7},
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
