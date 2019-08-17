package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"testing"
)

func TestSetRerollThreshold(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"3r6<=4", "(RRollComp (<= (RRollList 4 (RRoll 3 6)) 4))"},
		{"3r6+2r4<=2", "(RRollComp (<= (RRollList 2 (RRoll 3 6) (RRoll 2 4)) 2))"},
		{"(3+2)r6>=5", "(RRollComp (>= (RRollList 5 (RRoll (+ 3 2) 6)) 5))"},
		{"1r(2*3)>=4", "(RRollComp (>= (RRollList 4 (RRoll 1 (* 2 3))) 4))"},
		{"3r6>1*4", "(RRollComp (> (RRollList 4 (RRoll 3 6)) (* 1 4)))"},
		{"6R6[6]>=5", "(RRollComp (>= (RRollList 6 (RRoll 6 6)) 5))"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(*ast.RRollComp)

			// 可変ノードの引数を評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evalErr := evaluator.SetRerollThreshold(node)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			actual := node.SExp()
			if actual != test.expected {
				t.Errorf("異なる評価結果: got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
