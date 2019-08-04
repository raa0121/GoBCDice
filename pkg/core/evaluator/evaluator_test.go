package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
)

// 抽象構文木を評価し、値のオブジェクトに変換する例。
func ExampleEvaluator_Eval() {
	// 構文解析する
	ast, parseErr := parser.Parse("(2*3-4)d6-1d4+1")
	if parseErr != nil {
		return
	}

	fmt.Println("抽象構文木: " + ast.SExp())

	// ノードを評価する
	dieFeeder := feeder.NewQueue([]dice.Die{{6, 6}, {2, 6}, {3, 4}})
	evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

	obj, evalErr := evaluator.Eval(ast)
	if evalErr != nil {
		return
	}

	fmt.Println("ダイスロール結果: " + dice.FormatDice(evaluator.RolledDice()))
	fmt.Println("評価結果: " + obj.Inspect())
	// Output:
	// 抽象構文木: (DRollExpr (+ (- (DRoll (- (* 2 3) 4) 6) (DRoll 1 4)) 1))
	// ダイスロール結果: 6/6, 2/6, 3/4
	// 評価結果: 6
}
