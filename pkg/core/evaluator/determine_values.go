package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
)

// DetermneValuesは、可変ノードの値を決定する
func (e *Evaluator) DetermineValues(node ast.Node) error {
	switch n := node.(type) {
	case *ast.Command:
		return e.determineValuesInCommand(n)
	case ast.PrefixExpression:
		return e.determineValuesInPrefixExpression(n)
	case ast.InfixExpression:
		return e.determineValuesInInfixExpression(n)
	}

	return fmt.Errorf("DetermineValues not implemented: %s", node.Type())
}

func (e *Evaluator) determineValueOfVariableExpr(node ast.Node) (ast.Node, error) {
	if node.Type() == ast.D_ROLL_NODE {
		return e.determineValueOfDRoll(node.(*ast.VariableInfixExpression))
	}

	return nil, fmt.Errorf("determineValueOfVariableExpr not implemented: %s", node.Type())
}

func (e *Evaluator) determineValueOfDRoll(
	node *ast.VariableInfixExpression,
) (*ast.SumRollResult, error) {
	num, numIsInt := node.Left().(*ast.Int)
	if !numIsInt {
		return nil, fmt.Errorf("num is not Int: %s", node.Left().Type())
	}

	sides, sidesIsInt := node.Right().(*ast.Int)
	if !sidesIsInt {
		return nil, fmt.Errorf("sides is not Int: %s", node.Right().Type())
	}

	numVal := num.Value
	sidesVal := sides.Value

	rolledDice, rollDiceErr := e.RollDice(numVal, sidesVal)
	if rollDiceErr != nil {
		return nil, rollDiceErr
	}

	return ast.NewSumRollResult(rolledDice), nil
}

type nodeSetter func(ast.Node)

func (e *Evaluator) replaceVariablePrimaryExpr(node ast.Node, setter nodeSetter) error {
	if node.IsVariable() {
		valueDeterminedNode, err := e.determineValueOfVariableExpr(node)
		if err != nil {
			return err
		}

		setter(valueDeterminedNode)
	}

	return nil
}

func (e *Evaluator) determineValuesInCommand(node *ast.Command) error {
	expr := node.Expression
	if expr.IsPrimaryExpression() {
		return e.replaceVariablePrimaryExpr(expr, func(newNode ast.Node) {
			node.Expression = newNode
		})
	}

	return e.DetermineValues(expr)
}

func (e *Evaluator) determineValuesInPrefixExpression(node ast.PrefixExpression) error {
	right := node.Right()
	if right.IsPrimaryExpression() {
		return e.replaceVariablePrimaryExpr(right, func(newNode ast.Node) {
			node.SetRight(newNode)
		})
	}

	return e.DetermineValues(right)
}

func (e *Evaluator) determineValuesInInfixExpression(node ast.InfixExpression) error {
	left := node.Left()
	var leftErr error

	if left.IsPrimaryExpression() {
		leftErr = e.replaceVariablePrimaryExpr(left, func(newNode ast.Node) {
			node.SetLeft(newNode)
		})
	} else {
		leftErr = e.DetermineValues(left)
	}

	if leftErr != nil {
		return leftErr
	}

	right := node.Right()
	if right.IsPrimaryExpression() {
		return e.replaceVariablePrimaryExpr(right, func(newNode ast.Node) {
			node.SetRight(newNode)
		})
	}

	return e.DetermineValues(right)
}
