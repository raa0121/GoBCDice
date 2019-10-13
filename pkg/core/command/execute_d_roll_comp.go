package command

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// executeDRollComp は加算ロール式の成功判定を実行する。
func executeDRollComp(
	node *ast.Command,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	compareNode, exprIsCompareNode := node.Expression.(*ast.Compare)
	if !exprIsCompareNode {
		return nil, fmt.Errorf("DRollComp: expression is not a Compare node: %s", node.Type())
	}

	// 左辺の可変ノードの引数および右辺を評価する
	infixNotation1, evalVarArgsErr := evalVarArgs(compareNode, evaluator)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 加算ロールなどの可変ノードの値を決定する
	infixNotation2, determineValuesErr :=
		determineCompareValues(compareNode, evaluator)
	if determineValuesErr != nil {
		return nil, determineValuesErr
	}

	// 左辺を評価する
	leftObj, leftEvalErr := evaluator.EvalCompareLeft(compareNode)
	if leftEvalErr != nil {
		return nil, leftEvalErr
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(compareNode)
	if evalErr != nil {
		return nil, evalErr
	}

	result.RolledDice = evaluator.RolledDice()

	var successCheckResultMessage string

	boolObj, objIsBoolean := obj.(*object.Boolean)
	if !objIsBoolean {
		return nil, fmt.Errorf("DRollComp: result is not a Boolean: %s", obj.Type())
	}

	if boolObj.Value == true {
		result.SuccessCheckResult = SUCCESS_CHECK_SUCCESS
		successCheckResultMessage = "成功"
	} else {
		result.SuccessCheckResult = SUCCESS_CHECK_FAILURE
		successCheckResultMessage = "失敗"
	}

	result.appendMessagePart(notation.Parenthesize(infixNotation1))
	result.appendMessagePart(infixNotation2)
	result.appendMessagePart(leftObj.Inspect())
	result.appendMessagePart(successCheckResultMessage)

	return result, nil
}

// determineCompareValues は比較式の可変ノードの値を決定する。
// 返り値はその結果の中置表記とエラー。
func determineCompareValues(node *ast.Compare, evaluator *evaluator.Evaluator) (string, error) {
	determineValuesErr := evaluator.DetermineValues(node)
	if determineValuesErr != nil {
		return "", determineValuesErr
	}

	infixNotation, infixNotationErr := notation.InfixNotation(node.Left(), true)
	if infixNotationErr != nil {
		return "", infixNotationErr
	}

	return infixNotation, nil
}
