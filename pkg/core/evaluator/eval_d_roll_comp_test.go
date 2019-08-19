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

func TestEvalDRollComp(t *testing.T) {
	testcases := []struct {
		input    string
		expected bool
		dice     []dice.Die
	}{
		{
			input:    "2D6=7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2D6=7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:    "2D6=7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:    "2D6<>7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2D6<>7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:    "2D6<>7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:    "2D6<7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:    "2D6<7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2D6>7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:    "2D6>7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2D6<=7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2D6<=7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:    "2D6>=7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "2D6>=7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:    "-2D6<-7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:    "-2D6<-7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "-2D6>-7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:    "-2D6>-7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "-2D6<=-7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "-2D6<=-7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:    "-2D6>=-7",
			expected: true,
			dice:     []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:    "-2D6>=-7",
			expected: false,
			dice:     []dice.Die{{3, 6}, {5, 6}},
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
			obj, typeMatched := evaluated.(*object.Boolean)
			if !typeMatched {
				t.Fatalf("論理値オブジェクトでない: %T (%+v)", obj, obj)
				return
			}

			if obj.Value != test.expected {
				t.Errorf("異なる評価結果: got=%t, want=%t", obj.Value, test.expected)
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
