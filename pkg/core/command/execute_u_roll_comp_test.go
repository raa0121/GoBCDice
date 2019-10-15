package command

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
)

func TestExecuteURollComp(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "3U6>=7",
			expected: "DiceBot : (3U6>=7) ＞ 2U6[5] のように振り足し目標値を指定してください",
		},
		{
			input:    "3U6[1]>=7",
			expected: "DiceBot : (3U6[1]>=7) ＞ 振り足し目標値として2以上の整数を指定してください",
		},
		{
			input:    "3u6[6]=7",
			expected: "DiceBot : (3U6[6]=7) ＞ 11[6,5],7[6,1],5 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]=7",
			expected: "DiceBot : (3U6[6]=7) ＞ 11[6,5],7[6,1],7[6,1] ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[6]-4=7",
			expected: "DiceBot : (3U6[6]-4=7) ＞ 11[6,5],7[6,1],5-4 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]<>7",
			expected: "DiceBot : (3U6[6]<>7) ＞ 11[6,5],7[6,1],5 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]<>7",
			expected: "DiceBot : (3U6[6]<>7) ＞ 11[6,5],7[6,1],7[6,1] ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[6]+1<>7",
			expected: "DiceBot : (3U6[6]+1<>7) ＞ 11[6,5],7[6,1],5+1 ＞ 成功数3",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]>6",
			expected: "DiceBot : (3U6[6]>6) ＞ 11[6,5],7[6,1],5 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]>6",
			expected: "DiceBot : (3U6[6]>6) ＞ 11[6,5],7[6,1],7[6,1] ＞ 成功数3",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[6]-1>6",
			expected: "DiceBot : (3U6[6]-1>6) ＞ 11[6,5],7[6,1],5-1 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]<6",
			expected: "DiceBot : (3U6[6]<6) ＞ 11[6,5],7[6,1],5 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]<6",
			expected: "DiceBot : (3U6[6]<6) ＞ 11[6,5],7[6,1],7[6,1] ＞ 成功数0",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[6]-2<6",
			expected: "DiceBot : (3U6[6]-2<6) ＞ 11[6,5],7[6,1],5-2 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]>=7",
			expected: "DiceBot : (3U6[6]>=7) ＞ 11[6,5],7[6,1],5 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]>=7",
			expected: "DiceBot : (3U6[6]>=7) ＞ 11[6,5],7[6,1],7[6,1] ＞ 成功数3",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[6]-1>=7",
			expected: "DiceBot : (3U6[6]-1>=7) ＞ 11[6,5],7[6,1],5-1 ＞ 成功数1",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]<=7",
			expected: "DiceBot : (3U6[6]<=7) ＞ 11[6,5],7[6,1],5 ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6[6]<=7",
			expected: "DiceBot : (3U6[6]<=7) ＞ 11[6,5],7[6,1],7[6,1] ＞ 成功数2",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "3u6[6]-4<=7",
			expected: "DiceBot : (3U6[6]-4<=7) ＞ 11[6,5],7[6,1],5-4 ＞ 成功数3",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {1, 6}, {5, 6}},
		},
		{
			input:    "3u6+5u6[6]>=7",
			expected: "DiceBot : (3U6+5U6[6]>=7) ＞ 3,5,3,10[6,4],1,15[6,6,3],5,1 ＞ 成功数2",
			dice:     []dice.Die{{3, 6}, {5, 6}, {3, 6}, {6, 6}, {4, 6}, {1, 6}, {6, 6}, {6, 6}, {3, 6}, {5, 6}, {1, 6}},
		},
		{
			input:    "(5+6)u10[10]+5>=8",
			expected: "DiceBot : (11U10[10]+5>=8) ＞ 3,2,7,4,7,5,4,4,7,16[10,6],1+5 ＞ 成功数9",
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

			if root.(ast.Node).Type() != ast.U_ROLL_COMP_NODE {
				t.Fatal("URollCompではない")
			}
			uRollCompNode := root.(*ast.Command)

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(uRollCompNode, "DiceBot", evaluator)
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
