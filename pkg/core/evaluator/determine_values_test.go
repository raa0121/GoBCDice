package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"github.com/raa0121/GoBCDice/pkg/core/die/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/die/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"testing"
)

func TestDetermineValues(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []die.Die
	}{
		{
			input:    "2D6",
			expected: "(DRollExpr (SumRollResult (Die 5 6) (Die 3 6)))",
			dice:     []die.Die{{5, 6}, {3, 6}},
		},
		{
			input:    "2D4",
			expected: "(DRollExpr (SumRollResult (Die 1 4) (Die 2 4)))",
			dice:     []die.Die{{1, 4}, {2, 4}},
		},
		{
			input:    "2D6+1",
			expected: "(DRollExpr (+ (SumRollResult (Die 2 6) (Die 6 6)) 1))",
			dice:     []die.Die{{2, 6}, {6, 6}},
		},
		{
			input:    "1+2D6",
			expected: "(DRollExpr (+ 1 (SumRollResult (Die 4 6) (Die 3 6))))",
			dice:     []die.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2d6+1-1-2-3-4",
			expected: "(DRollExpr (- (- (- (- (+ (SumRollResult (Die 1 6) (Die 6 6)) 1) 1) 2) 3) 4))",
			dice:     []die.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "2D6+4D10",
			expected: "(DRollExpr (+ (SumRollResult (Die 5 6) (Die 4 6)) (SumRollResult (Die 1 10) (Die 9 10) (Die 7 10) (Die 4 10))))",
			dice:     []die.Die{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
		},
		{
			input:    "2d6*3",
			expected: "(DRollExpr (* (SumRollResult (Die 2 6) (Die 4 6)) 3))",
			dice:     []die.Die{{2, 6}, {4, 6}},
		},
		{
			input:    "2d10+3-4",
			expected: "(DRollExpr (- (+ (SumRollResult (Die 3 10) (Die 5 10)) 3) 4))",
			dice:     []die.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d10+3*4",
			expected: "(DRollExpr (+ (SumRollResult (Die 3 10) (Die 5 10)) (* 3 4)))",
			dice:     []die.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d6*3-1d6+1",
			expected: "(DRollExpr (+ (- (* (SumRollResult (Die 6 6) (Die 2 6)) 3) (SumRollResult (Die 3 6))) 1))",
			dice:     []die.Die{{6, 6}, {2, 6}, {3, 6}},
		},
		{
			input:    "1D6/2",
			expected: "(DRollExpr (/ (SumRollResult (Die 1 6)) 2))",
			dice:     []die.Die{{1, 6}},
		},
		{
			input:    "3D6/2",
			expected: "(DRollExpr (/ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) 2))",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "3D6/2+1D6",
			expected: "(DRollExpr (+ (/ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) 2) (SumRollResult (Die 5 6))))",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2",
			expected: "(DRollExpr (+ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) (/ (SumRollResult (Die 5 6)) 2)))",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2U",
			expected: "(DRollExpr (+ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) (/U (SumRollResult (Die 5 6)) 2)))",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "5D6/10",
			expected: "(DRollExpr (/ (SumRollResult (Die 6 6) (Die 6 6) (Die 6 6) (Die 6 6) (Die 5 6)) 10))",
			dice:     []die.Die{{6, 6}, {6, 6}, {6, 6}, {6, 6}, {5, 6}},
		},
		{
			input:    "3D6/2U",
			expected: "(DRollExpr (/U (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) 2))",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "5D6/10u",
			expected: "(DRollExpr (/U (SumRollResult (Die 6 6) (Die 6 6) (Die 6 6) (Die 2 6) (Die 1 6)) 10))",
			dice:     []die.Die{{6, 6}, {6, 6}, {6, 6}, {2, 6}, {1, 6}},
		},
		{
			input:    "1D100/10R",
			expected: "(DRollExpr (/R (SumRollResult (Die 55 100)) 10))",
			dice:     []die.Die{{55, 100}},
		},
		{
			input:    "1D100/10r",
			expected: "(DRollExpr (/R (SumRollResult (Die 54 100)) 10))",
			dice:     []die.Die{{54, 100}},
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			// 可変ノードの値を決定する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			err := evaluator.DetermineValues(ast)
			if err != nil {
				t.Fatalf("評価エラー: %s", err)
				return
			}

			actual := ast.SExp()
			if actual != test.expected {
				t.Errorf("異なる評価結果: got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
