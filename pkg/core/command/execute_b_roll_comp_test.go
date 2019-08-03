package command

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"reflect"
	"testing"
)

func TestExecuteBRollComp(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "2b6=3",
			expected: "DiceBot : (2B6=3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6=3",
			expected: "DiceBot : (2B6=3) ＞ 2,1 ＞ 成功数0",
			dice:     []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:    "2b6<>3",
			expected: "DiceBot : (2B6<>3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6<>3",
			expected: "DiceBot : (2B6<>3) ＞ 2,1 ＞ 成功数2",
			dice:     []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:    "2b6>3",
			expected: "DiceBot : (2B6>3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6>3",
			expected: "DiceBot : (2B6>3) ＞ 2,1 ＞ 成功数0",
			dice:     []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:    "2b6<3",
			expected: "DiceBot : (2B6<3) ＞ 3,4 ＞ 成功数0",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6<3",
			expected: "DiceBot : (2B6<3) ＞ 2,1 ＞ 成功数2",
			dice:     []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:    "2b6>=3",
			expected: "DiceBot : (2B6>=3) ＞ 3,4 ＞ 成功数2",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6>=3",
			expected: "DiceBot : (2B6>=3) ＞ 2,1 ＞ 成功数0",
			dice:     []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:    "2b6<=3",
			expected: "DiceBot : (2B6<=3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6<=3",
			expected: "DiceBot : (2B6<=3) ＞ 2,1 ＞ 成功数2",
			dice:     []dice.Die{{2, 6}, {1, 6}},
		},
		{
			input:    "2b6+4b10>3",
			expected: "DiceBot : (2B6+4B10>3) ＞ 3,4,9,7,1,5 ＞ 成功数4",
			dice:     []dice.Die{{3, 6}, {4, 6}, {9, 10}, {7, 10}, {1, 10}, {5, 10}},
		},
		{
			input:    "2b6+4b10>3",
			expected: "DiceBot : (2B6+4B10>3) ＞ 3,2,5,7,1,3 ＞ 成功数2",
			dice:     []dice.Die{{3, 6}, {2, 6}, {5, 10}, {7, 10}, {1, 10}, {3, 10}},
		},
		{
			input:    "2b6+3b8+5b12<5",
			expected: "DiceBot : (2B6+3B8+5B12<5) ＞ 5,2,7,3,4,11,9,8,12,6 ＞ 成功数3",
			dice:     []dice.Die{{5, 6}, {2, 6}, {7, 8}, {3, 8}, {4, 8}, {11, 12}, {9, 12}, {8, 12}, {12, 12}, {6, 12}},
		},
		{
			input:    "2b6+3b8+5b12<5",
			expected: "DiceBot : (2B6+3B8+5B12<5) ＞ 3,4,7,1,5,11,3,4,10,9 ＞ 成功数5",
			dice:     []dice.Die{{3, 6}, {4, 6}, {7, 8}, {1, 8}, {5, 8}, {11, 12}, {3, 12}, {4, 12}, {10, 12}, {9, 12}},
		},
		{
			input:    "[1...3]b6>3",
			expected: "DiceBot : (1B6>3) ＞ 5 ＞ 成功数1",
			dice:     []dice.Die{{1, 3}, {5, 6}},
		},
		{
			input:    "[1...3]b6>3",
			expected: "DiceBot : (3B6>3) ＞ 2,4,6 ＞ 成功数2",
			dice:     []dice.Die{{3, 3}, {2, 6}, {4, 6}, {6, 6}},
		},
		{
			input:    "2b[4...6]>3",
			expected: "DiceBot : (2B4>3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{1, 3}, {3, 4}, {4, 4}},
		},
		{
			input:    "2b[4...6]>3",
			expected: "DiceBot : (2B5>3) ＞ 5,1 ＞ 成功数1",
			dice:     []dice.Die{{2, 3}, {5, 5}, {1, 5}},
		},
		{
			input:    "[1...3]b[4...6]>3",
			expected: "DiceBot : (3B6>3) ＞ 6,4,3 ＞ 成功数2",
			dice:     []dice.Die{{3, 3}, {3, 3}, {6, 6}, {4, 6}, {3, 6}},
		},
		{
			input:    "(1*2)b6>3",
			expected: "DiceBot : (2B6>3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "([1...3]+1)b6>3",
			expected: "DiceBot : (3B6>3) ＞ 2,4,3 ＞ 成功数1",
			dice:     []dice.Die{{2, 3}, {2, 6}, {4, 6}, {3, 6}},
		},
		{
			input:    "2b(2+4)>3",
			expected: "DiceBot : (2B6>3) ＞ 3,4 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b([3...5]+1)>3",
			expected: "DiceBot : (2B5>3) ＞ 1,3 ＞ 成功数0",
			dice:     []dice.Die{{2, 3}, {1, 5}, {3, 5}},
		},
		{
			input:    "(1*2)b(2+4)>3",
			expected: "DiceBot : (2B6>3) ＞ 4,3 ＞ 成功数1",
			dice:     []dice.Die{{4, 6}, {3, 6}},
		},
	}

	for _, test := range testcases {
		name := fmt.Sprintf(
			"%q[%s]",
			test.input,
			dice.FormatDiceWithoutSpaces(test.dice),
		)
		t.Run(name, func(t *testing.T) {
			root, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			bRollCompNode, rootIsBRollComp := root.(*ast.BRollComp)
			if !rootIsBRollComp {
				t.Fatal("BRollCompではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(bRollCompNode, "DiceBot", evaluator)
			if execErr != nil {
				t.Fatalf("コマンド実行エラー: %s", execErr)
				return
			}

			actualMessage := r.Message()
			if actualMessage != test.expected {
				t.Errorf("結果のメッセージが異なる: got %q, want %q", actualMessage, test.expected)
			}

			if !reflect.DeepEqual(r.RolledDice, test.dice) {
				t.Errorf("ダイスロール結果が異なる: got [%s], want [%s]",
					dice.FormatDice(r.RolledDice), dice.FormatDice(test.dice))
			}

			expectedSuccessCheckResult := SUCCESS_CHECK_UNSPECIFIED
			actualSuccessCheckResult := r.SuccessCheckResult
			if actualSuccessCheckResult != expectedSuccessCheckResult {
				t.Errorf("成功判定結果が異なる: got %s, want %s",
					actualSuccessCheckResult, expectedSuccessCheckResult)
			}
		})
	}
}
