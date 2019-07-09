package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/die"
	"github.com/raa0121/GoBCDice/internal/die/roller"
	"github.com/raa0121/GoBCDice/internal/object"
	"math"
)

// 評価器の構造体
type Evaluator struct {
	diceRoller *roller.DiceRoller
	env        *Environment
}

// NewEvaluatorは新しい評価器を返す。
//
// * diceRoller: ダイスローラー
// * env: 評価環境
func NewEvaluator(diceRoller *roller.DiceRoller, env *Environment) *Evaluator {
	return &Evaluator{
		diceRoller: diceRoller,
		env:        env,
	}
}

// RolledDiceは、ダイスロール結果を返す
func (e *Evaluator) RolledDice() []die.Die {
	return e.env.RolledDice()
}

// Evalはnodeを評価してObjectに変換し、返す
func (e *Evaluator) Eval(node ast.Node) (object.Object, error) {
	// 型で分岐する
	switch n := node.(type) {
	case ast.Command:
		// TODO: もしかしたらコマンドの種類で分岐する？
		return e.Eval(n.Expression())
	case ast.PrefixExpression:
		return e.evalPrefixExpression(n)
	case ast.InfixExpression:
		return e.evalInfixExpression(n)
	case *ast.Int:
		return &object.Integer{Value: n.Value}, nil
	}

	return nil, fmt.Errorf("unknown type: %s", node.Type())
}

func (e *Evaluator) evalPrefixExpression(
	node ast.PrefixExpression,
) (object.Object, error) {
	if node.Right() == nil {
		return nil, fmt.Errorf("operator %s: right is nil", node.Operator())
	}

	right, err := e.Eval(node.Right())
	if err != nil {
		return nil, err
	}

	if right == nil {
		return nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator())
	}

	if right.Type() == object.INTEGER_OBJ {
		return e.evalIntegerPrefixExpression(
			node.Operator(), right.(*object.Integer))
	}

	return nil, fmt.Errorf("operator not implemented: %s%s",
		node.Operator(), right.Type())
}

func (e *Evaluator) evalIntegerPrefixExpression(
	operator string,
	right *object.Integer,
) (object.Object, error) {
	value := right.Value

	switch operator {
	case "-":
		return &object.Integer{Value: -value}, nil
	}

	return nil, fmt.Errorf("operator not implemented: %s%s",
		operator, right.Type())
}

func (e *Evaluator) evalInfixExpression(
	node ast.InfixExpression,
) (object.Object, error) {
	if node.Left() == nil {
		return nil, fmt.Errorf("operator %s: left is nil", node.Operator())
	}

	if node.Right() == nil {
		return nil, fmt.Errorf("operator %s: right is nil", node.Operator())
	}

	left, leftErr := e.Eval(node.Left())
	if leftErr != nil {
		return nil, leftErr
	}

	if left == nil {
		return nil, fmt.Errorf("operator %s: evaluated left is nil", node.Operator())
	}

	right, rightErr := e.Eval(node.Right())
	if rightErr != nil {
		return nil, rightErr
	}

	if right == nil {
		return nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator())
	}

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return e.evalIntegerInfixExpression(
			node.Operator(),
			left.(*object.Integer),
			right.(*object.Integer),
		)
	}

	return nil, fmt.Errorf("operator not implemented: %s %s %s",
		left.Type(), node.Operator(), right.Type())
}

func (e *Evaluator) evalIntegerInfixExpression(
	operator string,
	left *object.Integer,
	right *object.Integer,
) (object.Object, error) {
	leftValue := left.Value
	rightValue := right.Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}, nil
	case "-":
		return &object.Integer{Value: leftValue - rightValue}, nil
	case "*":
		return &object.Integer{Value: leftValue * rightValue}, nil
	case "/":
		// 除算（小数点以下切り捨て）
		return &object.Integer{Value: leftValue / rightValue}, nil
	case "/U":
		{
			// 除算（小数点以下切り上げ）
			resultFloat := math.Ceil(float64(leftValue) / float64(rightValue))
			return &object.Integer{Value: int(resultFloat)}, nil
		}
	case "/R":
		{
			// 除算（小数点以下四捨五入）
			resultFloat := math.Round(float64(leftValue) / float64(rightValue))
			return &object.Integer{Value: int(resultFloat)}, nil
		}
	case "D":
		return e.evalSumRoll(left, right)
	}

	return nil, fmt.Errorf("operator not implemented: %s %s %s",
		left.Type(), operator, right.Type())
}

func (e *Evaluator) evalSumRoll(
	num *object.Integer,
	sides *object.Integer,
) (object.Object, error) {
	numVal := num.Value
	sidesVal := sides.Value

	rolledDice, err := e.rollDice(numVal, sidesVal)
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, d := range rolledDice {
		sum += d.Value
	}

	return &object.Integer{Value: sum}, nil
}

// rollDiceは、sides個の面を持つダイスをnum個振り、その結果を返す。
// また、ダイスロールの結果を記録する。
func (e *Evaluator) rollDice(num int, sides int) ([]die.Die, error) {
	rolledDice, err := e.diceRoller.RollDice(num, sides)
	if err != nil {
		return nil, err
	}

	e.env.AppendRolledDice(rolledDice)

	return rolledDice, nil
}
