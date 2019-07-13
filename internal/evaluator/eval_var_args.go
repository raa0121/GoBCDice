package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/object"
	"github.com/raa0121/GoBCDice/internal/token"
)

// EvalVarArgsは可変ノードの引数を評価して確定させる
func (e *Evaluator) EvalVarArgs(node ast.Node) error {
	if !node.IsVariable() {
		return nil
	}

	switch n := node.(type) {
	case *ast.DRoll:
		return e.evalVarArgsOfDRoll(n)
	case ast.Command:
		return e.evalVarArgsOfCommand(n)
	case ast.PrefixExpression:
		return e.evalVarArgsOfPrefixExpression(n)
	case ast.InfixExpression:
		return e.evalVarArgsOfInfixExpression(n)
	}

	return fmt.Errorf("EvalVarArgs not implemented: %s", node.Type())
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

func (e *Evaluator) evalVarArgsOfCommand(node ast.Command) error {
	err := e.EvalVarArgs(node.Expression())
	if err != nil {
		return err
	}

	return nil
}

func (e *Evaluator) evalVarArgsOfPrefixExpression(node ast.PrefixExpression) error {
	err := e.EvalVarArgs(node.Right())
	if err != nil {
		return err
	}

	return nil
}

func (e *Evaluator) evalVarArgsOfInfixExpression(node ast.InfixExpression) error {
	leftErr := e.EvalVarArgs(node.Left())
	if leftErr != nil {
		return leftErr
	}

	rightErr := e.EvalVarArgs(node.Right())
	if rightErr != nil {
		return rightErr
	}

	return nil
}
