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

func TestEvalRRollList(t *testing.T) {
	testcases := []struct {
		input    string
		err      bool
		expected [][]int
		dice     []dice.Die
	}{
		{
			input: "2R6",
			err:   true,
		},
		{
			input:    "2R6[3]",
			err:      false,
			expected: [][]int{{3, 1}, {2}},
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2R6[3]",
			err:      false,
			expected: [][]int{{6, 1}, {3}, {1}},
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{4, 1, 5, 4, 3, 2, 5, 1, 6, 3, 2, 5}, {5}},
			dice:     []dice.Die{{4, 6}, {1, 6}, {5, 6}, {4, 6}, {3, 6}, {2, 6}, {5, 6}, {1, 6}, {6, 6}, {3, 6}, {2, 6}, {5, 6}, {5, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{6, 1, 2, 2, 2, 1, 5, 5, 1, 5, 2, 6}, {4, 5}},
			dice:     []dice.Die{{6, 6}, {1, 6}, {2, 6}, {2, 6}, {2, 6}, {1, 6}, {5, 6}, {5, 6}, {1, 6}, {5, 6}, {2, 6}, {6, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{2, 4, 4, 5, 4, 1, 2, 6, 6, 6, 1, 4}, {6, 4, 6}, {3, 4}},
			dice:     []dice.Die{{2, 6}, {4, 6}, {4, 6}, {5, 6}, {4, 6}, {1, 6}, {2, 6}, {6, 6}, {6, 6}, {6, 6}, {1, 6}, {4, 6}, {6, 6}, {4, 6}, {6, 6}, {3, 6}, {4, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{2, 5, 4, 5, 3, 1, 3, 4, 2, 2, 4, 5}},
			dice:     []dice.Die{{2, 6}, {5, 6}, {4, 6}, {5, 6}, {3, 6}, {1, 6}, {3, 6}, {4, 6}, {2, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{5, 3, 5, 3, 1, 1, 6, 3, 2, 2, 5, 5}, {6}, {4}},
			dice:     []dice.Die{{5, 6}, {3, 6}, {5, 6}, {3, 6}, {1, 6}, {1, 6}, {6, 6}, {3, 6}, {2, 6}, {2, 6}, {5, 6}, {5, 6}, {6, 6}, {4, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{4, 4, 2, 1, 5, 6, 5, 1, 1, 4, 4, 6}, {3, 6}, {2}},
			dice:     []dice.Die{{4, 6}, {4, 6}, {2, 6}, {1, 6}, {5, 6}, {6, 6}, {5, 6}, {1, 6}, {1, 6}, {4, 6}, {4, 6}, {6, 6}, {3, 6}, {6, 6}, {2, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{2, 1, 2, 5, 2, 3, 1, 3, 4, 4, 4, 5}},
			dice:     []dice.Die{{2, 6}, {1, 6}, {2, 6}, {5, 6}, {2, 6}, {3, 6}, {1, 6}, {3, 6}, {4, 6}, {4, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{1, 2, 3, 4, 1, 1, 2, 3, 6, 4, 1, 4}, {2}},
			dice:     []dice.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {1, 6}, {1, 6}, {2, 6}, {3, 6}, {6, 6}, {4, 6}, {1, 6}, {4, 6}, {2, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{4, 3, 6, 2, 2, 1, 3, 5, 1, 3, 2, 1}, {6}, {2}},
			dice:     []dice.Die{{4, 6}, {3, 6}, {6, 6}, {2, 6}, {2, 6}, {1, 6}, {3, 6}, {5, 6}, {1, 6}, {3, 6}, {2, 6}, {1, 6}, {6, 6}, {2, 6}},
		},
		{
			input:    "12R6[6]",
			err:      false,
			expected: [][]int{{6, 6, 3, 1, 4, 2, 4, 5, 4, 3, 6, 6}, {4, 4, 3, 4}},
			dice:     []dice.Die{{6, 6}, {6, 6}, {3, 6}, {1, 6}, {4, 6}, {2, 6}, {4, 6}, {5, 6}, {4, 6}, {3, 6}, {6, 6}, {6, 6}, {4, 6}, {4, 6}, {3, 6}, {4, 6}},
		},
		{
			input:    "2R4+2R6[4]",
			err:      false,
			expected: [][]int{{4, 3}, {3, 5}, {1}, {2}},
			dice:     []dice.Die{{4, 4}, {3, 4}, {3, 6}, {5, 6}, {1, 4}, {2, 6}},
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
				if !test.err {
					t.Fatalf("評価エラー: %s", evalErr)
					return
				}

				return
			}

			if test.err {
				t.Fatal("評価エラーが発生しなかった")
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

			valueGroupsLength := obj.Length()
			if valueGroupsLength != len(test.expected) {
				t.Fatalf("異なる配列の長さ（回転数）: got=%d, want=%d", valueGroupsLength, len(test.expected))
				return
			}

			for i, expectedValues := range test.expected {
				ei := obj.At(i)

				t.Run(fmt.Sprintf("%v", expectedValues), func(t *testing.T) {
					eiArray, eiTypeMatched := ei.(*object.Array)
					if !eiTypeMatched {
						t.Fatalf("配列オブジェクトでない: %T (%+v)", ei, ei)
						return
					}

					valuesLength := eiArray.Length()
					if valuesLength != len(expectedValues) {
						t.Fatalf("異なる配列の長さ（振られたダイスの数）: got=%d, want=%d",
							valuesLength, len(expectedValues))
						return
					}

					for j, e := range expectedValues {
						ej := eiArray.At(j)

						t.Run(fmt.Sprintf("%d", e), func(t *testing.T) {
							ejInt, ejTypeMatched := ej.(*object.Integer)
							if !ejTypeMatched {
								t.Fatalf("整数オブジェクトでない: %T (%+v)", ej, ej)
								return
							}

							x := ejInt.Value
							if x != e {
								t.Errorf("異なる値: got=%d, want=%d", x, e)
							}
						})
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
