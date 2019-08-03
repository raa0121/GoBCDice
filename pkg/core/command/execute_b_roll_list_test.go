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

func TestExecuteBRollList(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "2b6",
			expected: "DiceBot : (2B6) ＞ 3,4",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b6",
			expected: "DiceBot : (2B6) ＞ 6,1",
			dice:     []dice.Die{{6, 6}, {1, 6}},
		},
		{
			input:    "[1...3]b6",
			expected: "DiceBot : (1B6) ＞ 5",
			dice:     []dice.Die{{1, 3}, {5, 6}},
		},
		{
			input:    "[1...3]b6",
			expected: "DiceBot : (3B6) ＞ 2,4,6",
			dice:     []dice.Die{{3, 3}, {2, 6}, {4, 6}, {6, 6}},
		},
		{
			input:    "2b[4...6]",
			expected: "DiceBot : (2B4) ＞ 3,4",
			dice:     []dice.Die{{1, 3}, {3, 4}, {4, 4}},
		},
		{
			input:    "2b[4...6]",
			expected: "DiceBot : (2B5) ＞ 5,1",
			dice:     []dice.Die{{2, 3}, {5, 5}, {1, 5}},
		},
		{
			input:    "[1...3]b[4...6]",
			expected: "DiceBot : (3B6) ＞ 6,4,3",
			dice:     []dice.Die{{3, 3}, {3, 3}, {6, 6}, {4, 6}, {3, 6}},
		},
		{
			input:    "(1*2)b6",
			expected: "DiceBot : (2B6) ＞ 3,4",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "([1...3]+1)b6",
			expected: "DiceBot : (3B6) ＞ 2,4,3",
			dice:     []dice.Die{{2, 3}, {2, 6}, {4, 6}, {3, 6}},
		},
		{
			input:    "2b(2+4)",
			expected: "DiceBot : (2B6) ＞ 3,4",
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2b([3...5]+1)",
			expected: "DiceBot : (2B5) ＞ 1,3",
			dice:     []dice.Die{{2, 3}, {1, 5}, {3, 5}},
		},
		{
			input:    "[1...5]b(2*3)",
			expected: "DiceBot : (5B6) ＞ 3,5,1,5,6",
			dice:     []dice.Die{{5, 5}, {3, 6}, {5, 6}, {1, 6}, {5, 6}, {6, 6}},
		},
		{
			input:    "(1+1)b[1...5]",
			expected: "DiceBot : (2B4) ＞ 3,4",
			dice:     []dice.Die{{4, 5}, {3, 4}, {4, 4}},
		},
		{
			input:    "(1*2)b(2+4)",
			expected: "DiceBot : (2B6) ＞ 4,3",
			dice:     []dice.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2b6+4b10",
			expected: "DiceBot : (2B6+4B10) ＞ 1,3,1,10,3,1",
			dice:     []dice.Die{{1, 6}, {3, 6}, {1, 10}, {10, 10}, {3, 10}, {1, 10}},
		},
		{
			input:    "2b6+4b10",
			expected: "DiceBot : (2B6+4B10) ＞ 3,2,5,7,1,3",
			dice:     []dice.Die{{3, 6}, {2, 6}, {5, 10}, {7, 10}, {1, 10}, {3, 10}},
		},
		{
			input:    "2b6+3b8+5b12",
			expected: "DiceBot : (2B6+3B8+5B12) ＞ 3,4,7,1,5,11,3,4,10,9",
			dice:     []dice.Die{{3, 6}, {4, 6}, {7, 8}, {1, 8}, {5, 8}, {11, 12}, {3, 12}, {4, 12}, {10, 12}, {9, 12}},
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

			bRollListNode, rootIsBRollList := root.(*ast.BRollList)
			if !rootIsBRollList {
				t.Fatal("BRollListではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(bRollListNode, "DiceBot", evaluator)
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
