package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"reflect"
	"testing"
)

func TestEvalBRollList(t *testing.T) {
	testcases := []struct {
		input    string
		expected []int
		dice     []dice.Die
	}{
		{
			input:    "2b6",
			expected: []int{3, 4},
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6",
			expected: []int{1, 6},
			dice:     []dice.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "(1*2)b(2+4)",
			expected: []int{3, 4},
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "[1...3]b[4...6]",
			expected: []int{3, 2, 1},
			dice:     []dice.Die{{3, 3}, {1, 3}, {3, 4}, {2, 4}, {1, 4}},
		},
		{
			input:    "2b6+4b10",
			expected: []int{3, 4, 9, 7, 1, 5},
			dice:     []dice.Die{{3, 6}, {4, 6}, {9, 10}, {7, 10}, {1, 10}, {5, 10}},
		},
		{
			input:    "2b6+3b8+5b12",
			expected: []int{5, 2, 7, 3, 4, 11, 9, 8, 12, 6},
			dice:     []dice.Die{{5, 6}, {2, 6}, {7, 8}, {3, 8}, {4, 8}, {11, 12}, {9, 12}, {8, 12}, {12, 12}, {6, 12}},
		},
	}

	for _, test := range testcases {
		name := fmt.Sprintf("%q[%s]",
			test.input, dice.FormatDiceWithoutSpaces(test.dice))
		t.Run(name, func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(ast.Node)

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evaluated, evalErr := evaluator.Eval(node)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			if evaluated == nil {
				t.Fatalf("Evalの対象外 (nil)")
				return
			}

			// 型が合っているか？
			obj, typeMatched := evaluated.(*object.Array)
			if !typeMatched {
				t.Fatalf("配列オブジェクトでない: %T (%+v)", obj, obj)
				return
			}

			elements := obj.Elements

			if len(elements) != len(test.expected) {
				t.Fatalf("異なる配列の長さ: got=%d, want=%d", len(elements), len(test.expected))
				return
			}

			for i, e := range test.expected {
				ei := elements[i]

				t.Run(fmt.Sprintf("%d", e), func(t *testing.T) {
					eiInt, ok := ei.(*object.Integer)
					if !ok {
						t.Fatalf("整数オブジェクトでない: %T (%+v)", ei, ei)
						return
					}

					x := eiInt.Value
					if x != e {
						t.Errorf("異なる値: got=%d, want=%d", x, e)
					}
				})
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
