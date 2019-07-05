package evaluator

import (
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/object"
)

// Evalはnodeを評価してObjectに変換し、返す
func Eval(node ast.Node) object.Object {
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
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalPrefixExpression(node *ast.PrefixExpression) object.Object {
	if node.Right == nil {
		return nil
	}

	right := Eval(node.Right)
	if right == nil {
		return nil
	}

	switch node.Operator {
	case "-":
		return evalUnaryMinusOperatorExpression(right)
	}

	return nil
}

func evalUnaryMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return nil
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(node *ast.InfixExpression) object.Object {
	if node.Left == nil || node.Right == nil {
		return nil
	}

	left := Eval(node.Left)
	if left == nil {
		return nil
	}

	right := Eval(node.Right)
	if right == nil {
		return nil
	}

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(
			node.Operator,
			left.(*object.Integer),
			right.(*object.Integer),
		)
	}

	return nil
}

func evalIntegerInfixExpression(
	operator string,
	left *object.Integer,
	right *object.Integer,
) object.Object {
	leftValue := left.Value
	rightValue := right.Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	}

	return nil
}
