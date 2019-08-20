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
	"strings"
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
	case *ast.Choice:
		return executeChoice(c, gameID, evaluator)
	}

	return nil, fmt.Errorf("command execution not implemented: %s", node.Type())
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
	infixNotation1, evalVarArgsErr := evalVarArgs(node, evaluator)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 加算ロールなどの可変ノードの値を決定する
	infixNotation2, determineValuesErr := determineValues(node, evaluator)
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
	result.appendMessagePart(notation.Parenthesize(infixNotation1))
	result.appendMessagePart(infixNotation2)
	result.appendMessagePart(obj.Inspect())

	return result, nil
}

// executeDRollComp は加算ロール式の成功判定を実行する。
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
	result.appendMessagePart(notation.Parenthesize(infixNotation))
	result.appendMessagePart(arrayObj.JoinedElements(","))

	return result, nil
}

// executeBRollComp はバラバラロールの成功数カウントを実行する。
func executeBRollComp(
	node *ast.BRollComp,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	// 左辺の可変ノードの引数および右辺を評価する
	infixNotation, evalVarArgsErr := evalVarArgs(
		node.Expression().(*ast.Compare),
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
	result.appendMessagePart(notation.Parenthesize(infixNotation))
	result.appendMessagePart(resultObj.Values.JoinedElements(","))
	result.appendMessagePart("成功数" + resultObj.NumOfSuccesses.Inspect())

	return result, nil
}

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
	evalVarArgsErr := evaluator.EvalVarArgs(node)
	if evalVarArgsErr != nil {
		return nil, evalVarArgsErr
	}

	// 中置表記を生成する
	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return nil, infixNotationErr
	}

	result.appendMessagePart(notation.Parenthesize(infixNotation))

	// 振り足しの閾値を確認する
	checkRerollThresholdErr := evaluator.CheckRerollThreshold(node)
	if checkRerollThresholdErr != nil {
		result.appendMessagePart(checkRerollThresholdErr.Error())
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
	result.appendMessagePart(formatRRollValues(valueGroups))

	return result, nil
}

// executeRRollComp は個数振り足しロールの成功数カウントを実行する。
func executeRRollComp(
	node *ast.RRollComp,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	compareNode := node.Expression().(*ast.Compare)

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
		evaluator.CheckRerollThreshold(compareNode.Left().(*ast.RRollList))
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

// executeChoice はランダム選択を実行する。
func executeChoice(
	node *ast.Choice,
	gameID string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameID: gameID,
	}

	// 中置表記を記録しておく
	infixNotation, infixNotationErr := notation.InfixNotation(node, true)
	if infixNotationErr != nil {
		return nil, infixNotationErr
	}

	// 抽象構文木を評価する
	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	resultObj := obj.(*object.String)
	result.RolledDice = evaluator.RolledDice()

	// 結果のメッセージを作る
	result.appendMessagePart(notation.Parenthesize(infixNotation))
	result.appendMessagePart(resultObj.Value)

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

// formatRRollValues は個数振り足しロールの出目を整形する。
func formatRRollValues(valueGroups *object.Array) string {
	valueGroupStrs := make([]string, 0, len(valueGroups.Elements))
	for _, valuesObj := range valueGroups.Elements {
		valuesArray := valuesObj.(*object.Array)
		valueGroupStrs = append(valueGroupStrs, valuesArray.JoinedElements(","))
	}

	return strings.Join(valueGroupStrs, " + ")
}
