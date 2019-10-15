package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalRRollComp は個数振り足しロールの成功数カウントを評価する。
func (e *Evaluator) evalRRollComp(node *ast.Command) (*object.RRollCompResult, error) {
	compareNode := node.Expression.(*ast.BasicInfixExpression)

	// 両辺を評価する
	valueGroupsObj, targetObj, evalOperandsErr :=
		e.evalInfixExpressionOperands(compareNode)
	if evalOperandsErr != nil {
		return nil, evalOperandsErr
	}

	valueGroupsArray := valueGroupsObj.(*object.Array)
	targetIntObj := targetObj.(*object.Integer)

	// 成功判定を行い、出目のグループごとの成功数を加算する
	numOfSuccesses := 0
	for _, valueGroupObj := range valueGroupsArray.Elements {
		numOfSuccessesInGroup, countErr := e.countNumOfSuccesses(
			valueGroupObj.(*object.Array),
			compareNode.Operator(),
			targetIntObj,
		)
		if countErr != nil {
			return nil, countErr
		}

		numOfSuccesses += numOfSuccessesInGroup.Value
	}

	return object.NewRRollCompResult(
		valueGroupsArray,
		object.NewInteger(numOfSuccesses),
	), nil
}
