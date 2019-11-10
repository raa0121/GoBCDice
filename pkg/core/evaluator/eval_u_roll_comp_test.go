package evaluator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
)

func TestEvalURollComp(t *testing.T) {
	testcases := []struct {
		input                  string
		err                    bool
		expectedValueGroups    [][]int
		expectedSumOfGroups    []int
		expectedNumOfSuccesses int
		dice                   []dice.Die
	}{
		{
			input: "3U6>=7",
			err:   true,
		},
		{
			input: "3U6[1]>=7",
			err:   true,
		},
		{
			input:                  "3u6[6]=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups:    []int{11, 7, 7},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:                  "3u6[6]-4=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]<>7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]<>7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups:    []int{11, 7, 7},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:                  "3u6[6]+1<>7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]>6",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]>6",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups:    []int{11, 7, 7},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:                  "3u6[6]-1>6",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]<6",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]<6",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups:    []int{11, 7, 7},
			expectedNumOfSuccesses: 0,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:                  "3u6[6]-2<6",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]>=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]>=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups:    []int{11, 7, 7},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:                  "3u6[6]-1>=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]<=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6[6]<=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups:    []int{11, 7, 7},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:                  "3u6[6]-4<=7",
			err:                    false,
			expectedValueGroups:    [][]int{{6, 5}, {6, 1}, {5}},
			expectedSumOfGroups:    []int{11, 7, 5},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:                  "3u6+5u6[6]>=7",
			err:                    false,
			expectedValueGroups:    [][]int{{3}, {5}, {3}, {6, 4}, {1}, {6, 6, 3}, {5}, {1}},
			expectedSumOfGroups:    []int{3, 5, 3, 10, 1, 15, 5, 1},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{3, 6}, {5, 6}, {3, 6}, {6, 6}, {4, 6}, {1, 6}, {6, 6}, {6, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:                  "(5+6)u10[10]+5>=8",
			err:                    false,
			expectedValueGroups:    [][]int{{3}, {2}, {7}, {4}, {7}, {5}, {4}, {4}, {7}, {10, 6}, {1}},
			expectedSumOfGroups:    []int{3, 2, 7, 4, 7, 5, 4, 4, 7, 16, 1},
			expectedNumOfSuccesses: 9,
			dice:                   []dice.Die{{3, 10}, {2, 10}, {7, 10}, {4, 10}, {7, 10}, {5, 10}, {4, 10}, {4, 10}, {7, 10}, {10, 10}, {6, 10}, {1, 10}},
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
			obj, typeMatched := evaluated.(*object.URollCompResult)
			if !typeMatched {
				t.Fatalf("URollCompResultでない: %T (%+v)", obj, obj)
				return
			}

			valueGroups := obj.RollResult.ValueGroups()
			valueGroupsLength := valueGroups.Length()
			if valueGroupsLength != len(test.expectedValueGroups) {
				t.Fatalf("異なる配列の長さ（回転数）: got=%d, want=%d",
					valueGroupsLength, len(test.expectedValueGroups))
				return
			}

			for i, expectedValues := range test.expectedValueGroups {
				ei := valueGroups.At(i)

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

			sumOfGroups := obj.RollResult.SumOfGroups()
			sumOfGroupsLength := sumOfGroups.Length()
			if sumOfGroupsLength != len(test.expectedSumOfGroups) {
				t.Fatalf("異なる配列の長さ（グループの合計）: got=%d, want=%d",
					sumOfGroupsLength, len(test.expectedSumOfGroups))
				return
			}

			for i, expectedSumOfGroup := range test.expectedSumOfGroups {
				ei := sumOfGroups.At(i)

				t.Run(fmt.Sprintf("SumOfGroup:%d", expectedSumOfGroup), func(t *testing.T) {
					eiInt, eiTypeMatched := ei.(*object.Integer)
					if !eiTypeMatched {
						t.Fatalf("整数オブジェクトでない: %T (%+v)", ei, ei)
						return
					}

					actualSomeOfGroup := eiInt.Value
					if actualSomeOfGroup != expectedSumOfGroup {
						t.Errorf("異なる値: got=%d, want=%d",
							actualSomeOfGroup, expectedSumOfGroup)
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
