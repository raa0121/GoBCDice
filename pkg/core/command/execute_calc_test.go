package command

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
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
			root, parseErr := parser.Parse("test", []byte(test.input))
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

			expectedSuccessCheckResult := SUCCESS_CHECK_UNSPECIFIED
			actualSuccessCheckResult := r.SuccessCheckResult
			if actualSuccessCheckResult != expectedSuccessCheckResult {
				t.Errorf("成功判定結果が異なる: got %s, want %s",
					actualSuccessCheckResult, expectedSuccessCheckResult)
			}
		})
	}
}
