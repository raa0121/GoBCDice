package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/object"
)

// Evalはnodeを評価してObjectに変換し、返す
func Eval(node ast.Node) (object.Object, error) {
	// 型で分岐する
	switch node := node.(type) {
	case *ast.Command:
		// TODO: もしかしたらコマンドの種類で分岐する？
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node)
	case *ast.InfixExpression:
		return evalInfixExpression(node)
	case *ast.Int:
		return &object.Integer{Value: node.Value}, nil
	}

	return nil, fmt.Errorf("unknown type: %s", node.Type())
}

func evalPrefixExpression(node *ast.PrefixExpression) (object.Object, error) {
	if node.Right == nil {
		return nil, fmt.Errorf("operator %s: right is nil", node.Operator)
	}

	right, err := Eval(node.Right)
	if err != nil {
		return nil, err
	}

	if right == nil {
		return nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator)
	}

	if right.Type() == object.INTEGER_OBJ {
		return evalIntegerPrefixExpression(node.Operator, right.(*object.Integer))
	}

	return nil, fmt.Errorf("unknown operator: %s%s", node.Operator, right.Type())
}

func evalIntegerPrefixExpression(operator string, right *object.Integer) (object.Object, error) {
	value := right.Value

	switch operator {
	case "-":
		return &object.Integer{Value: -value}, nil
	}

	return nil, fmt.Errorf("unknown operator: %s%s", operator, right.Type())
}

func evalInfixExpression(node *ast.InfixExpression) (object.Object, error) {
	if node.Left == nil {
		return nil, fmt.Errorf("operator %s: left is nil", node.Operator)
	}

	if node.Right == nil {
		return nil, fmt.Errorf("operator %s: right is nil", node.Operator)
	}

	left, leftErr := Eval(node.Left)
	if leftErr != nil {
		return nil, leftErr
	}

	if left == nil {
		return nil, fmt.Errorf("operator %s: evaluated left is nil", node.Operator)
	}

	right, rightErr := Eval(node.Right)
	if rightErr != nil {
		return nil, rightErr
	}

	if right == nil {
		return nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator)
	}

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(
			node.Operator,
			left.(*object.Integer),
			right.(*object.Integer),
		)
	}

	return nil, fmt.Errorf("unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
}

func evalIntegerInfixExpression(
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
		return &object.Integer{Value: leftValue / rightValue}, nil
	}

	return nil, fmt.Errorf("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}
