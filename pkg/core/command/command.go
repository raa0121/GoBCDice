/*
BCDiceコマンドの実行処理のパッケージ。
構文解析で得たコマンドのノードを評価して、最終的な出力のメッセージを生成することができる。

このパッケージにおいて、コマンドの種類ごとに実行の仕方を定義する。
*/
package command

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
)

// Execute は指定されたコマンドを実行する。
//
// node: コマンドのノード,
// gameID: ゲーム識別子,
// evaluator: 評価器。
func Execute(
	node ast.Node,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	switch c := node.(type) {
	case *ast.Calc:
		return executeCalc(c, gameID, evaluator)
	case *ast.DRollExpr:
		return executeDRollExpr(c, gameID, evaluator)
	case *ast.DRollComp:
		return executeDRollComp(c, gameID, evaluator)
	case *ast.BRollList:
		return executeBRollList(c, gameID, evaluator)
	case *ast.BRollComp:
		return executeBRollComp(c, gameID, evaluator)
	case *ast.RRollList:
		return executeRRollList(c, gameID, evaluator)
	case *ast.RRollComp:
		return executeRRollComp(c, gameID, evaluator)
	case *ast.URollExpr:
		return executeURollExpr(c, gameID, evaluator)
	case *ast.Choice:
		return executeChoice(c, gameID, evaluator)
	}

	return nil, fmt.Errorf("command execution not implemented: %s", node.Type())
}

// evalVarArgs は、加算ロールなどの可変ノードの引数を評価して整数に変換する。
// 返り値はその結果の中置表記とエラー。
func evalVarArgs(node ast.Node, evaluator *evaluator.Evaluator) (string, error) {
	evalErr := evaluator.EvalVarArgs(node)
	if evalErr != nil {
		return "", evalErr
	}

	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return "", infixNotationErr
	}

	return infixNotation, nil
}
