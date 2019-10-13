package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
)

// executeCalc は計算を実行する。
func executeCalc(
	node *ast.Command,
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
