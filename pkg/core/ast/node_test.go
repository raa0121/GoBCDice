package ast

import (
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
		{&RRollList{}, "RRollList"},
		{&Calc{}, "Calc"},
		{&Choice{}, "Choice"},

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
		{&RRoll{}, "RRoll"},
		{&RandomNumber{}, "RandomNumber"},

		{&Int{}, "Int"},
		{&String{}, "String"},
		{&Nil{}, "Nil"},
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

func TestNode_IsNil(t *testing.T) {
	testcases := []struct {
		node     Node
		expected bool
	}{
		{&DRollExpr{}, false},
		{&DRollComp{}, false},
		{&BRollList{}, false},
		{&BRollComp{}, false},
		{&RRollList{}, false},
		{&Calc{}, false},
		{&Choice{}, false},

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

		{&DRoll{}, false},
		{&BRoll{}, false},
		{&RRoll{}, false},
		{&RandomNumber{}, false},

		{&Int{}, false},
		{&String{}, false},
		{&Nil{}, true},
		{&SumRollResult{}, false},
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
		{&DRollExpr{}, false},
		{&DRollComp{}, false},
		{&BRollList{}, false},
		{&BRollComp{}, false},
		{&RRollList{}, false},
		{&Calc{}, false},
		{&Choice{}, false},

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
		{&RRoll{}, true},
		{&RandomNumber{}, true},

		{&Int{}, true},
		{&String{}, true},
		{&Nil{}, true},
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
