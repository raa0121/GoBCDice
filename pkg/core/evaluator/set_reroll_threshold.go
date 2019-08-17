package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// SetRerollThreshold は振り足しの閾値を設定する。
func (e *Evaluator) SetRerollThreshold(node *ast.RRollComp) error {
	compareNode := node.Expression().(*ast.Compare)
	rRollList := compareNode.Left().(*ast.RRollList)

	// xRn[...] で閾値が設定されていた場合は何もしない
	if !rRollList.Threshold.IsNil() {
		return nil
	}

	thresholdObj, evalErr := e.Eval(compareNode.Right())
	if evalErr != nil {
		return evalErr
	}

	rRollList.Threshold = ast.NewInt(thresholdObj.(*object.Integer).Value)

	return nil
}
