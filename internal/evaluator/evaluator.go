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
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.Int:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
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
