/*
BCDiceコマンドの評価処理のパッケージ。

主な処理は、以下の3つ。

1. ダイスロールなどの可変ノードの引数を評価し、整数に変換する。

2. 可変ノードの値を決定する。

3. 抽象構文木を評価し、値のオブジェクトに変換する。

*/
package evaluator

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// 評価器の構造体。
type Evaluator struct {
	diceRoller *roller.DiceRoller
	env        *Environment
	// 個数振り足しロールにおける最大振り足し数
	// TODO: 外部から変更するためのインターフェースを作る
	MaxRerolls int
}

// NewEvaluator は新しい評価器を返す。
//
// diceRoller: ダイスローラー,
// env: 評価環境
func NewEvaluator(diceRoller *roller.DiceRoller, env *Environment) *Evaluator {
	return &Evaluator{
		diceRoller: diceRoller,
		env:        env,
		MaxRerolls: 10000,
	}
}

// RolledDice はダイスロール結果を返す。
func (e *Evaluator) RolledDice() []dice.Die {
	return e.env.RolledDice()
}

// Eval はnodeを評価してObjectに変換し、返す。
func (e *Evaluator) Eval(node ast.Node) (object.Object, error) {
	// 型で分岐する
	switch n := node.(type) {
	case *ast.BRollList:
		return e.evalBRollList(n)
	case *ast.BRollComp:
		return e.evalBRollComp(n)
	case *ast.RRollList:
		return e.evalRRollList(n)
	case *ast.RRollComp:
		return e.evalRRollComp(n)
	case *ast.URollExpr:
		return e.evalURollExpr(n)
	case *ast.Choice:
		return e.evalChoice(n)
	case ast.Command:
		return e.Eval(n.Expression())
	case ast.PrefixExpression:
		return e.evalPrefixExpression(n)
	case ast.InfixExpression:
		return e.evalInfixExpression(n)
	case *ast.Int:
		return object.NewInteger(n.Value), nil
	case *ast.SumRollResult:
		return object.NewInteger(n.Value()), nil
	}

	return nil, fmt.Errorf("unknown type: %s", node.Type())
}

// RollDice は、sides個の面を持つダイスをnum個振り、その結果を返す。
// また、ダイスロールの結果を記録する。
func (e *Evaluator) RollDice(num int, sides int) ([]dice.Die, error) {
	rolledDice, err := e.diceRoller.RollDice(num, sides)
	if err != nil {
		return nil, err
	}

	e.env.AppendRolledDice(rolledDice)

	return rolledDice, nil
}
