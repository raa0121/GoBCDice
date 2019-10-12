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

func TestExecuteURollExpr(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "3u6",
			expected: "DiceBot : (3U6) ＞ 2U6[5] のように振り足し目標値を指定してください",
		},
		{
			input:    "3u6[1]",
			expected: "DiceBot : (3U6[1]) ＞ 振り足し目標値として2以上の整数を指定してください",
		},
		{
			input:    "3u6[6]",
			expected: "DiceBot : (3U6[6]) ＞ 11[6,5],7[6,1],7[6,1] ＞ 11/25 (最大/合計)",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "(1+2)u6[6]",
			expected: "DiceBot : (3U6[6]) ＞ 11[6,5],7[6,1],7[6,1] ＞ 11/25 (最大/合計)",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u(2*3)[6]",
			expected: "DiceBot : (3U6[6]) ＞ 11[6,5],7[6,1],7[6,1] ＞ 11/25 (最大/合計)",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[12/2]",
			expected: "DiceBot : (3U6[6]) ＞ 11[6,5],7[6,1],7[6,1] ＞ 11/25 (最大/合計)",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "6U6[4]",
			expected: "DiceBot : (6U6[4]) ＞ 2,7[5,2],1,11[4,6,1],8[5,3],3 ＞ 11/32 (最大/合計)",
			dice:     []dice.Die{{2, 6}, {5, 6}, {2, 6}, {1, 6}, {4, 6}, {6, 6}, {1, 6}, {5, 6}, {3, 6}, {3, 6}},
		},
		{
			input:    "6U6[4]",
			expected: "DiceBot : (6U6[4]) ＞ 7[4,3],1,1,23[4,6,6,4,3],1,14[6,5,3] ＞ 23/47 (最大/合計)",
			dice:     []dice.Die{{4, 6}, {3, 6}, {1, 6}, {1, 6}, {4, 6}, {6, 6}, {6, 6}, {4, 6}, {3, 6}, {1, 6}, {6, 6}, {5, 6}, {3, 6}},
		},
		{
			input:    "1U100[96]+3",
			expected: "DiceBot : (1U100[96]+3) ＞ 155[98,57]+3 ＞ 158/158 (最大/合計)",
			dice:     []dice.Die{{98, 100}, {57, 100}},
		},
		{
			input:    "3u6[5]+10",
			expected: "DiceBot : (3U6[5]+10) ＞ 1,3,6[5,1]+10 ＞ 16/20 (最大/合計)",
			dice:     []dice.Die{{1, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:    "3u6[5]-(-2-2*4)",
			expected: "DiceBot : (3U6[5]+10) ＞ 1,3,6[5,1]+10 ＞ 16/20 (最大/合計)",
			dice:     []dice.Die{{1, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:    "3u6[5]-10",
			expected: "DiceBot : (3U6[5]-10) ＞ 1,3,6[5,1]-10 ＞ -4/0 (最大/合計)",
			dice:     []dice.Die{{1, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:    "3u6[5]+(-2-2*4)",
			expected: "DiceBot : (3U6[5]-10) ＞ 1,3,6[5,1]-10 ＞ -4/0 (最大/合計)",
			dice:     []dice.Die{{1, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:    "3u6+5u6[6]",
			expected: "DiceBot : (3U6+5U6[6]) ＞ 3,5,3,10[6,4],1,15[6,6,3],5,1 ＞ 15/43 (最大/合計)",
			dice:     []dice.Die{{3, 6}, {5, 6}, {3, 6}, {6, 6}, {4, 6}, {1, 6}, {6, 6}, {6, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:    "(5+6)u10[10]+5",
			expected: "DiceBot : (11U10[10]+5) ＞ 3,2,7,4,7,5,4,4,7,16[10,6],1+5 ＞ 21/65 (最大/合計)",
			dice:     []dice.Die{{3, 10}, {2, 10}, {7, 10}, {4, 10}, {7, 10}, {5, 10}, {4, 10}, {4, 10}, {7, 10}, {10, 10}, {6, 10}, {1, 10}},
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

			uRollExprNode, rootIsURollExpr := root.(*ast.URollExpr)
			if !rootIsURollExpr {
				t.Fatal("URollExprではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(uRollExprNode, "DiceBot", evaluator)
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
