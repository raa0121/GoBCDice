package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalBRollComp はバラバラロールの成功数カウントを評価する。
func (e *Evaluator) evalBRollComp(node *ast.Command) (*object.BRollCompResult, error) {
	compareNode := node.Expression.(*ast.BasicInfixExpression)

	// 左辺を評価する
	valuesObj, evalBRollListErr := e.Eval(compareNode.Left())
	if evalBRollListErr != nil {
		return nil, evalBRollListErr
	}

	// 右辺を評価する
	evaluatedTargetObj, evalTargetErr := e.Eval(compareNode.Right())
	if evalTargetErr != nil {
		return nil, evalTargetErr
	}

	valuesArray := valuesObj.(*object.Array)
	evaluatedTargetNode :=
		ast.NewInt(evaluatedTargetObj.(*object.Integer).Value)

	// 振られた各ダイスに対して成功判定を行い、成功数を数える
	numOfSuccesses := 0
	for _, el := range valuesArray.Elements {
		valueNode := ast.NewInt(el.(*object.Integer).Value)
		valueCompareNode := ast.NewCompare(
			valueNode,
			compareNode.Operator(),
			evaluatedTargetNode,
		)

		r, compErr := e.Eval(valueCompareNode)
		if compErr != nil {
			return nil, compErr
		}

		if success := r.(*object.Boolean).Value; success {
			numOfSuccesses++
		}
	}

	return object.NewBRollCompResult(valuesArray, object.NewInteger(numOfSuccesses)), nil
}
