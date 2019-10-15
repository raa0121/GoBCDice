package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalBRollComp はバラバラロールの成功数カウントを評価する。
func (e *Evaluator) evalBRollComp(node *ast.Command) (*object.BRollCompResult, error) {
	compareNode := node.Expression.(*ast.BasicInfixExpression)

	// 両辺を評価する
	valuesObj, targetObj, evalOperandsErr :=
		e.evalInfixExpressionOperands(compareNode)
	if evalOperandsErr != nil {
		return nil, evalOperandsErr
	}

	valuesArray := valuesObj.(*object.Array)

	// 成功数を数える
	numOfSuccesses, countErr := e.countNumOfSuccesses(
		valuesObj.(*object.Array),
		compareNode.Operator(),
		targetObj.(*object.Integer),
	)
	if countErr != nil {
		return nil, countErr
	}

	return object.NewBRollCompResult(valuesArray, numOfSuccesses), nil
}
