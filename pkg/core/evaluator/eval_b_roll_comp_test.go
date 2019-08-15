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

func TestEvalBRollComp(t *testing.T) {
	testcases := []struct {
		input                  string
		expectedValues         []int
		expectedNumOfSuccesses int
		dice                   []dice.Die
	}{
		{
			input:                  "2b6=3",
			expectedValues:         []int{3, 4},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                  "2b6=3",
			expectedValues:         []int{2, 1},
			expectedNumOfSuccesses: 0,
			dice:                   []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:                  "2b6<>3",
			expectedValues:         []int{3, 4},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                  "2b6<>3",
			expectedValues:         []int{2, 1},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:                  "2b6>3",
			expectedValues:         []int{3, 4},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                  "2b6>3",
			expectedValues:         []int{2, 1},
			expectedNumOfSuccesses: 0,
			dice:                   []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:                  "2b6<3",
			expectedValues:         []int{3, 4},
			expectedNumOfSuccesses: 0,
			dice:                   []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                  "2b6<3",
			expectedValues:         []int{2, 1},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:                  "2b6>=3",
			expectedValues:         []int{3, 4},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                  "2b6>=3",
			expectedValues:         []int{2, 1},
			expectedNumOfSuccesses: 0,
			dice:                   []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:                  "2b6<=3",
			expectedValues:         []int{3, 4},
			expectedNumOfSuccesses: 1,
			dice:                   []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                  "2b6<=3",
			expectedValues:         []int{2, 1},
			expectedNumOfSuccesses: 2,
			dice:                   []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:                  "2b6+4b10>3",
			expectedValues:         []int{3, 4, 9, 7, 1, 5},
			expectedNumOfSuccesses: 4,
			dice:                   []dice.Die{{3, 6}, {4, 6}, {9, 10}, {7, 10}, {1, 10}, {5, 10}},
		},
		{
			input:                  "2b6+3b8+5b12<5",
			expectedValues:         []int{5, 2, 7, 3, 4, 11, 9, 8, 12, 6},
			expectedNumOfSuccesses: 3,
			dice:                   []dice.Die{{5, 6}, {2, 6}, {7, 8}, {3, 8}, {4, 8}, {11, 12}, {9, 12}, {8, 12}, {12, 12}, {6, 12}},
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
			obj, typeMatched := evaluated.(*object.BRollCompResult)
			if !typeMatched {
				t.Fatalf("BRollCompResultでない: %T (%+v)", obj, obj)
				return
			}

			elements := obj.Values.Elements

			if len(elements) != len(test.expectedValues) {
				t.Fatalf("異なる値の配列の長さ: got=%d, want=%d",
					len(elements), len(test.expectedValues))
				return
			}

			for i, e := range test.expectedValues {
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
