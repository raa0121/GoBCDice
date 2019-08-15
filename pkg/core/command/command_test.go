package command

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
)

// 加算ロールコマンドの例。
func ExampleExecute_sumRoll() {
	// 構文解析する
	root, parseErr := parser.Parse("ExampleExecute_sumRoll", []byte("(2*3-4)d6-1d4+1"))
	if parseErr != nil {
		return
	}

	// コマンドと認識されたことを確認する
	commandNode, rootIsCommand := root.(ast.Command)
	if !rootIsCommand {
		return
	}

	// コマンドを実行する
	dieFeeder := feeder.NewQueue([]dice.Die{{6, 6}, {5, 6}, {2, 4}})
	evaluator := evaluator.NewEvaluator(
		roller.New(dieFeeder),
		evaluator.NewEnvironment(),
	)

	result, execErr := Execute(commandNode, "DiceBot", evaluator)
	if execErr != nil {
		return
	}

	// 結果のメッセージを出力する
	fmt.Println(result.Message())
	// Output: DiceBot : (2D6-1D4+1) ＞ 11[6,5]-2[2]+1 ＞ 10
}
