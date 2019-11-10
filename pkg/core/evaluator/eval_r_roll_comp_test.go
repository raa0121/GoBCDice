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

func TestEvalRRollComp(t *testing.T) {
	testcases := []struct {
		input                  string
		expectedValueGroups    [][]int
		expectedNumOfSuccesses int
		dice                   []dice.Die
	}{
		{
			input:                  "2R6[3]=3",
			expectedValueGroups:    [][]int{{3, 1}, {2}},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:                  "2R6[3]=3",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R6[3]<>3",
			expectedValueGroups:    [][]int{{3, 1}, {2}},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:                  "2R6[3]<>3",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R6[3]>3",
			expectedValueGroups:    [][]int{{3, 1}, {2}},
			expectedNumOfSuccesses: 0,
			dice:                   []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:                  "2R6[3]>3",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R6[3]<3",
			expectedValueGroups:    [][]int{{3, 1}, {2}},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:                  "2R6[3]<3",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R6[3]>=3",
			expectedValueGroups:    [][]int{{3, 1}, {2}},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:                  "2R6[3]>=3",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R6[3]<=3",
			expectedValueGroups:    [][]int{{3, 1}, {2}},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:                  "2R6[3]<=3",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R6[3]>=4",
			expectedValueGroups:    [][]int{{6, 1}, {3}, {1}},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:                  "2R4+2R6[4]>=4",
			expectedValueGroups:    [][]int{{4, 3}, {3, 5}, {1}, {2}},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{4, 4}, {3, 4}, {3, 6}, {5, 6}, {1, 4}, {2, 6}},
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
			obj, typeMatched := evaluated.(*object.RRollCompResult)
			if !typeMatched {
				t.Fatalf("RRollCompResultでない: %T (%+v)", obj, obj)
				return
			}

			valueGroupsLength := obj.ValueGroups.Length()
			if valueGroupsLength != len(test.expectedValueGroups) {
				t.Fatalf("異なる配列の長さ（回転数）: got=%d, want=%d",
					valueGroupsLength, len(test.expectedValueGroups))
				return
			}

			for i, expectedValues := range test.expectedValueGroups {
				ei := obj.ValueGroups.At(i)

				t.Run(fmt.Sprintf("%v", expectedValues), func(t *testing.T) {
					eiArray, eiTypeMatched := ei.(*object.Array)
					if !eiTypeMatched {
						t.Fatalf("配列オブジェクトでない: %T (%+v)", ei, ei)
						return
					}

					eiLength := eiArray.Length()
					if eiLength != len(expectedValues) {
						t.Fatalf("異なる配列の長さ（振られたダイスの数）: got=%d, want=%d",
							eiLength, len(expectedValues))
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

			actualNumOfSuccesses := obj.NumOfSuccesses.Value
			if actualNumOfSuccesses != test.expectedNumOfSuccesses {
				t.Errorf("異なる成功数: got=%d, want=%d",
					actualNumOfSuccesses, test.expectedNumOfSuccesses)
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
