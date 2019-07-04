package evaluator

import (
	"github.com/raa0121/GoBCDice/internal/object"
	"github.com/raa0121/GoBCDice/internal/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
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
	}

	for i, test := range testcases {
		ast, err := parser.Parse(test.input)
		if err != nil {
			t.Errorf("#%d: 構文解析エラー: %s", i, err)
			continue
		}

		// ノードを評価する
		evaluated := Eval(ast)

		// 型が合っているか？
		obj, ok := evaluated.(*object.Integer)
		if !ok {
			t.Errorf("整数オブジェクトでない: got=%T (%+v)", obj, obj)
			continue
		}

		if obj.Value != test.expected {
			t.Errorf("異なる値: got=%d, want=%d", obj.Value, test.expected)
		}
	}
}
