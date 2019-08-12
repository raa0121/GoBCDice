package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"reflect"
	"testing"
)

func TestEvalChoice(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "A",
			dice:     []dice.Die{{1, 4}},
		},
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "B",
			dice:     []dice.Die{{2, 4}},
		},
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "C",
			dice:     []dice.Die{{3, 4}},
		},
		{
			input:    "choice[A,B,C,D]どれにしよう",
			expected: "D",
			dice:     []dice.Die{{4, 4}},
		},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "Call of Cthulhu",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "Sword World",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "Double Cross",
			dice:     []dice.Die{{3, 3}},
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "日本語",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "でも",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "だいじょうぶ",
			dice:     []dice.Die{{3, 3}},
		},
		{
			input:    "choice[1+2, (3*4), 5d6]",
			expected: "1+2",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "choice[1+2, (3*4), 5d6]",
			expected: "(3*4)",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "choice[1+2, (3*4), 5d6]",
			expected: "5d6",
			dice:     []dice.Die{{3, 3}},
		},
	}

	for _, test := range testcases {
		name := fmt.Sprintf("%q[%s]",
			test.input, dice.FormatDiceWithoutSpaces(test.dice))
		t.Run(name, func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evaluated, evalErr := evaluator.Eval(ast)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			if evaluated == nil {
				t.Fatalf("Evalの対象外 (nil)")
				return
			}

			// 型が合っているか？
			obj, typeMatched := evaluated.(*object.String)
			if !typeMatched {
				t.Fatalf("文字列オブジェクトでない: %T (%+v)", obj, obj)
				return
			}

			actual := obj.Value
			if actual != test.expected {
				t.Errorf("異なる値: got=%q, want=%q", actual, test.expected)
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
