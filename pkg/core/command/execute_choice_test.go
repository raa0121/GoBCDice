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

func TestExecuteChoice(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "DiceBot : (CHOICE[A,B,C,D]) ＞ A",
			dice:     []dice.Die{{1, 4}},
		},
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "DiceBot : (CHOICE[A,B,C,D]) ＞ B",
			dice:     []dice.Die{{2, 4}},
		},
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "DiceBot : (CHOICE[A,B,C,D]) ＞ C",
			dice:     []dice.Die{{3, 4}},
		},
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "DiceBot : (CHOICE[A,B,C,D]) ＞ D",
			dice:     []dice.Die{{4, 4}},
		},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "DiceBot : (CHOICE[Call of Cthulhu,Sword World,Double Cross]) ＞ Call of Cthulhu",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "DiceBot : (CHOICE[Call of Cthulhu,Sword World,Double Cross]) ＞ Sword World",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "DiceBot : (CHOICE[Call of Cthulhu,Sword World,Double Cross]) ＞ Double Cross",
			dice:     []dice.Die{{3, 3}},
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "DiceBot : (CHOICE[日本語,でも,だいじょうぶ]) ＞ 日本語",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "DiceBot : (CHOICE[日本語,でも,だいじょうぶ]) ＞ でも",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "DiceBot : (CHOICE[日本語,でも,だいじょうぶ]) ＞ だいじょうぶ",
			dice:     []dice.Die{{3, 3}},
		},
		{
			input:    "choice[1+2, (3*4), 5d6]",
			expected: "DiceBot : (CHOICE[1+2,(3*4),5d6]) ＞ 1+2",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "choice[1+2, (3*4), 5d6]",
			expected: "DiceBot : (CHOICE[1+2,(3*4),5d6]) ＞ (3*4)",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "choice[1+2, (3*4), 5d6]",
			expected: "DiceBot : (CHOICE[1+2,(3*4),5d6]) ＞ 5d6",
			dice:     []dice.Die{{3, 3}},
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

			choiceNode, rootIsChoice := root.(*ast.Choice)
			if !rootIsChoice {
				t.Fatal("Choiceではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(choiceNode, "DiceBot", evaluator)
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
