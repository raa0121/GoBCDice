package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"reflect"
	"testing"
)

func TestEvalDRollExpr(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
		dice     []dice.Die
	}{
		{
			input:    "2D6",
			expected: 8,
			dice:     []dice.Die{{5, 6}, {3, 6}},
		},
		{
			input:    "2D4",
			expected: 3,
			dice:     []dice.Die{{1, 4}, {2, 4}},
		},
		{
			input:    "2D6+1",
			expected: 9,
			dice:     []dice.Die{{2, 6}, {6, 6}},
		},
		{
			input:    "1+2D6",
			expected: 8,
			dice:     []dice.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2d6+1-1-2-3-4",
			expected: -2,
			dice:     []dice.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "2D6+4D10",
			expected: 30,
			dice:     []dice.Die{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
		},
		{
			input:    "2d6*3",
			expected: 18,
			dice:     []dice.Die{{2, 6}, {4, 6}},
		},
		{
			input:    "2d10+3-4",
			expected: 7,
			dice:     []dice.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d10+3*4",
			expected: 20,
			dice:     []dice.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d6*3-1d6+1",
			expected: 22,
			dice:     []dice.Die{{6, 6}, {2, 6}, {3, 6}},
		},
		{
			input:    "(2+3)d6-1+3d6+2",
			expected: 31,
			dice:     []dice.Die{{2, 6}, {3, 6}, {1, 6}, {5, 6}, {6, 6}, {5, 6}, {4, 6}, {4, 6}},
		},
		{
			input:    "(2*3-4)d6-1d4+1",
			expected: 10,
			dice:     []dice.Die{{6, 6}, {5, 6}, {2, 4}},
		},
		{
			input:    "((2+3)*4/3)d6*2+5",
			expected: 53,
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {2, 6}, {1, 6}, {4, 6}},
		},
		{
			input:    "1D6/2",
			expected: 0,
			dice:     []dice.Die{{1, 6}},
		},
		{
			input:    "3D6/2",
			expected: 3,
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "3D6/2+1D6",
			expected: 8,
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2",
			expected: 9,
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2U",
			expected: 10,
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "5D6/10",
			expected: 2,
			dice:     []dice.Die{{6, 6}, {6, 6}, {6, 6}, {6, 6}, {5, 6}},
		},
		{
			input:    "3D6/2U",
			expected: 4,
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "5D6/10u",
			expected: 3,
			dice:     []dice.Die{{6, 6}, {6, 6}, {6, 6}, {2, 6}, {1, 6}},
		},
		{
			input:    "1D100/10R",
			expected: 6,
			dice:     []dice.Die{{55, 100}},
		},
		{
			input:    "1D100/10r",
			expected: 5,
			dice:     []dice.Die{{54, 100}},
		},
		{
			input:    "[1...5]D6",
			expected: 15,
			dice:     []dice.Die{{4, 5}, {5, 6}, {3, 6}, {4, 6}, {3, 6}},
		},
		{
			input:    "([2...4]+2)D10",
			expected: 29,
			dice:     []dice.Die{{3, 3}, {8, 10}, {7, 10}, {2, 10}, {1, 10}, {6, 10}, {5, 10}},
		},
		{
			input:    "[(2+3)...8]D6",
			expected: 14,
			dice:     []dice.Die{{1, 4}, {1, 6}, {2, 6}, {4, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "[5...(7+1)]D6",
			expected: 14,
			dice:     []dice.Die{{1, 4}, {1, 6}, {2, 6}, {4, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "2d[1...5]",
			expected: 3,
			dice:     []dice.Die{{2, 5}, {1, 2}, {2, 2}},
		},
		{
			input:    "2d([2...4]+2)",
			expected: 7,
			dice:     []dice.Die{{2, 3}, {4, 5}, {3, 5}},
		},
		{
			input:    "2d[(2+3)...8]",
			expected: 10,
			dice:     []dice.Die{{4, 4}, {3, 8}, {7, 8}},
		},
		{
			input:    "2d[5...(7+1)]",
			expected: 10,
			dice:     []dice.Die{{4, 4}, {3, 8}, {7, 8}},
		},
		{
			input:    "([1...4]+1)d([2...4]+2)-1",
			expected: 13,
			dice:     []dice.Die{{2, 4}, {3, 3}, {5, 6}, {5, 6}, {4, 6}},
		},
	}

	for _, test := range testcases {
		name := fmt.Sprintf("%q[%s]",
			test.input, dice.FormatDiceWithoutSpaces(test.dice))
		t.Run(name, func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evaluated, evalErr := evaluator.Eval(ast)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			if evaluated == nil {
				t.Fatalf("Evalの対象外 (nil)")
				return
			}

			// 型が合っているか？
			obj, typeMatched := evaluated.(*object.Integer)
			if !typeMatched {
				t.Fatalf("整数オブジェクトでない: %T (%+v)", obj, obj)
				return
			}

			if obj.Value != test.expected {
				t.Errorf("異なる評価結果: got=%d, want=%d", obj.Value, test.expected)
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
