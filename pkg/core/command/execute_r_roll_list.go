package command

import (
	"strings"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeRRollList は個数振り足しロールを実行する。
func executeRRollList(
	node *ast.RRollList,
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

	result.AppendMessagePart(notation.Parenthesize(infixNotation))

	// 振り足しの閾値を確認する
	checkRerollThresholdErr := evaluator.CheckRRollThreshold(node)
	if checkRerollThresholdErr != nil {
		result.AppendMessagePart(checkRerollThresholdErr.Error())
		return result, nil
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	valueGroups := obj.(*object.Array)
	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.AppendMessagePart(formatRRollValues(valueGroups))

	return result, nil
}

// formatRRollValues は個数振り足しロールの出目を整形する。
func formatRRollValues(valueGroups *object.Array) string {
	valueGroupStrs := make([]string, 0, valueGroups.Length())
	for _, valuesObj := range valueGroups.Elements {
		valuesArray := valuesObj.(*object.Array)
		valueGroupStrs = append(valueGroupStrs, valuesArray.JoinedElements(","))
	}

	return strings.Join(valueGroupStrs, " + ")
}
