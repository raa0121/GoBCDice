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

func TestExecuteDRollComp(t *testing.T) {
	testcases := []struct {
		input                      string
		expectedMessage            string
		expectedSuccessCheckResult SuccessCheckResultType
		dice                       []dice.Die
	}{
		{
			input:                      "2D6=7",
			expectedMessage:            "DiceBot : (2D6=7) ＞ 7[3,4] ＞ 7 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6=7",
			expectedMessage:            "DiceBot : (2D6=7) ＞ 6[3,3] ＞ 6 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "2D6=7",
			expectedMessage:            "DiceBot : (2D6=7) ＞ 6[3,3] ＞ 6 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "2D6+1=3+4",
			expectedMessage:            "DiceBot : (2D6+1=7) ＞ 6[2,4]+1 ＞ 7 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{2, 6}, {4, 6}},
		},
		{
			input:                      "2D6+1=3+4",
			expectedMessage:            "DiceBot : (2D6+1=7) ＞ 7[3,4]+1 ＞ 8 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6<>7",
			expectedMessage:            "DiceBot : (2D6<>7) ＞ 7[3,4] ＞ 7 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6<>7",
			expectedMessage:            "DiceBot : (2D6<>7) ＞ 6[3,3] ＞ 6 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "2D6<>7",
			expectedMessage:            "DiceBot : (2D6<>7) ＞ 8[3,5] ＞ 8 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:                      "2D6<7",
			expectedMessage:            "DiceBot : (2D6<7) ＞ 6[3,3] ＞ 6 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "2D6<7",
			expectedMessage:            "DiceBot : (2D6<7) ＞ 7[3,4] ＞ 7 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6>7",
			expectedMessage:            "DiceBot : (2D6>7) ＞ 8[3,5] ＞ 8 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:                      "2D6>7",
			expectedMessage:            "DiceBot : (2D6>7) ＞ 7[3,4] ＞ 7 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6<=7",
			expectedMessage:            "DiceBot : (2D6<=7) ＞ 7[3,4] ＞ 7 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6<=7",
			expectedMessage:            "DiceBot : (2D6<=7) ＞ 8[3,5] ＞ 8 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:                      "2D6>=7",
			expectedMessage:            "DiceBot : (2D6>=7) ＞ 7[3,4] ＞ 7 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "2D6>=7",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			expectedMessage:            "DiceBot : (2D6>=7) ＞ 6[3,3] ＞ 6 ＞ 失敗",
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "-2D6<-7",
			expectedMessage:            "DiceBot : (-2D6<-7) ＞ -8[3,5] ＞ -8 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {5, 6}},
		},
		{
			input:                      "-2D6<-7",
			expectedMessage:            "DiceBot : (-2D6<-7) ＞ -7[3,4] ＞ -7 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "-2D6>-7",
			expectedMessage:            "DiceBot : (-2D6>-7) ＞ -6[3,3] ＞ -6 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "-2D6>-7",
			expectedMessage:            "DiceBot : (-2D6>-7) ＞ -7[3,4] ＞ -7 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "-2D6<=-7",
			expectedMessage:            "DiceBot : (-2D6<=-7) ＞ -7[3,4] ＞ -7 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "-2D6<=-7",
			expectedMessage:            "DiceBot : (-2D6<=-7) ＞ -6[3,3] ＞ -6 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {3, 6}},
		},
		{
			input:                      "-2D6>=-7",
			expectedMessage:            "DiceBot : (-2D6>=-7) ＞ -7[3,4] ＞ -7 ＞ 成功",
			expectedSuccessCheckResult: SUCCESS_CHECK_SUCCESS,
			dice:                       []dice.Die{{3, 6}, {4, 6}},
		},
		{
			input:                      "-2D6>=-7",
			expectedMessage:            "DiceBot : (-2D6>=-7) ＞ -8[3,5] ＞ -8 ＞ 失敗",
			expectedSuccessCheckResult: SUCCESS_CHECK_FAILURE,
			dice:                       []dice.Die{{3, 6}, {5, 6}},
		},
	}

	for _, test := range testcases {
		name := fmt.Sprintf(
			"%q[%s]",
			test.input,
			dice.FormatDiceWithoutSpaces(test.dice),
		)
		t.Run(name, func(t *testing.T) {
			root, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			if root.(ast.Node).Type() != ast.D_ROLL_COMP_NODE {
				t.Fatal("DRollCompではない")
			}
			dRollCompNode := root.(*ast.Command)

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := evaluator.NewEvaluator(
				roller.New(dieFeeder),
				evaluator.NewEnvironment(),
			)

			r, execErr := Execute(dRollCompNode, "DiceBot", evaluator)
			if execErr != nil {
				t.Fatalf("コマンド実行エラー: %s", execErr)
				return
			}

			actualMessage := r.Message()
			if actualMessage != test.expectedMessage {
				t.Errorf("結果のメッセージが異なる: got %q, want %q",
					actualMessage, test.expectedMessage)
			}

			if !reflect.DeepEqual(r.RolledDice, test.dice) {
				t.Errorf("ダイスロール結果が異なる: got [%s], want [%s]",
					dice.FormatDice(r.RolledDice), dice.FormatDice(test.dice))
			}

			actualSuccessCheckResult := r.SuccessCheckResult
			if actualSuccessCheckResult != test.expectedSuccessCheckResult {
				t.Errorf("成功判定結果が異なる: got %q, want %q",
					actualSuccessCheckResult, test.expectedSuccessCheckResult)
			}
		})
	}
}
