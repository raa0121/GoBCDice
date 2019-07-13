package command

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/die"
	"github.com/raa0121/GoBCDice/internal/die/feeder"
	"github.com/raa0121/GoBCDice/internal/die/roller"
	"github.com/raa0121/GoBCDice/internal/evaluator"
	"github.com/raa0121/GoBCDice/internal/parser"
	"testing"
)

func TestExecuteCalc(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"C(5)", "DiceBot : C(5) ＞ 計算結果 ＞ 5"},
		{"C(10)", "DiceBot : C(10) ＞ 計算結果 ＞ 10"},
		{"C(42)", "DiceBot : C(42) ＞ 計算結果 ＞ 42"},
		{"C(65535)", "DiceBot : C(65535) ＞ 計算結果 ＞ 65535"},
		{"C(-5)", "DiceBot : C(-5) ＞ 計算結果 ＞ -5"},
		{"C(-10)", "DiceBot : C(-10) ＞ 計算結果 ＞ -10"},
		{"C(-42)", "DiceBot : C(-42) ＞ 計算結果 ＞ -42"},
		{"C(-65535)", "DiceBot : C(-65535) ＞ 計算結果 ＞ -65535"},
		{"C(+5)", "DiceBot : C(5) ＞ 計算結果 ＞ 5"},
		{"C(+10)", "DiceBot : C(10) ＞ 計算結果 ＞ 10"},
		{"C(+42)", "DiceBot : C(42) ＞ 計算結果 ＞ 42"},
		{"C(+65535)", "DiceBot : C(65535) ＞ 計算結果 ＞ 65535"},
		{"C(1+2)", "DiceBot : C(1+2) ＞ 計算結果 ＞ 3"},
		{"C(1-2)", "DiceBot : C(1-2) ＞ 計算結果 ＞ -1"},
		{"C(1*2)", "DiceBot : C(1*2) ＞ 計算結果 ＞ 2"},
		{"C(1/2)", "DiceBot : C(1/2) ＞ 計算結果 ＞ 0"},
		{"C(-1+2)", "DiceBot : C(-1+2) ＞ 計算結果 ＞ 1"},
		{"C(+1+2)", "DiceBot : C(1+2) ＞ 計算結果 ＞ 3"},
		{"C(1+2-3)", "DiceBot : C(1+2-3) ＞ 計算結果 ＞ 0"},
		{"C(1*2+3)", "DiceBot : C(1*2+3) ＞ 計算結果 ＞ 5"},
		{"C(1/2+3)", "DiceBot : C(1/2+3) ＞ 計算結果 ＞ 3"},
		{"C(1+2*3)", "DiceBot : C(1+2*3) ＞ 計算結果 ＞ 7"},
		{"C(1+2/3)", "DiceBot : C(1+2/3) ＞ 計算結果 ＞ 1"},
		{"C(1+(2-3))", "DiceBot : C(1+2-3) ＞ 計算結果 ＞ 0"},
		{"C((1+2)*3)", "DiceBot : C((1+2)*3) ＞ 計算結果 ＞ 9"},
		{"C((1+2)/3)", "DiceBot : C((1+2)/3) ＞ 計算結果 ＞ 1"},
		{"C(-(1+2))", "DiceBot : C(-(1+2)) ＞ 計算結果 ＞ -3"},
		{"C(+(1+2))", "DiceBot : C(1+2) ＞ 計算結果 ＞ 3"},
		{"C(1+50/3-2)", "DiceBot : C(1+50/3-2) ＞ 計算結果 ＞ 15"},
		{"C(1+50/3u-2)", "DiceBot : C(1+50/3U-2) ＞ 計算結果 ＞ 16"},
		{"C(1+50/3r-2)", "DiceBot : C(1+50/3R-2) ＞ 計算結果 ＞ 16"},
		{"C(1+100/3-2)", "DiceBot : C(1+100/3-2) ＞ 計算結果 ＞ 32"},
		{"C(1+100/3u-2)", "DiceBot : C(1+100/3U-2) ＞ 計算結果 ＞ 33"},
		{"C(1+100/3r-2)", "DiceBot : C(1+100/3R-2) ＞ 計算結果 ＞ 32"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			root, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			calcNode, rootIsCalc := root.(*ast.Calc)
			if !rootIsCalc {
				t.Fatal("Calcではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(calcNode, "DiceBot", evaluator)
			if execErr != nil {
				t.Fatalf("コマンド実行エラー: %s", execErr)
				return
			}

			actual := r.Message()
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}

func TestExecuteDRollExpr(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []die.Die
	}{
		{
			input:    "2D6",
			expected: "DiceBot : (2D6) ＞ 8[5,3] ＞ 8",
			dice:     []die.Die{{5, 6}, {3, 6}},
		},
		{
			input:    "2D4",
			expected: "DiceBot : (2D4) ＞ 3[1,2] ＞ 3",
			dice:     []die.Die{{1, 4}, {2, 4}},
		},
		{
			input:    "2D6+1",
			expected: "DiceBot : (2D6+1) ＞ 8[2,6]+1 ＞ 9",
			dice:     []die.Die{{2, 6}, {6, 6}},
		},
		{
			input:    "1+2D6",
			expected: "DiceBot : (1+2D6) ＞ 1+7[4,3] ＞ 8",
			dice:     []die.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2d6+1-1-2-3-4",
			expected: "DiceBot : (2D6+1-1-2-3-4) ＞ 7[1,6]+1-1-2-3-4 ＞ -2",
			dice:     []die.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "2D6+4D10",
			expected: "DiceBot : (2D6+4D10) ＞ 9[5,4]+21[1,9,7,4] ＞ 30",
			dice:     []die.Die{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
		},
		{
			input:    "2d6*3",
			expected: "DiceBot : (2D6*3) ＞ 6[2,4]*3 ＞ 18",
			dice:     []die.Die{{2, 6}, {4, 6}},
		},
		{
			input:    "2d10+3-4",
			expected: "DiceBot : (2D10+3-4) ＞ 8[3,5]+3-4 ＞ 7",
			dice:     []die.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d10+3*4",
			expected: "DiceBot : (2D10+3*4) ＞ 8[3,5]+3*4 ＞ 20",
			dice:     []die.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d6*3-1d6+1",
			expected: "DiceBot : (2D6*3-1D6+1) ＞ 8[6,2]*3-3[3]+1 ＞ 22",
			dice:     []die.Die{{6, 6}, {2, 6}, {3, 6}},
		},
		{
			input:    "(2+3)d6-1+3d6+2",
			expected: "DiceBot : (5D6-1+3D6+2) ＞ 17[2,3,1,5,6]-1+13[5,4,4]+2 ＞ 31",
			dice:     []die.Die{{2, 6}, {3, 6}, {1, 6}, {5, 6}, {6, 6}, {5, 6}, {4, 6}, {4, 6}},
		},
		{
			input:    "(2*3-4)d6-1d4+1",
			expected: "DiceBot : (2D6-1D4+1) ＞ 11[6,5]-2[2]+1 ＞ 10",
			dice:     []die.Die{{6, 6}, {5, 6}, {2, 6}},
		},
		{
			input:    "((2+3)*4/3)d6*2+5",
			expected: "DiceBot : (6D6*2+5) ＞ 24[6,5,6,2,1,4]*2+5 ＞ 53",
			dice:     []die.Die{{6, 6}, {5, 6}, {6, 6}, {2, 6}, {1, 6}, {4, 6}},
		},
		{
			input:    "1D6/2",
			expected: "DiceBot : (1D6/2) ＞ 1[1]/2 ＞ 0",
			dice:     []die.Die{{1, 6}},
		},
		{
			input:    "3D6/2",
			expected: "DiceBot : (3D6/2) ＞ 7[1,2,4]/2 ＞ 3",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "3D6/2+1D6",
			expected: "DiceBot : (3D6/2+1D6) ＞ 7[1,2,4]/2+5[5] ＞ 8",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2",
			expected: "DiceBot : (3D6+1D6/2) ＞ 7[1,2,4]+5[5]/2 ＞ 9",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2U",
			expected: "DiceBot : (3D6+1D6/2U) ＞ 7[1,2,4]+5[5]/2U ＞ 10",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "5D6/10",
			expected: "DiceBot : (5D6/10) ＞ 29[6,6,6,6,5]/10 ＞ 2",
			dice:     []die.Die{{6, 6}, {6, 6}, {6, 6}, {6, 6}, {5, 6}},
		},
		{
			input:    "3D6/2U",
			expected: "DiceBot : (3D6/2U) ＞ 7[1,2,4]/2U ＞ 4",
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "5D6/10u",
			expected: "DiceBot : (5D6/10U) ＞ 21[6,6,6,2,1]/10U ＞ 3",
			dice:     []die.Die{{6, 6}, {6, 6}, {6, 6}, {2, 6}, {1, 6}},
		},
		{
			input:    "1D100/10R",
			expected: "DiceBot : (1D100/10R) ＞ 55[55]/10R ＞ 6",
			dice:     []die.Die{{55, 100}},
		},
		{
			input:    "1D100/10r",
			expected: "DiceBot : (1D100/10R) ＞ 54[54]/10R ＞ 5",
			dice:     []die.Die{{54, 100}},
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			root, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			dRollExprNode, rootIsDRollExpr := root.(*ast.DRollExpr)
			if !rootIsDRollExpr {
				t.Fatal("DRollExprではない")
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(dRollExprNode, "DiceBot", evaluator)
			if execErr != nil {
				t.Fatalf("コマンド実行エラー: %s", execErr)
				return
			}

			actual := r.Message()
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}
