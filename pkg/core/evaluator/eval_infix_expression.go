package evaluator

import (
	"fmt"
	"math"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalInfixExpression は中置式を評価する。
func (e *Evaluator) evalInfixExpression(
	node ast.InfixExpression,
) (object.Object, error) {
	left, right, err := e.evalInfixExpressionOperands(node)
	if err != nil {
		return nil, err
	}

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		leftInteger := left.(*object.Integer)
		rightInteger := right.(*object.Integer)

		switch n := node.(type) {
		case ast.Divide:
			return e.evalIntegerDivide(n, leftInteger, rightInteger)
		default:
			return e.evalIntegerInfixExpression(n.Operator(), leftInteger, rightInteger)
		}
	}

	return nil, fmt.Errorf("operator not implemented: %s %s %s",
		left.Type(), node.Operator(), right.Type())
}

// evalInfixExpressionOperands は中置式の左右のオペランドを評価する。
//
// 戻り値: 左のオペランドの評価結果, 右のオペランドの評価結果, エラー
func (e *Evaluator) evalInfixExpressionOperands(
	node ast.InfixExpression,
) (object.Object, object.Object, error) {
	if node.Left() == nil {
		return nil, nil, fmt.Errorf("operator %s: left is nil", node.Operator())
	}

	if node.Right() == nil {
		return nil, nil, fmt.Errorf("operator %s: right is nil", node.Operator())
	}

	left, leftErr := e.Eval(node.Left())
	if leftErr != nil {
		return nil, nil, leftErr
	}

	if left == nil {
		return nil, nil, fmt.Errorf("operator %s: evaluated left is nil", node.Operator())
	}

	right, rightErr := e.Eval(node.Right())
	if rightErr != nil {
		return left, nil, rightErr
	}

	if right == nil {
		return left, nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator())
	}

	return left, right, nil
}

// evalIntegerInfixExpression は整数ノード同士を子に持つ中置式を評価する。
func (e *Evaluator) evalIntegerInfixExpression(
	operator string,
	left *object.Integer,
	right *object.Integer,
) (object.Object, error) {
	leftValue := left.Value
	rightValue := right.Value

	switch operator {
	case "+":
		return object.NewInteger(leftValue + rightValue), nil
	case "-":
		return object.NewInteger(leftValue - rightValue), nil
	case "*":
		return object.NewInteger(leftValue * rightValue), nil
	case "D":
		return e.evalSumRoll(left, right)
	case "B", "R":
		return e.evalBasicRoll(left, right)
	case "...":
		return e.evalRandomNumber(left, right)
	case "=":
		return object.NewBoolean(leftValue == rightValue), nil
	case "<>":
		return object.NewBoolean(leftValue != rightValue), nil
	case "<":
		return object.NewBoolean(leftValue < rightValue), nil
	case ">":
		return object.NewBoolean(leftValue > rightValue), nil
	case "<=":
		return object.NewBoolean(leftValue <= rightValue), nil
	case ">=":
		return object.NewBoolean(leftValue >= rightValue), nil
	}

	return nil, fmt.Errorf("operator not implemented: %s %s %s",
		left.Type(), operator, right.Type())
}

// evalIntegerDivide は除算を評価する。
func (e *Evaluator) evalIntegerDivide(
	divide ast.Divide,
	left *object.Integer,
	right *object.Integer,
) (object.Object, error) {
	leftValue := left.Value
	rightValue := right.Value

	if rightValue == 0 {
		return nil, fmt.Errorf("%d divided by zero", leftValue)
	}

	switch divide.RoundingMethod() {
	case ast.ROUNDING_METHOD_ROUND_DOWN:
		// 除算（小数点以下切り捨て）
		return object.NewInteger(leftValue / rightValue), nil
	case ast.ROUNDING_METHOD_ROUND:
		{
			// 除算（小数点以下四捨五入）
			resultFloat := math.Round(float64(leftValue) / float64(rightValue))
			return object.NewInteger(int(resultFloat)), nil
		}
	case ast.ROUNDING_METHOD_ROUND_UP:
		{
			// 除算（小数点以下切り上げ）
			resultFloat := math.Ceil(float64(leftValue) / float64(rightValue))
			return object.NewInteger(int(resultFloat)), nil
		}
	default:
		return nil, fmt.Errorf("evalIntegerDivide: unknown rounding method")
	}
}

// evalSumRoll は加算ロールを評価する。
func (e *Evaluator) evalSumRoll(
	num *object.Integer,
	sides *object.Integer,
) (object.Object, error) {
	numVal := num.Value
	sidesVal := sides.Value

	rolledDice, err := e.RollDice(numVal, sidesVal)
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, d := range rolledDice {
		sum += d.Value
	}

	return object.NewInteger(sum), nil
}

// evalBasicRoll はバラバラロールを評価する。
// 返り値は、整数オブジェクトを要素として持つ配列オブジェクト、およびエラー。
func (e *Evaluator) evalBasicRoll(
	num *object.Integer,
	sides *object.Integer,
) (*object.Array, error) {
	numVal := num.Value
	sidesVal := sides.Value

	rolledDice, err := e.RollDice(numVal, sidesVal)
	if err != nil {
		return nil, err
	}

	intObjs := make([]object.Object, 0, len(rolledDice))
	for _, d := range rolledDice {
		intObjs = append(intObjs, object.NewInteger(d.Value))
	}

	return object.NewArrayByMove(intObjs), nil
}

// evalRandomNumber はランダム数値取り出しを評価する。
func (e *Evaluator) evalRandomNumber(
	min *object.Integer,
	max *object.Integer,
) (object.Object, error) {
	minValue := min.Value
	maxValue := max.Value

	if minValue >= maxValue {
		return nil,
			fmt.Errorf(
				"evalRandomNumber: min (%d) must be less than max (%d)",
				minValue,
				maxValue,
			)
	}

	randRange := maxValue - minValue + 1
	rolledDice, err := e.RollDice(1, randRange)
	if err != nil {
		return nil, err
	}

	resultValue := (minValue - 1) + rolledDice[0].Value

	return object.NewInteger(resultValue), nil
}
