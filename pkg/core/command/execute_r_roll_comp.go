package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeRRollComp は個数振り足しロールの成功数カウントを実行する。
func executeRRollComp(
	node *ast.Command,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	compareNode := node.Expression.(*ast.Compare)

	// 左辺の可変ノードの引数および右辺を評価する
	evalVarArgsErr := evaluator.EvalVarArgs(compareNode)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 振り足しの閾値を設定する
	setRerollThresholdErr := evaluator.SetRerollThreshold(node)
	if setRerollThresholdErr != nil {
		return nil, setRerollThresholdErr
	}

	// 中置表記を生成する
	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return nil, infixNotationErr
	}

	result.appendMessagePart(notation.Parenthesize(infixNotation))

	// 振り足しの閾値を確認する
	checkRerollThresholdErr :=
		evaluator.CheckRRollThreshold(compareNode.Left().(*ast.RRollList))
	if checkRerollThresholdErr != nil {
		result.appendMessagePart(checkRerollThresholdErr.Error())
		return result, nil
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	resultObj := obj.(*object.RRollCompResult)
	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.appendMessagePart(formatRRollValues(resultObj.ValueGroups))
	result.appendMessagePart("成功数" + resultObj.NumOfSuccesses.Inspect())

	return result, nil
}
