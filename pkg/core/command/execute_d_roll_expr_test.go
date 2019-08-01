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

func TestExecuteDRollExpr(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		{
			input:    "2D6",
			expected: "DiceBot : (2D6) ＞ 8[5,3] ＞ 8",
			dice:     []dice.Die{{5, 6}, {3, 6}},
		},
		{
			input:    "2D4",
			expected: "DiceBot : (2D4) ＞ 3[1,2] ＞ 3",
			dice:     []dice.Die{{1, 4}, {2, 4}},
		},
		{
			input:    "2D6+1",
			expected: "DiceBot : (2D6+1) ＞ 8[2,6]+1 ＞ 9",
			dice:     []dice.Die{{2, 6}, {6, 6}},
		},
		{
			input:    "1+2D6",
			expected: "DiceBot : (1+2D6) ＞ 1+7[4,3] ＞ 8",
			dice:     []dice.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2d6+1-1-2-3-4",
			expected: "DiceBot : (2D6+1-1-2-3-4) ＞ 7[1,6]+1-1-2-3-4 ＞ -2",
			dice:     []dice.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "2D6+4D10",
			expected: "DiceBot : (2D6+4D10) ＞ 9[5,4]+21[1,9,7,4] ＞ 30",
			dice:     []dice.Die{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
		},
		{
			input:    "2d6*3",
			expected: "DiceBot : (2D6*3) ＞ 6[2,4]*3 ＞ 18",
			dice:     []dice.Die{{2, 6}, {4, 6}},
		},
		{
			input:    "2d10+3-4",
			expected: "DiceBot : (2D10+3-4) ＞ 8[3,5]+3-4 ＞ 7",
			dice:     []dice.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d10+3*4",
			expected: "DiceBot : (2D10+3*4) ＞ 8[3,5]+3*4 ＞ 20",
			dice:     []dice.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d6*3-1d6+1",
			expected: "DiceBot : (2D6*3-1D6+1) ＞ 8[6,2]*3-3[3]+1 ＞ 22",
			dice:     []dice.Die{{6, 6}, {2, 6}, {3, 6}},
		},
		{
			input:    "(2+3)d6-1+3d6+2",
			expected: "DiceBot : (5D6-1+3D6+2) ＞ 17[2,3,1,5,6]-1+13[5,4,4]+2 ＞ 31",
			dice:     []dice.Die{{2, 6}, {3, 6}, {1, 6}, {5, 6}, {6, 6}, {5, 6}, {4, 6}, {4, 6}},
		},
		{
			input:    "(2*3-4)d6-1d4+1",
			expected: "DiceBot : (2D6-1D4+1) ＞ 11[6,5]-2[2]+1 ＞ 10",
			dice:     []dice.Die{{6, 6}, {5, 6}, {2, 6}},
		},
		{
			input:    "((2+3)*4/3)d6*2+5",
			expected: "DiceBot : (6D6*2+5) ＞ 24[6,5,6,2,1,4]*2+5 ＞ 53",
			dice:     []dice.Die{{6, 6}, {5, 6}, {6, 6}, {2, 6}, {1, 6}, {4, 6}},
		},
		{
			input:    "1D6/2",
			expected: "DiceBot : (1D6/2) ＞ 1[1]/2 ＞ 0",
			dice:     []dice.Die{{1, 6}},
		},
		{
			input:    "3D6/2",
			expected: "DiceBot : (3D6/2) ＞ 7[1,2,4]/2 ＞ 3",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "3D6/2+1D6",
			expected: "DiceBot : (3D6/2+1D6) ＞ 7[1,2,4]/2+5[5] ＞ 8",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2",
			expected: "DiceBot : (3D6+1D6/2) ＞ 7[1,2,4]+5[5]/2 ＞ 9",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2U",
			expected: "DiceBot : (3D6+1D6/2U) ＞ 7[1,2,4]+5[5]/2U ＞ 10",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "5D6/10",
			expected: "DiceBot : (5D6/10) ＞ 29[6,6,6,6,5]/10 ＞ 2",
			dice:     []dice.Die{{6, 6}, {6, 6}, {6, 6}, {6, 6}, {5, 6}},
		},
		{
			input:    "3D6/2U",
			expected: "DiceBot : (3D6/2U) ＞ 7[1,2,4]/2U ＞ 4",
			dice:     []dice.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "5D6/10u",
			expected: "DiceBot : (5D6/10U) ＞ 21[6,6,6,2,1]/10U ＞ 3",
			dice:     []dice.Die{{6, 6}, {6, 6}, {6, 6}, {2, 6}, {1, 6}},
		},
		{
			input:    "1D100/10R",
			expected: "DiceBot : (1D100/10R) ＞ 55[55]/10R ＞ 6",
			dice:     []dice.Die{{55, 100}},
		},
		{
			input:    "1D100/10r",
			expected: "DiceBot : (1D100/10R) ＞ 54[54]/10R ＞ 5",
			dice:     []dice.Die{{54, 100}},
		},
		{
			input:    "[1...5]D6",
			expected: "DiceBot : (4D6) ＞ 15[5,3,4,3] ＞ 15",
			dice:     []dice.Die{{4, 5}, {5, 6}, {3, 6}, {4, 6}, {3, 6}},
		},
		{
			input:    "([2...4]+2)D10",
			expected: "DiceBot : (6D10) ＞ 29[8,7,2,1,6,5] ＞ 29",
			dice:     []dice.Die{{3, 3}, {8, 10}, {7, 10}, {2, 10}, {1, 10}, {6, 10}, {5, 10}},
		},
		{
			input:    "[(2+3)...8]D6",
			expected: "DiceBot : (5D6) ＞ 14[1,2,4,6,1] ＞ 14",
			dice:     []dice.Die{{1, 4}, {1, 6}, {2, 6}, {4, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "[5...(7+1)]D6",
			expected: "DiceBot : (5D6) ＞ 14[1,2,4,6,1] ＞ 14",
			dice:     []dice.Die{{1, 4}, {1, 6}, {2, 6}, {4, 6}, {6, 6}, {1, 6}},
		},
		{
			input:    "2d[1...5]",
			expected: "DiceBot : (2D2) ＞ 3[1,2] ＞ 3",
			dice:     []dice.Die{{2, 5}, {1, 2}, {2, 2}},
		},
		{
			input:    "2d([2...4]+2)",
			expected: "DiceBot : (2D5) ＞ 7[4,3] ＞ 7",
			dice:     []dice.Die{{2, 3}, {4, 5}, {3, 5}},
		},
		{
			input:    "2d[(2+3)...8]",
			expected: "DiceBot : (2D8) ＞ 10[3,7] ＞ 10",
			dice:     []dice.Die{{4, 4}, {3, 8}, {7, 8}},
		},
		{
			input:    "2d[5...(7+1)]",
			expected: "DiceBot : (2D8) ＞ 10[3,7] ＞ 10",
			dice:     []dice.Die{{4, 4}, {3, 8}, {7, 8}},
		},
		{
			input:    "([1...4]+1)d([2...4]+2)-1",
			expected: "DiceBot : (3D6-1) ＞ 14[5,5,4]-1 ＞ 13",
			dice:     []dice.Die{{2, 4}, {3, 3}, {5, 6}, {5, 6}, {4, 6}},
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
