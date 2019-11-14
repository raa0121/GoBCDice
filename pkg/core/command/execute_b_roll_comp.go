package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeBRollComp はバラバラロールの成功数カウントを実行する。
func executeBRollComp(
	node *ast.Command,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	// 左辺の可変ノードの引数および右辺を評価する
	infixNotation, evalVarArgsErr := evalVarArgs(
		node.Expression.(*ast.BasicInfixExpression),
		evaluator,
	)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	resultObj := obj.(*object.BRollCompResult)
	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.AppendMessagePart(notation.Parenthesize(infixNotation))
	result.AppendMessagePart(resultObj.Values.JoinedElements(","))
	result.AppendMessagePart("成功数" + resultObj.NumOfSuccesses.Inspect())

	return result, nil
}
