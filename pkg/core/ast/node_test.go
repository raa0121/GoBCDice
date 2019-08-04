package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
	"testing"
)

func TestNode_Type(t *testing.T) {
	testcases := []struct {
		node     Node
		expected string
	}{
		{&DRollExpr{}, "DRollExpr"},
		{&DRollComp{}, "DRollComp"},
		{&BRollList{}, "BRollList"},
		{&BRollComp{}, "BRollComp"},
		{&Calc{}, "Calc"},

		{&PrefixExpressionImpl{}, "PrefixExpression"},
		{&UnaryMinus{}, "UnaryMinus"},

		{&InfixExpressionImpl{}, "InfixExpression"},
		{&Compare{}, "Compare"},
		{&Add{}, "Add"},
		{&Subtract{}, "Subtract"},
		{&Multiply{}, "Multiply"},
		{&DivideWithRoundingUp{}, "DivideWithRoundingUp"},
		{&DivideWithRounding{}, "DivideWithRounding"},
		{&DivideWithRoundingDown{}, "DivideWithRoundingDown"},
		{&DRoll{}, "DRoll"},
		{&BRoll{}, "BRoll"},
		{&RandomNumber{}, "RandomNumber"},

		{&Int{}, "Int"},
		{&SumRollResult{}, "SumRollResult"},
	}

	for _, test := range testcases {
		t.Run(test.expected, func(t *testing.T) {
			actual := test.node.Type().String()
			if actual != test.expected {
				t.Errorf("got: %q, want: %q", actual, test.expected)
			}
		})
	}
}

func TestNode_IsPrimaryExpression(t *testing.T) {
	testcases := []struct {
		node     Node
		expected bool
	}{
		{&DRollExpr{}, false},
		{&DRollComp{}, false},
		{&BRollList{}, false},
		{&BRollComp{}, false},
		{&Calc{}, false},

		{&PrefixExpressionImpl{}, false},
		{&UnaryMinus{}, false},

		{&InfixExpressionImpl{}, false},
		{&Compare{}, false},
		{&Add{}, false},
		{&Subtract{}, false},
		{&Multiply{}, false},
		{&DivideWithRoundingUp{}, false},
		{&DivideWithRounding{}, false},
		{&DivideWithRoundingDown{}, false},

		{&DRoll{}, true},
		{&BRoll{}, true},
		{&RandomNumber{}, true},

		{&Int{}, true},
		{&SumRollResult{}, true},
	}

	for _, test := range testcases {
		t.Run(test.node.Type().String(), func(t *testing.T) {
			actual := test.node.IsPrimaryExpression()
			if actual != test.expected {
				t.Errorf("got: %v, want: %v", actual, test.expected)
			}
		})
	}
}

func TestNode_IsVariable(t *testing.T) {
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
				token.Token{token.D, "D", 2},
				NewInt(6, token.Token{token.INT, "6", 3}),
			),
			expected: true,
		},
		{
			node: NewBRoll(
				NewInt(2, token.Token{token.INT, "2", 1}),
				token.Token{token.B, "B", 2},
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
					token.Token{token.D, "D", 3},
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
					token.Token{token.D, "D", 2},
					NewInt(6, token.Token{token.INT, "6", 3}),
				),
				token.Token{token.PLUS, "+", 4},
				NewInt(1, token.Token{token.INT, "1", 5}),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewInt(1, token.Token{token.INT, "1", 1}),
				token.Token{token.PLUS, "+", 2},
				NewDRoll(
					NewInt(2, token.Token{token.INT, "2", 3}),
					token.Token{token.D, "D", 4},
					NewInt(6, token.Token{token.INT, "6", 5}),
				),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewDRoll(
					NewInt(2, token.Token{token.INT, "2", 1}),
					token.Token{token.D, "D", 2},
					NewInt(6, token.Token{token.INT, "6", 3}),
				),
				token.Token{token.PLUS, "+", 4},
				NewDRoll(
					NewInt(3, token.Token{token.INT, "3", 5}),
					token.Token{token.D, "D", 6},
					NewInt(10, token.Token{token.INT, "10", 7}),
				),
			),
			expected: true,
		},
		{
			node: NewCalc(
				token.Token{token.CALC, "C", 1},
				NewAdd(
					NewInt(1, token.Token{token.INT, "1", 3}),
					token.Token{token.PLUS, "+", 4},
					NewInt(2, token.Token{token.INT, "2", 5}),
				),
			),
			expected: false,
		},
		{
			node: NewDRollExpr(
				token.Token{token.PLUS, "+", 4},
				NewAdd(
					NewDRoll(
						NewInt(2, token.Token{token.INT, "2", 1}),
						token.Token{token.D, "D", 2},
						NewInt(6, token.Token{token.INT, "6", 3}),
					),
					token.Token{token.PLUS, "+", 4},
					NewInt(1, token.Token{token.INT, "1", 5}),
				),
			),
			expected: true,
		},
		{
			node: NewDRollComp(
				token.Token{token.GT, ">", 4},
				NewCompare(
					NewDRoll(
						NewInt(2, token.Token{token.INT, "2", 1}),
						token.Token{token.D, "D", 2},
						NewInt(6, token.Token{token.INT, "6", 3}),
					),
					token.Token{token.GT, ">", 4},
					NewInt(7, token.Token{token.INT, "7", 5}),
				),
			),
			expected: true,
		},
		{
			node: NewBRollList(
				NewBRoll(
					NewInt(2, token.Token{token.INT, "2", 1}),
					token.Token{token.B, "B", 2},
					NewInt(6, token.Token{token.INT, "6", 3}),
				),
			),
			expected: true,
		},
		{
			node: NewBRollComp(
				token.Token{token.GT, ">", 4},
				NewCompare(
					NewBRollList(
						NewBRoll(
							NewInt(2, token.Token{token.INT, "2", 1}),
							token.Token{token.B, "B", 2},
							NewInt(6, token.Token{token.INT, "6", 3}),
						),
					),
					token.Token{token.GT, ">", 4},
					NewInt(7, token.Token{token.INT, "7", 5}),
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
