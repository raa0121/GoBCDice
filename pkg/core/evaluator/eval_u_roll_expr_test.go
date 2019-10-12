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

func TestEvalURollExpr(t *testing.T) {
	testcases := []struct {
		input               string
		err                 bool
		expectedValueGroups [][]int
		expectedSumOfGroups []int
		expectedMaxValue    int
		expectedSumOfValues int
		dice                []dice.Die
	}{
		{
			input: "3U6",
			err:   true,
		},
		{
			input:               "3u6[6]",
			err:                 false,
			expectedValueGroups: [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups: []int{11, 7, 7},
			expectedMaxValue:    11,
			expectedSumOfValues: 25,
			dice:                []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:               "(1+2)u6[6]",
			err:                 false,
			expectedValueGroups: [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups: []int{11, 7, 7},
			expectedMaxValue:    11,
			expectedSumOfValues: 25,
			dice:                []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:               "3u(2*3)[6]",
			err:                 false,
			expectedValueGroups: [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups: []int{11, 7, 7},
			expectedMaxValue:    11,
			expectedSumOfValues: 25,
			dice:                []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:               "3u6[12/2]",
			err:                 false,
			expectedValueGroups: [][]int{{6, 5}, {6, 1}, {6, 1}},
			expectedSumOfGroups: []int{11, 7, 7},
			expectedMaxValue:    11,
			expectedSumOfValues: 25,
			dice:                []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:               "6U6[4]",
			err:                 false,
			expectedValueGroups: [][]int{{2}, {5, 2}, {1}, {4, 6, 1}, {5, 3}, {3}},
			expectedSumOfGroups: []int{2, 7, 1, 11, 8, 3},
			expectedMaxValue:    11,
			expectedSumOfValues: 32,
			dice:                []dice.Die{{2, 6}, {5, 6}, {2, 6}, {1, 6}, {4, 6}, {6, 6}, {1, 6}, {5, 6}, {3, 6}, {3, 6}},
		},
		{
			input:               "6U6[4]",
			err:                 false,
			expectedValueGroups: [][]int{{4, 3}, {1}, {1}, {4, 6, 6, 4, 3}, {1}, {6, 5, 3}},
			expectedSumOfGroups: []int{7, 1, 1, 23, 1, 14},
			expectedMaxValue:    23,
			expectedSumOfValues: 47,
			dice:                []dice.Die{{4, 6}, {3, 6}, {1, 6}, {1, 6}, {4, 6}, {6, 6}, {6, 6}, {4, 6}, {3, 6}, {1, 6}, {6, 6}, {5, 6}, {3, 6}},
		},
		{
			input:               "1U100[96]+3",
			err:                 false,
			expectedValueGroups: [][]int{{98, 57}},
			expectedSumOfGroups: []int{155},
			expectedMaxValue:    158,
			expectedSumOfValues: 158,
			dice:                []dice.Die{{98, 100}, {57, 100}},
		},
		{
			input:               "3u6[5]+10",
			err:                 false,
			expectedValueGroups: [][]int{{1}, {3}, {5, 1}},
			expectedSumOfGroups: []int{1, 3, 6},
			expectedMaxValue:    16,
			expectedSumOfValues: 20,
			dice:                []dice.Die{{1, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:               "3u6[5]-10",
			err:                 false,
			expectedValueGroups: [][]int{{1}, {3}, {5, 1}},
			expectedSumOfGroups: []int{1, 3, 6},
			expectedMaxValue:    -4,
			expectedSumOfValues: 0,
			dice:                []dice.Die{{1, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:               "3u6+5u6[6]",
			err:                 false,
			expectedValueGroups: [][]int{{3}, {5}, {3}, {6, 4}, {1}, {6, 6, 3}, {5}, {1}},
			expectedSumOfGroups: []int{3, 5, 3, 10, 1, 15, 5, 1},
			expectedMaxValue:    15,
			expectedSumOfValues: 43,
			dice:                []dice.Die{{3, 6}, {5, 6}, {3, 6}, {6, 6}, {4, 6}, {1, 6}, {6, 6}, {6, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:               "(5+6)u10[10]+5",
			err:                 false,
			expectedValueGroups: [][]int{{3}, {2}, {7}, {4}, {7}, {5}, {4}, {4}, {7}, {10, 6}, {1}},
			expectedSumOfGroups: []int{3, 2, 7, 4, 7, 5, 4, 4, 7, 16, 1},
			expectedMaxValue:    21,
			expectedSumOfValues: 65,
			dice:                []dice.Die{{3, 10}, {2, 10}, {7, 10}, {4, 10}, {7, 10}, {5, 10}, {4, 10}, {4, 10}, {7, 10}, {10, 10}, {6, 10}, {1, 10}},
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
			obj, typeMatched := evaluated.(*object.URollExprResult)
			if !typeMatched {
				t.Fatalf("URollExprResultでない: %T (%+v)", obj, obj)
				return
			}

			valueGroupsLength := obj.ValueGroups().Length()
			if valueGroupsLength != len(test.expectedValueGroups) {
				t.Fatalf("異なる配列の長さ（回転数）: got=%d, want=%d",
					valueGroupsLength, len(test.expectedValueGroups))
				return
			}

			for i, expectedValues := range test.expectedValueGroups {
				ei := obj.ValueGroups().At(i)

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

			sumOfGroupsLength := obj.SumOfGroups().Length()
			if sumOfGroupsLength != len(test.expectedSumOfGroups) {
				t.Fatalf("異なる配列の長さ（グループの合計）: got=%d, want=%d",
					sumOfGroupsLength, len(test.expectedSumOfGroups))
				return
			}

			for i, expectedSumOfGroup := range test.expectedSumOfGroups {
				ei := obj.SumOfGroups().At(i)

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

			actualMaxValue := obj.MaxValue().Value
			if actualMaxValue != test.expectedMaxValue {
				t.Errorf("異なる最大値: got=%d, want=%d",
					actualMaxValue, test.expectedMaxValue)
			}

			actualSumOfValues := obj.SumOfValues().Value
			if actualSumOfValues != test.expectedSumOfValues {
				t.Errorf("異なる合計値: got=%d, want=%d",
					actualSumOfValues, test.expectedSumOfValues)
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
