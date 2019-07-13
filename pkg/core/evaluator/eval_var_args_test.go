package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/die/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/die/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"testing"
)

func TestEvalVarArgs(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"(1+2)d6", "(DRollExpr (DRoll 3 6))"},
		{"4d(3*2)", "(DRollExpr (DRoll 4 6))"},
		{"(8/2)D(4+6)", "(DRollExpr (DRoll 4 10))"},
		{"-(1+2)d6", "(DRollExpr (- (DRoll 3 6)))"},
		{"(2+3)d6-1+3d6+2", "(DRollExpr (+ (+ (- (DRoll 5 6) 1) (DRoll 3 6)) 2))"},
		{"(2*3-4)d6-1d4+1", "(DRollExpr (+ (- (DRoll 2 6) (DRoll 1 4)) 1))"},
		{"((2+3)*4/3)d6*2+5", "(DRollExpr (+ (* (DRoll 6 6) 2) 5))"},
		{"(2-1)d(8/2)*(1+1)d(3*4/2)+2*3", "(DRollExpr (+ (* (DRoll 1 4) (DRoll 2 6)) (* 2 3)))"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			// 可変ノードの引数を評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evalErr := evaluator.EvalVarArgs(ast)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			actual := ast.SExp()
			if actual != test.expected {
				t.Errorf("異なる評価結果: got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
