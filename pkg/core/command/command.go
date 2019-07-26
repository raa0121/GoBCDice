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
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// Execute は指定されたコマンドを実行する。
//
// commandNode: コマンドのノード,
// gameID: ゲーム識別子,
// evaluator: 評価器。
func Execute(
	commandNode ast.Command,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	switch c := commandNode.(type) {
	case *ast.Calc:
		return executeCalc(c, gameID, evaluator)
	case *ast.DRollExpr:
		return executeDRollExpr(c, gameID, evaluator)
	case *ast.DRollComp:
		return executeDRollComp(c, gameID, evaluator)
	}

	return nil, fmt.Errorf("command execution not implemented: %s", commandNode.Type())
}

// executeCalc は計算を実行する。
func executeCalc(
	node *ast.Calc,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	// 抽象構文木を中置表記に変換する
	infixNotation, notationErr := notation.InfixNotation(node, true)
	if notationErr != nil {
		return nil, notationErr
	}

	// 抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	// 結果のメッセージを作る
	result.appendMessagePart(infixNotation)
	result.appendMessagePart("計算結果")
	result.appendMessagePart(obj.Inspect())

	return result, nil
}

// executeDRollExpr は加算ロールを実行する。
func executeDRollExpr(
	node *ast.DRollExpr,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	// 加算ロールなどの可変ノードの引数を評価して整数に変換する
	infixNotationOfNodeWithEvaluatedVarArgs, evalVarArgsErr :=
		evalVarArgs(node, evaluator)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 加算ロールなどの可変ノードの値を決定する
	infixNotationOfNodeWithDeterminedValues, determineValuesErr :=
		determineValues(node, evaluator)
	if determineValuesErr != nil {
		return nil, determineValuesErr
	}

	// 変換された抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.appendMessagePart(notation.Parenthesize(infixNotationOfNodeWithEvaluatedVarArgs))
	result.appendMessagePart(infixNotationOfNodeWithDeterminedValues)
	result.appendMessagePart(obj.Inspect())

	return result, nil
}

// executeDRollCompp は加算ロール式の成功判定を実行する。
func executeDRollComp(
	node *ast.DRollComp,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	compareNode, exprIsCompareNode := node.Expression().(*ast.Compare)
	if !exprIsCompareNode {
		return nil, fmt.Errorf("DRollComp: expression is not a Compare node: %s", node.Type())
	}

	// 左辺の可変ノードの引数および右辺を評価する
	infixNotationOfNodeWithEvaluatedVarArgs, evalVarArgsAndRightErr :=
		evalDRollCompVarArgsAndRight(compareNode, evaluator)
	if evalVarArgsAndRightErr != nil {
		return nil, evalVarArgsAndRightErr
	}

	// 加算ロールなどの可変ノードの値を決定する
	infixNotationOfLeftWithDeterminedValues, determineValuesErr :=
		determineDRollCompValues(compareNode, evaluator)
	if determineValuesErr != nil {
		return nil, determineValuesErr
	}

	// 左辺を評価する
	leftObj, leftEvalErr := evalDRollCompLeft(compareNode, evaluator)
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

	result.appendMessagePart(notation.Parenthesize(infixNotationOfNodeWithEvaluatedVarArgs))
	result.appendMessagePart(infixNotationOfLeftWithDeterminedValues)
	result.appendMessagePart(leftObj.Inspect())
	result.appendMessagePart(successCheckResultMessage)

	return result, nil
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

// determineValues は加算ロールなどの可変ノードの値を決定する。
// 返り値はその結果の中置表記とエラー。
func determineValues(node ast.Node, evaluator *evaluator.Evaluator) (string, error) {
	determineValuesErr := evaluator.DetermineValues(node)
	if determineValuesErr != nil {
		return "", determineValuesErr
	}

	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return "", infixNotationErr
	}

	return infixNotation, nil
}

// evalDRollCompVarArgsAndRight は、加算ロール式成功判定内の左辺の可変ノードの引数および右辺を評価する。
// 返り値はその結果の中置表記とエラー。
func evalDRollCompVarArgsAndRight(
	node *ast.Compare,
	evaluator *evaluator.Evaluator,
) (string, error) {
	// 加算ロールを含む左辺の可変ノードの引数を評価して整数に変換する
	leftEvalErr := evaluator.EvalVarArgs(node.Left())
	if leftEvalErr != nil {
		return "", leftEvalErr
	}

	// 右辺（閾値）を評価して整数に変換する
	rightObj, rightEvalErr := evaluator.Eval(node.Right())
	if rightEvalErr != nil {
		return "", rightEvalErr
	}

	evaluatedRight := ast.NewInt(rightObj.(*object.Integer).Value, token.Token{})
	node.SetRight(evaluatedRight)

	// 中置表記を生成する
	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return "", infixNotationErr
	}

	return infixNotation, nil
}

// determineDRollCompValues は加算ロール式成功判定の可変ノードの値を決定する。
// 返り値はその結果の中置表記とエラー。
func determineDRollCompValues(node *ast.Compare, evaluator *evaluator.Evaluator) (string, error) {
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

// evalDRollCompLeft は加算ロール式成功判定の左辺を評価する。
func evalDRollCompLeft(
	node *ast.Compare,
	evaluator *evaluator.Evaluator,
) (object.Object, error) {
	leftObj, leftEvalErr := evaluator.Eval(node.Left())
	if leftEvalErr != nil {
		return nil, leftEvalErr
	}

	evaluatedLeft := ast.NewInt(leftObj.(*object.Integer).Value, token.Token{})
	node.SetLeft(evaluatedLeft)

	return leftObj, nil
}
