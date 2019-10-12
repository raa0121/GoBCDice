package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
)

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
