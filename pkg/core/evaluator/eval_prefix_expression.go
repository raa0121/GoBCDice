package evaluator

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalPrefixExpression は前置式を評価する。
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

// evalIntegerPrefixExpression は整数ノードを子に持つ前置式を評価する。
func (e *Evaluator) evalIntegerPrefixExpression(
	operator string,
	right *object.Integer,
) (object.Object, error) {
	value := right.Value

	switch operator {
	case "-":
		return object.NewInteger(-value), nil
	}

	return nil, fmt.Errorf("operator not implemented: %s%s",
		operator, right.Type())
}
