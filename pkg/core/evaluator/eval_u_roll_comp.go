package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalURollComp は上方無限ロールの成功数カウントを評価する。
func (e *Evaluator) evalURollComp(node *ast.Command) (*object.URollCompResult, error) {
	compareNode := node.Expression.(*ast.BasicInfixExpression)

	// 両辺を評価する
	leftObj, targetObj, evalOperandsErr :=
		e.evalInfixExpressionOperands(compareNode)
	if evalOperandsErr != nil {
		return nil, evalOperandsErr
	}

	rollResult := leftObj.(*object.URollExprResult)

	// 出目のグループごとに、合計値に修正値を加算する
	sumOfGroups := rollResult.SumOfGroups()
	modifiedSumOfGroups := make([]object.Object, 0, sumOfGroups.Length())
	for _, sum := range sumOfGroups.Elements {
		modifiedSumOfGroups = append(
			modifiedSumOfGroups,
			sum.(*object.Integer).Add(rollResult.Modifier()),
		)
	}

	// 成功数を数える
	numOfSuccesses, countErr := e.countNumOfSuccesses(
		object.NewArrayByMove(modifiedSumOfGroups),
		compareNode.Operator(),
		targetObj.(*object.Integer),
	)
	if countErr != nil {
		return nil, countErr
	}

	return object.NewURollCompResult(rollResult, numOfSuccesses), nil
}
