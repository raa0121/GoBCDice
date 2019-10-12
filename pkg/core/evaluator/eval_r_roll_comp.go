package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalRRollComp は個数振り足しロールの成功数カウントを評価する。
func (e *Evaluator) evalRRollComp(node *ast.RRollComp) (*object.RRollCompResult, error) {
	compareNode := node.Expression().(*ast.Compare)

	// 左辺を評価する
	valueGroupsObj, evalRRollListErr := e.Eval(compareNode.Left())
	if evalRRollListErr != nil {
		return nil, evalRRollListErr
	}

	// 右辺を評価する
	evaluatedTargetObj, evalTargetErr := e.Eval(compareNode.Right())
	if evalTargetErr != nil {
		return nil, evalTargetErr
	}

	valueGroupsArray := valueGroupsObj.(*object.Array)
	evaluatedTargetNode :=
		ast.NewInt(evaluatedTargetObj.(*object.Integer).Value)

	// 振られた各ダイスに対して成功判定を行い、成功数を数える
	numOfSuccesses := 0
	for _, vg := range valueGroupsArray.Elements {
		valuesArray := vg.(*object.Array)
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
	}

	return object.NewRRollCompResult(
		valueGroupsArray,
		object.NewInteger(numOfSuccesses),
	), nil
}
