package command

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeURollExpr は上方無限ロール式を実行する。
func executeURollExpr(
	node *ast.URollExpr,
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
	checkRerollThresholdErr := evaluator.CheckURollThreshold(node.URollList)
	if checkRerollThresholdErr != nil {
		result.AppendMessagePart(checkRerollThresholdErr.Error())
		return result, nil
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	uRollExprResult := obj.(*object.URollExprResult)
	result.RolledDice = evaluator.RolledDice()

	result.AppendMessagePart(formatURollExprValueGroupsAndModifier(uRollExprResult))
	result.AppendMessagePart(fmt.Sprintf(
		"%d/%d (最大/合計)",
		uRollExprResult.MaxValue().Value,
		uRollExprResult.SumOfValues().Value,
	))

	return result, nil
}

// formatURollExprValueGroupsAndModifier は上方無限ロールの出目および修正値を整形する。
func formatURollExprValueGroupsAndModifier(result *object.URollExprResult) string {
	valueGroups := result.ValueGroups()
	n := valueGroups.Length()
	sumOfGroups := result.SumOfGroups()
	formattedValueGroups := make([]string, 0, n)

	for i := 0; i < n; i++ {
		var buf bytes.Buffer
		valuesArray := valueGroups.At(i).(*object.Array)

		buf.WriteString(sumOfGroups.At(i).Inspect())

		if valuesArray.Length() > 1 {
			buf.WriteString(valuesArray.InspectWithoutSpaces())
		}

		formattedValueGroups = append(formattedValueGroups, buf.String())
	}

	var buf bytes.Buffer
	buf.WriteString(strings.Join(formattedValueGroups, ","))
	buf.WriteString(notation.FormatModifier(result.Modifier().Value))

	return buf.String()
}
