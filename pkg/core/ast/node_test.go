package ast

import (
	"testing"
)

func TestNode_Type(t *testing.T) {
	testcases := []struct {
		node     Node
		expected string
	}{
		{NewDRollExpr(nil), "DRollExpr"},
		{NewDRollComp(nil), "DRollComp"},
		{NewBRollList(nil), "BRollList"},
		{NewBRollComp(nil), "BRollComp"},
		{NewRRollList(nil, nil), "RRollList"},
		{NewURollExpr(nil, nil), "URollExpr"},
		{NewURollComp(nil), "URollComp"},
		{NewCalc(nil), "Calc"},
		{NewChoice(nil), "Choice"},

		{NewUnaryMinus(nil), "UnaryMinus"},

		{NewCompare(nil, "", nil), "Compare"},
		{NewAdd(nil, nil), "Add"},
		{NewSubtract(nil, nil), "Subtract"},
		{NewMultiply(nil, nil), "Multiply"},
		{NewDivideWithRoundingUp(nil, nil), "DivideWithRoundingUp"},
		{NewDivideWithRounding(nil, nil), "DivideWithRounding"},
		{NewDivideWithRoundingDown(nil, nil), "DivideWithRoundingDown"},
		{NewDRoll(nil, nil), "DRoll"},
		{NewBRoll(nil, nil), "BRoll"},
		{NewRRoll(nil, nil), "RRoll"},
		{NewURoll(nil, nil), "RRoll"},
		{NewRandomNumber(nil, nil), "RandomNumber"},

		{NewInt(0), "Int"},
		{NewString(""), "String"},
		{NilInstance(), "Nil"},
		{NewSumRollResult(nil), "SumRollResult"},
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

func TestNode_IsNil(t *testing.T) {
	testcases := []struct {
		node     Node
		expected bool
	}{
		{NewDRollExpr(nil), false},
		{NewDRollComp(nil), false},
		{NewBRollList(nil), false},
		{NewBRollComp(nil), false},
		{NewRRollList(nil, nil), false},
		{NewURollExpr(nil, nil), false},
		{NewURollComp(nil), false},
		{NewCalc(nil), false},
		{NewChoice(nil), false},

		{NewUnaryMinus(nil), false},

		{NewCompare(nil, "", nil), false},
		{NewAdd(nil, nil), false},
		{NewSubtract(nil, nil), false},
		{NewMultiply(nil, nil), false},
		{NewDivideWithRoundingUp(nil, nil), false},
		{NewDivideWithRounding(nil, nil), false},
		{NewDivideWithRoundingDown(nil, nil), false},

		{NewDRoll(nil, nil), false},
		{NewBRoll(nil, nil), false},
		{NewRRoll(nil, nil), false},
		{NewRandomNumber(nil, nil), false},

		{NewInt(0), false},
		{NewString(""), false},
		{NilInstance(), true},
		{NewSumRollResult(nil), false},
	}

	for _, test := range testcases {
		t.Run(test.node.Type().String(), func(t *testing.T) {
			actual := test.node.IsNil()
			if actual != test.expected {
				t.Errorf("got: %v, want: %v", actual, test.expected)
			}
		})
	}
}

func TestNode_IsPrimaryExpression(t *testing.T) {
	testcases := []struct {
		node     Node
		expected bool
	}{
		{NewDRollExpr(nil), false},
		{NewDRollComp(nil), false},
		{NewBRollList(nil), false},
		{NewBRollComp(nil), false},
		{NewRRollList(nil, nil), false},
		{NewURollExpr(nil, nil), false},
		{NewCalc(nil), false},
		{NewChoice(nil), false},

		{NewUnaryMinus(nil), false},

		{NewCompare(nil, "", nil), false},
		{NewAdd(nil, nil), false},
		{NewSubtract(nil, nil), false},
		{NewMultiply(nil, nil), false},
		{NewDivideWithRoundingUp(nil, nil), false},
		{NewDivideWithRounding(nil, nil), false},
		{NewDivideWithRoundingDown(nil, nil), false},

		{NewDRoll(nil, nil), true},
		{NewBRoll(nil, nil), true},
		{NewRRoll(nil, nil), true},
		{NewRandomNumber(nil, nil), true},

		{NewInt(0), true},
		{NewString(""), true},
		{NilInstance(), true},
		{NewSumRollResult(nil), true},
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
			node:     NilInstance(),
			expected: false,
		},
		{
			node:     NewInt(42),
			expected: false,
		},
		{
			node: NewDRoll(
				NewInt(2),
				NewInt(6),
			),
			expected: true,
		},
		{
			node: NewBRoll(
				NewInt(2),
				NewInt(6),
			),
			expected: true,
		},
		{
			node: NewRRoll(
				NewInt(2),
				NewInt(6),
			),
			expected: true,
		},
		{
			node: NewURoll(
				NewInt(2),
				NewInt(6),
			),
			expected: true,
		},
		{
			node: NewRandomNumber(
				NewInt(3),
				NewInt(5),
			),
			expected: true,
		},
		{
			node: NewUnaryMinus(
				NewInt(1),
			),
			expected: false,
		},
		{
			node: NewUnaryMinus(
				NewDRoll(
					NewInt(2),
					NewInt(6),
				),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewInt(1),
				NewInt(2),
			),
			expected: false,
		},
		{
			node: NewAdd(
				NewDRoll(
					NewInt(2),
					NewInt(6),
				),
				NewInt(1),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewInt(1),
				NewDRoll(
					NewInt(2),
					NewInt(6),
				),
			),
			expected: true,
		},
		{
			node: NewAdd(
				NewDRoll(
					NewInt(2),
					NewInt(6),
				),
				NewDRoll(
					NewInt(3),
					NewInt(10),
				),
			),
			expected: true,
		},
		{
			node: NewCalc(
				NewAdd(
					NewInt(1),
					NewInt(2),
				),
			),
			expected: false,
		},
		{
			node: NewDRollExpr(
				NewAdd(
					NewDRoll(
						NewInt(2),
						NewInt(6),
					),
					NewInt(1),
				),
			),
			expected: true,
		},
		{
			node: NewDRollComp(
				NewCompare(
					NewDRoll(
						NewInt(2),
						NewInt(6),
					),
					">",
					NewInt(7),
				),
			),
			expected: true,
		},
		{
			node: NewBRollList(
				NewBRoll(
					NewInt(2),
					NewInt(6),
				),
			),
			expected: true,
		},
		{
			node: NewBRollComp(
				NewCompare(
					NewBRollList(
						NewBRoll(
							NewInt(2),
							NewInt(6),
						),
					),
					">",
					NewInt(7),
				),
			),
			expected: true,
		},
		{
			node: NewRRollList(
				NewRRoll(
					NewInt(2),
					NewInt(6),
				),
				NilInstance(),
			),
			expected: true,
		},
		{
			node: NewRRollComp(
				NewCompare(
					NewRRollList(
						NewRRoll(
							NewInt(2),
							NewInt(6),
						),
						NilInstance(),
					),
					">=",
					NewInt(3),
				),
			),
			expected: true,
		},
		{
			node: NewURollExpr(
				NewRRollList(
					NewURoll(
						NewInt(2),
						NewInt(6),
					),
					NilInstance(),
				),
				nil,
			),
			expected: true,
		},
		{
			node: NewURollExpr(
				NewRRollList(
					NewURoll(
						NewInt(2),
						NewInt(6),
					),
					NilInstance(),
				),
				NewAdd(nil, NewInt(1)),
			),
			expected: true,
		},
		{
			node: NewURollComp(
				NewCompare(
					NewURollExpr(
						NewRRollList(
							NewURoll(
								NewInt(3),
								NewInt(6),
							),
							NewInt(6),
						),
						nil,
					),
					">=",
					NewInt(10),
				),
			),
			expected: true,
		},
		{
			node: NewRRollList(
				NewURoll(
					NewInt(2),
					NewInt(6),
				),
				NilInstance(),
			),
			expected: true,
		},
		{
			node:     NewString("hello"),
			expected: false,
		},
		{
			node:     NewChoice(NewString("hello")),
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
