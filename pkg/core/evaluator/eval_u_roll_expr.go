package evaluator

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalURollList は上方無限ロール列を評価する。
func (e *Evaluator) evalURollList(node *ast.RRollList) (*object.Array, error) {
	if node.Threshold.IsNil() {
		return nil, fmt.Errorf("evalURollList: threshold is nil")
	}

	// 閾値を評価する
	thresholdObj, evalThresholdErr := e.Eval(node.Threshold)
	if evalThresholdErr != nil {
		return nil, evalThresholdErr
	}
	thresholdInt := thresholdObj.(*object.Integer)

	// ダイスロール結果を格納する配列
	valueGroups := []object.Object{}
	for _, uRoll := range node.RRolls {
		valueGroupsParts, evalURollListErr := e.evalCompoundingRoll(uRoll, thresholdInt)
		if evalURollListErr != nil {
			return nil, evalURollListErr
		}

		valueGroups = append(valueGroups, valueGroupsParts.Elements...)
	}

	return object.NewArrayByMove(valueGroups), nil
}

// evalURollExpr は上方無限ロール式を評価する。
func (e *Evaluator) evalURollExpr(node *ast.URollExpr) (*object.URollExprResult, error) {
	var modifier *object.Integer

	if node.Bonus == nil {
		modifier = object.NewInteger(0)
	} else {
		// ボーナスを評価する
		bonusObj, evalBonusErr := e.Eval(node.Bonus)
		if evalBonusErr != nil {
			return nil, evalBonusErr
		}

		modifier = bonusObj.(*object.Integer)
	}

	valueGroups, valueGroupsErr := e.evalURollList(node.URollList)
	if valueGroupsErr != nil {
		return nil, valueGroupsErr
	}

	return object.NewURollExprResult(valueGroups, modifier), nil
}

// evalCompoundingRoll は上方無限ロールを評価する。
// 返り値は、出目のグループを要素として持つ配列オブジェクト、およびエラー。
func (e *Evaluator) evalCompoundingRoll(
	node *ast.VariableInfixExpression,
	threshold *object.Integer,
) (*object.Array, error) {
	left, right, operandsEvalErr := e.evalInfixExpressionOperands(node)
	if operandsEvalErr != nil {
		return nil, operandsEvalErr
	}

	if left.Type() != object.INTEGER_OBJ {
		return nil, fmt.Errorf("evalCompoundingRoll: evaluated num is not an Integer")
	}

	if right.Type() != object.INTEGER_OBJ {
		return nil, fmt.Errorf("evalCompoundingRoll: evaluated sides is not an Integer")
	}

	num := left.(*object.Integer)
	sides := right.(*object.Integer)

	numVal := num.Value
	sidesVal := sides.Value
	thresholdVal := threshold.Value

	if sidesVal < 1 {
		return nil, fmt.Errorf(
			"evalCompoundingRoll(num: %d, sides: %d): ダイスの面数が少なすぎます",
			numVal,
			sidesVal,
		)
	}

	if numVal < 1 {
		return nil, fmt.Errorf(
			"evalCompoundingRoll(num: %d, sides: %d): 振るダイス数が少なすぎます",
			numVal,
			sidesVal,
		)
	}

	valueGroups := make([]object.Object, 0, numVal)
	for i := 0; i < numVal; i++ {
		valueGroup := []object.Object{}

		for j := 0; j < e.MaxRerolls; j++ {
			rolledDice, err := e.RollDice(1, sidesVal)
			if err != nil {
				return nil, err
			}

			value := rolledDice[0].Value
			valueGroup = append(valueGroup, object.NewInteger(value))

			if value < thresholdVal {
				break
			}
		}

		valueGroupArray := object.NewArrayByMove(valueGroup)
		valueGroups = append(valueGroups, valueGroupArray)
	}

	return object.NewArrayByMove(valueGroups), nil
}
