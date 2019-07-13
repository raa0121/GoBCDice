package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
	"testing"
)

func TestIsVariable(t *testing.T) {
	testcases := []struct {
		node     Node
		expected bool
	}{
		{
			node:     NewInt(42, token.Token{token.INT, "42", 1}),
			expected: false,
		},
		{
			node: NewDRoll(
				NewInt(2, token.Token{token.INT, "2", 1}),
				token.Token{token.D_ROLL, "D", 2},
				NewInt(6, token.Token{token.INT, "6", 3}),
			),
			expected: true,
		},
		{
			node: NewRandomNumber(
				NewInt(3, token.Token{token.INT, "3", 1}),
				token.Token{token.DOTS, "...", 2},
				NewInt(5, token.Token{token.INT, "5", 5}),
			),
			expected: true,
		},
		{
			node: NewUnaryMinus(
				token.Token{token.MINUS, "-", 1},
				NewInt(1, token.Token{token.INT, "1", 2}),
			),
			expected: false,
		},
		{
			node: NewUnaryMinus(
				token.Token{token.MINUS, "-", 1},
				NewDRoll(
					NewInt(2, token.Token{token.INT, "2", 2}),
					token.Token{token.D_ROLL, "D", 3},
					NewInt(6, token.Token{token.INT, "6", 4}),
				),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewInt(1, token.Token{token.INT, "1", 1}),
				token.Token{token.PLUS, "+", 2},
				NewInt(2, token.Token{token.INT, "2", 3}),
			),
			expected: false,
		},
		{
			node: NewAdd(
				NewDRoll(
					NewInt(2, token.Token{token.INT, "2", 1}),
					token.Token{token.D_ROLL, "D", 2},
					NewInt(6, token.Token{token.INT, "6", 3}),
				),
				token.Token{token.PLUS, "+", 4},
				NewInt(2, token.Token{token.INT, "2", 5}),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewInt(1, token.Token{token.INT, "2", 1}),
				token.Token{token.PLUS, "+", 2},
				NewDRoll(
					NewInt(2, token.Token{token.INT, "2", 3}),
					token.Token{token.D_ROLL, "D", 4},
					NewInt(6, token.Token{token.INT, "6", 5}),
				),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewDRoll(
					NewInt(2, token.Token{token.INT, "2", 1}),
					token.Token{token.D_ROLL, "D", 2},
					NewInt(6, token.Token{token.INT, "6", 3}),
				),
				token.Token{token.PLUS, "+", 4},
				NewDRoll(
					NewInt(3, token.Token{token.INT, "3", 5}),
					token.Token{token.D_ROLL, "D", 6},
					NewInt(10, token.Token{token.INT, "10", 7}),
				),
			),
			expected: true,
		},
	}

	for _, test := range testcases {
		t.Run(test.node.SExp(), func(t *testing.T) {
			actual := test.node.IsVariable()
			if actual != test.expected {
				t.Errorf("got: %v, want: %v", actual, test.expected)
			}
		})
	}
}
