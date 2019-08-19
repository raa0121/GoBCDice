package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"testing"
)

// 可変ノードの値を決定する例。
func ExampleEvaluator_DetermineValues() {
	r, parseErr := parser.Parse("ExampleEvaluator_DetermineValues", []byte("2d6*3-1d6+1"))
	if parseErr != nil {
		return
	}

	node := r.(ast.Node)

	fmt.Println("構文解析直後の抽象構文木: " + node.SExp())

	// 可変ノードの値を決定する
	dieFeeder := feeder.NewQueue([]dice.Die{{6, 6}, {2, 6}, {3, 6}})
	evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

	err := evaluator.DetermineValues(node)
	if err != nil {
		return
	}

	fmt.Println("ダイスロール結果: " + dice.FormatDice(evaluator.RolledDice()))
	fmt.Println("値決定後の抽象構文木: " + node.SExp())
	// Output:
	// 構文解析直後の抽象構文木: (DRollExpr (+ (- (* (DRoll 2 6) 3) (DRoll 1 6)) 1))
	// ダイスロール結果: 6/6, 2/6, 3/6
	// 値決定後の抽象構文木: (DRollExpr (+ (- (* (SumRollResult (Die 6 6) (Die 2 6)) 3) (SumRollResult (Die 3 6))) 1))
}

func TestDetermineValues(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "2D6",
			expected: "(DRollExpr (SumRollResult (Die 5 6) (Die 3 6)))",
			dice:     []dice.Die{{5, 6}, {3, 6}},
		},
		{
			input:    "2D4",
			expected: "(DRollExpr (SumRollResult (Die 1 4) (Die 2 4)))",
			dice:     []dice.Die{{1, 4}, {2, 4}},
		},
		{
			input:    "2D6+1",
			expected: "(DRollExpr (+ (SumRollResult (Die 2 6) (Die 6 6)) 1))",
			dice:     []dice.Die{{2, 6}, {6, 6}},
		},
		{
			input:    "1+2D6",
			expected: "(DRollExpr (+ 1 (SumRollResult (Die 4 6) (Die 3 6))))",
			dice:     []dice.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2d6+1-1-2-3-4",
			expected: "(DRollExpr (- (- (- (- (+ (SumRollResult (Die 1 6) (Die 6 6)) 1) 1) 2) 3) 4))",
			dice:     []dice.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "2D6+4D10",
			expected: "(DRollExpr (+ (SumRollResult (Die 5 6) (Die 4 6)) (SumRollResult (Die 1 10) (Die 9 10) (Die 7 10) (Die 4 10))))",
			dice:     []dice.Die{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
		},
		{
			input:    "2d6*3",
			expected: "(DRollExpr (* (SumRollResult (Die 2 6) (Die 4 6)) 3))",
			dice:     []dice.Die{{2, 6}, {4, 6}},
		},
		{
			input:    "2d10+3-4",
			expected: "(DRollExpr (- (+ (SumRollResult (Die 3 10) (Die 5 10)) 3) 4))",
			dice:     []dice.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d10+3*4",
			expected: "(DRollExpr (+ (SumRollResult (Die 3 10) (Die 5 10)) (* 3 4)))",
			dice:     []dice.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d6*3-1d6+1",
			expected: "(DRollExpr (+ (- (* (SumRollResult (Die 6 6) (Die 2 6)) 3) (SumRollResult (Die 3 6))) 1))",
			dice:     []dice.Die{{6, 6}, {2, 6}, {3, 6}},
		},
		{
			input:    "1D6/2",
			expected: "(DRollExpr (/ (SumRollResult (Die 1 6)) 2))",
			dice:     []dice.Die{{1, 6}},
		},
		{
			input:    "3D6/2",
			expected: "(DRollExpr (/ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) 2))",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "3D6/2+1D6",
			expected: "(DRollExpr (+ (/ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) 2) (SumRollResult (Die 5 6))))",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2",
			expected: "(DRollExpr (+ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) (/ (SumRollResult (Die 5 6)) 2)))",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2U",
			expected: "(DRollExpr (+ (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) (/U (SumRollResult (Die 5 6)) 2)))",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "5D6/10",
			expected: "(DRollExpr (/ (SumRollResult (Die 6 6) (Die 6 6) (Die 6 6) (Die 6 6) (Die 5 6)) 10))",
			dice:     []dice.Die{{6, 6}, {6, 6}, {6, 6}, {6, 6}, {5, 6}},
		},
		{
			input:    "3D6/2U",
			expected: "(DRollExpr (/U (SumRollResult (Die 1 6) (Die 2 6) (Die 4 6)) 2))",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "5D6/10u",
			expected: "(DRollExpr (/U (SumRollResult (Die 6 6) (Die 6 6) (Die 6 6) (Die 2 6) (Die 1 6)) 10))",
			dice:     []dice.Die{{6, 6}, {6, 6}, {6, 6}, {2, 6}, {1, 6}},
		},
		{
			input:    "1D100/10R",
			expected: "(DRollExpr (/R (SumRollResult (Die 55 100)) 10))",
			dice:     []dice.Die{{55, 100}},
		},
		{
			input:    "1D100/10r",
			expected: "(DRollExpr (/R (SumRollResult (Die 54 100)) 10))",
			dice:     []dice.Die{{54, 100}},
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(ast.Node)

			// 可変ノードの値を決定する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			err := evaluator.DetermineValues(node)
			if err != nil {
				t.Fatalf("評価エラー: %s", err)
				return
			}

			actual := node.SExp()
			if actual != test.expected {
				t.Errorf("異なる評価結果: got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
