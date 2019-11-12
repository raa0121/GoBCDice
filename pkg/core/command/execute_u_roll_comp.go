package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeURollComp は上方無限ロールの成功数カウントを実行する。
func executeURollComp(
	node *ast.Command,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	compareNode := node.Expression.(*ast.BasicInfixExpression)

	// 左辺の可変ノードの引数および右辺を評価する
	evalVarArgsErr := evaluator.EvalVarArgs(compareNode)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 中置表記を生成する
	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return nil, infixNotationErr
	}

	result.AppendMessagePart(notation.Parenthesize(infixNotation))

	// 振り足しの閾値を確認する
	checkRerollThresholdErr := evaluator.CheckURollThreshold(
		compareNode.Left().(*ast.URollExpr).URollList,
	)
	if checkRerollThresholdErr != nil {
		result.AppendMessagePart(checkRerollThresholdErr.Error())
		return result, nil
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	resultObj := obj.(*object.URollCompResult)
	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.AppendMessagePart(
		formatURollExprValueGroupsAndModifier(resultObj.RollResult),
	)
	result.AppendMessagePart("成功数" + resultObj.NumOfSuccesses.Inspect())

	return result, nil
}
