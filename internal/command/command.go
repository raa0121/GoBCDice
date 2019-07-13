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

	infixNotation, notationErr := notation.InfixNotation(node)
	if notationErr != nil {
		return nil, notationErr
	}

	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

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

	infixNotationOfNodeWithEvaluatedVarArgs, evaluateVarArgsErr :=
		evaluateVarArgs(node, evaluator)
	if evaluateVarArgsErr != nil {
		return nil, evaluateVarArgsErr
	}

	infixNotationOfNodeWithDeterminedValues, determineValuesErr :=
		determineValues(node, evaluator)
	if determineValuesErr != nil {
		return nil, determineValuesErr
	}

	obj, evalErr := evaluator.Eval(node)
	if evalErr != nil {
		return nil, evalErr
	}

	result.appendMessagePart(notation.Parenthesize(infixNotationOfNodeWithEvaluatedVarArgs))
	result.appendMessagePart(infixNotationOfNodeWithDeterminedValues)
	result.appendMessagePart(obj.Inspect())

	return result, nil
}

// 加算ロールなどの可変ノードの引数を評価して整数に変換する
func evaluateVarArgs(node ast.Node, evaluator *evaluator.Evaluator) (string, error) {
	evalVarArgsErr := evaluator.EvalVarArgs(node)
	if evalVarArgsErr != nil {
		return "", evalVarArgsErr
	}

	infixNotationOfNodeWithEvaluatedVarArgs, iNForEvaluatedVarArgsErr :=
		notation.InfixNotation(node)
	if iNForEvaluatedVarArgsErr != nil {
		return "", iNForEvaluatedVarArgsErr
	}

	return infixNotationOfNodeWithEvaluatedVarArgs, nil
}

// 加算ロールなどの可変ノードの値を確定させる
func determineValues(node ast.Node, evaluator *evaluator.Evaluator) (string, error) {
	nodeForDetermineValues := node
	// TODO: 加算ロールなどの可変ノードの値を確定させる
	nodeWithDeterminedValues := nodeForDetermineValues

	infixNotationOfNodeWithDeterminedValues, iNForDeterminedValuesErr :=
		notation.InfixNotation(nodeWithDeterminedValues)
	if iNForDeterminedValuesErr != nil {
		return "", iNForDeterminedValuesErr
	}

	return infixNotationOfNodeWithDeterminedValues, nil
}
