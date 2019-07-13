package command

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/evaluator"
	"github.com/raa0121/GoBCDice/internal/notation"
)

// Executeは指定されたコマンドを実行する。
//
// * commandNode: コマンドのノード
// * gameId: ゲーム識別子
// * evaluator: 評価器
func Execute(
	commandNode ast.Command,
	gameId string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	switch c := commandNode.(type) {
	case *ast.Calc:
		return executeCalc(c, gameId, evaluator)
	case *ast.DRollExpr:
		return executeDRollExpr(c, gameId, evaluator)
	}

	return nil, fmt.Errorf("command execution not implemented: %s", commandNode.Type())
}

// executeCalcは計算を実行する
func executeCalc(
	node *ast.Calc,
	gameId string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameId: gameId,
	}

	// 抽象構文木を中置表記に変換する
	infixNotation, notationErr := notation.InfixNotation(node)
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

// executeDRollExprは加算ロールを実行する
func executeDRollExpr(
	node *ast.DRollExpr,
	gameId string,
	evaluator *evaluator.Evaluator,
) (*Result, error) {
	result := &Result{
		GameId: gameId,
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

	// 結果のメッセージを作る
	result.appendMessagePart(notation.Parenthesize(infixNotationOfNodeWithEvaluatedVarArgs))
	result.appendMessagePart(infixNotationOfNodeWithDeterminedValues)
	result.appendMessagePart(obj.Inspect())

	return result, nil
}

// 加算ロールなどの可変ノードの引数を評価して整数に変換する。
// 返り値はその結果の中置表記とエラー。
func evalVarArgs(node ast.Node, evaluator *evaluator.Evaluator) (string, error) {
	evalErr := evaluator.EvalVarArgs(node)
	if evalErr != nil {
		return "", evalErr
	}

	infixNotation, infixNotationErr := notation.InfixNotation(node)
	if infixNotationErr != nil {
		return "", infixNotationErr
	}

	return infixNotation, nil
}

// determineValuesは加算ロールなどの可変ノードの値を決定する。
// 返り値はその結果の中置表記とエラー。
func determineValues(node ast.Node, evaluator *evaluator.Evaluator) (string, error) {
	determineValuesErr := evaluator.DetermineValues(node)
	if determineValuesErr != nil {
		return "", determineValuesErr
	}

	infixNotation, infixNotationErr := notation.InfixNotation(node)
	if infixNotationErr != nil {
		return "", infixNotationErr
	}

	return infixNotation, nil
}
