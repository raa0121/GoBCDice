package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// SetRerollThreshold は振り足しの閾値を設定する。
func (e *Evaluator) SetRerollThreshold(node *ast.Command) error {
	compareNode := node.Expression.(*ast.BasicInfixExpression)
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

// CheckRRollThreshold は個数振り足しロールの振り足しの閾値をチェックする。
func (e *Evaluator) CheckRRollThreshold(node *ast.RRollList) error {
	if node.Threshold.IsNil() {
		return fmt.Errorf("2R6>=5 あるいは 2R6[5] のように振り足し目標値を指定してください")
	}

	return e.checkRollThreshold(node.Threshold)
}

// CheckURollThreshold は上方無限ロールの振り足しの閾値をチェックする。
func (e *Evaluator) CheckURollThreshold(node *ast.RRollList) error {
	if node.Threshold.IsNil() {
		return fmt.Errorf("2U6[5] のように振り足し目標値を指定してください")
	}

	return e.checkRollThreshold(node.Threshold)
}

// checkRollThreshold は指定されていた閾値が条件を満たしているか確認する
func (e *Evaluator) checkRollThreshold(thresholdNode ast.Node) error {
	thresholdObj, evalErr := e.Eval(thresholdNode)
	if evalErr != nil {
		return fmt.Errorf("閾値評価エラー: %s", evalErr)
	}

	thresholdInt := thresholdObj.(*object.Integer)
	threshold := thresholdInt.Value

	if threshold < 2 {
		return fmt.Errorf("振り足し目標値として2以上の整数を指定してください")
	}

	return nil
}
