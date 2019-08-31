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

func TestExecuteRRollComp(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "2r6>=1",
			expected: "DiceBot : (2R6[1]>=1) ＞ 振り足し目標値として2以上の整数を指定してください",
		},
		{
			input:    "2r6=3",
			expected: "DiceBot : (2R6[3]=3) ＞ 3,1 + 2 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2r6=3",
			expected: "DiceBot : (2R6[3]=3) ＞ 6,1 + 3 + 1 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R6<>3",
			expected: "DiceBot : (2R6[3]<>3) ＞ 3,1 + 2 ＞ 成功数2",
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2R6<>3",
			expected: "DiceBot : (2R6[3]<>3) ＞ 6,1 + 3 + 1 ＞ 成功数3",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R6>3",
			expected: "DiceBot : (2R6[3]>3) ＞ 3,1 + 2 ＞ 成功数0",
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2R6>3",
			expected: "DiceBot : (2R6[3]>3) ＞ 6,1 + 3 + 1 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R6<3",
			expected: "DiceBot : (2R6[3]<3) ＞ 3,1 + 2 ＞ 成功数2",
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2R6<3",
			expected: "DiceBot : (2R6[3]<3) ＞ 6,1 + 3 + 1 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R6>=3",
			expected: "DiceBot : (2R6[3]>=3) ＞ 3,1 + 2 ＞ 成功数1",
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2R6>=3",
			expected: "DiceBot : (2R6[3]>=3) ＞ 6,1 + 3 + 1 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R6<=3",
			expected: "DiceBot : (2R6[3]<=3) ＞ 3,1 + 2 ＞ 成功数3",
			dice:     []dice.Die{{3, 6}, {1, 6}, {2, 6}},
		},
		{
			input:    "2R6<=3",
			expected: "DiceBot : (2R6[3]<=3) ＞ 6,1 + 3 + 1 ＞ 成功数3",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R6[3]>=4",
			expected: "DiceBot : (2R6[3]>=4) ＞ 6,1 + 3 + 1 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {1, 6}, {3, 6}, {1, 6}},
		},
		{
			input:    "2R4+2R6>=4",
			expected: "DiceBot : (2R4+2R6[4]>=4) ＞ 4,3 + 3,5 + 1 + 2 ＞ 成功数2",
			dice:     []dice.Die{{4, 4}, {3, 4}, {3, 6}, {5, 6}, {1, 4}, {2, 6}},
		},
	}

	for _, test := range testcases {
		name := fmt.Sprintf(
			"%q[%s]",
			test.input,
			dice.FormatDiceWithoutSpaces(test.dice),
		)
		t.Run(name, func(t *testing.T) {
			root, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			rRollCompNode, rootIsRRollComp := root.(*ast.RRollComp)
			if !rootIsRRollComp {
				t.Fatal("RRollCompではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(rRollCompNode, "DiceBot", evaluator)
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
