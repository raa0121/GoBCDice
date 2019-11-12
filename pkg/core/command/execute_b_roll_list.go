package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeBRollList はバラバラロールを実行する。
func executeBRollList(
	node *ast.BRollList,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	// 可変ノードの引数を評価して整数に変換する
	infixNotation, evalVarArgsErr := evalVarArgs(node, evaluator)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	arrayObj := obj.(*object.Array)

	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.AppendMessagePart(notation.Parenthesize(infixNotation))
	result.AppendMessagePart(arrayObj.JoinedElements(","))

	return result, nil
}
