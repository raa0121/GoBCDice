package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/notation"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

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
	result.AppendMessagePart(notation.Parenthesize(infixNotation))
	result.AppendMessagePart(resultObj.Value)

	return result, nil
}
