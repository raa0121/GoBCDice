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
