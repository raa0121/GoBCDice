package evaluator

import (
	"github.com/raa0121/GoBCDice/internal/object"
	"github.com/raa0121/GoBCDice/internal/parser"
	"testing"
)

func TestEvalCalc(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{"C(5)", 5},
		{"C(10)", 10},
		{"C(42)", 42},
		{"C(65535)", 65535},
		{"C(-5)", -5},
		{"C(-10)", -10},
		{"C(-42)", -42},
		{"C(-65535)", -65535},
		{"C(+5)", 5},
		{"C(+10)", 10},
		{"C(+42)", 42},
		{"C(+65535)", 65535},
		{"C(1+2)", 3},
		{"C(1-2)", -1},
		{"C(1*2)", 2},
		{"C(1/2)", 0},
		{"C(-1+2)", 1},
		{"C(+1+2)", 3},
		{"C(1+2-3)", 0},
		{"C(1*2+3)", 5},
		{"C(1/2+3)", 3},
		{"C(1+2*3)", 7},
		{"C(1+2/3)", 1},
		{"C(1+(2-3))", 0},
		{"C((1+2)*3)", 9},
		{"C((1+2)/3)", 1},
		{"C(-(1+2))", -3},
		{"C(+(1+2))", 3},
	}

	for i, test := range testcases {
		input := test.input
		ast, parseErr := parser.Parse(input)
		if parseErr != nil {
			t.Errorf("#%d: %q: 構文解析エラー: %s", i, input, parseErr)
			continue
		}

		// ノードを評価する
		evaluated, evalErr := Eval(ast)
		if evalErr != nil {
			t.Errorf("#%d: %q: 評価エラー: %s", i, input, evalErr)
			continue
		}

		if evaluated == nil {
			t.Errorf("#%d: %q: Evalの対象外 (nil)", i, input)
			continue
		}

		// 型が合っているか？
		obj, typeMatched := evaluated.(*object.Integer)
		if !typeMatched {
			t.Errorf("#%d: %q: 整数オブジェクトでない: got=%T (%+v)", i, input, obj, obj)
			continue
		}

		if obj.Value != test.expected {
			t.Errorf("#%d: %q: 異なる値: got=%d, want=%d", i, input, obj.Value, test.expected)
		}
	}
}
