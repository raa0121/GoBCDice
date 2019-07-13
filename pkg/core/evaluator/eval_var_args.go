package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// EvalVarArgsは可変ノードの引数を評価して確定させる
func (e *Evaluator) EvalVarArgs(node ast.Node) error {
	switch n := node.(type) {
	case ast.Command:
		return e.evalVarArgsInCommand(n)
	case ast.PrefixExpression:
		return e.evalVarArgsInPrefixExpression(n)
	case ast.InfixExpression:
		return e.evalVarArgsInInfixExpression(n)
	}

	return fmt.Errorf("EvalVarArgs not implemented: %s", node.Type())
}

func (e *Evaluator) evalVarArgsOfVariableExpr(node ast.Node) error {
	switch n := node.(type) {
	case *ast.DRoll:
		return e.evalVarArgsOfDRoll(n)
	}

	return fmt.Errorf("evalVarArgsOfVariableExpr not implemented: %s", node.Type())
}

func (e *Evaluator) evalVarArgsOfDRoll(node *ast.DRoll) error {
	leftObj, leftErr := e.Eval(node.Left())
	if leftErr != nil {
		return leftErr
	}

	rightObj, rightErr := e.Eval(node.Right())
	if rightErr != nil {
		return rightErr
	}

	evaluatedLeft := ast.NewInt(leftObj.(*object.Integer).Value, token.Token{})
	evaluatedRight := ast.NewInt(rightObj.(*object.Integer).Value, token.Token{})

	node.SetLeft(evaluatedLeft)
	node.SetRight(evaluatedRight)

	return nil
}

func (e *Evaluator) evalVarArgsInCommand(node ast.Command) error {
	expr := node.Expression()
	if expr.IsPrimaryExpression() {
		if expr.IsVariable() {
			return e.evalVarArgsOfVariableExpr(expr)
		}

		return nil
	}

	return e.EvalVarArgs(expr)
}

func (e *Evaluator) evalVarArgsInPrefixExpression(node ast.PrefixExpression) error {
	right := node.Right()
	if right.IsPrimaryExpression() {
		if right.IsVariable() {
			return e.evalVarArgsOfVariableExpr(right)
		}

		return nil
	}

	return e.EvalVarArgs(right)
}

func (e *Evaluator) evalVarArgsInInfixExpression(node ast.InfixExpression) error {
	left := node.Left()
	var leftErr error

	if left.IsPrimaryExpression() {
		if left.IsVariable() {
			leftErr = e.evalVarArgsOfVariableExpr(left)
		}
	} else {
		leftErr = e.EvalVarArgs(left)
	}

	if leftErr != nil {
		return leftErr
	}

	right := node.Right()
	if right.IsPrimaryExpression() {
		if right.IsVariable() {
			return e.evalVarArgsOfVariableExpr(right)
		}

		return nil
	}

	return e.EvalVarArgs(right)
}
