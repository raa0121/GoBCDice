package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// EvalVarArgs は可変ノードの引数を評価して整数に変換する。
func (e *Evaluator) EvalVarArgs(node ast.Node) error {
	switch n := node.(type) {
	case *ast.BRollList:
		return e.evalVarArgsInBRollList(n)
	case ast.Command:
		return e.evalVarArgsInCommand(n)
	case ast.PrefixExpression:
		return e.evalVarArgsInPrefixExpression(n)
	case ast.InfixExpression:
		return e.evalVarArgsInInfixExpression(n)
	}

	return fmt.Errorf("EvalVarArgs not implemented: %s", node.Type())
}

// evalVarArgsOfVariableExpr は可変ノードの引数を評価して整数に変換する。
//
// このメソッドには、DRollなど、実際に引数を評価して整数に変換する必要がある、可変一次式のノードを渡す。
// このメソッドは、ノードの型に合わせて処理を振り分ける。
func (e *Evaluator) evalVarArgsOfVariableExpr(node ast.Node) error {
	switch n := node.(type) {
	case *ast.DRoll:
		return e.evalVarArgsOfRoll(n)
	case *ast.BRoll:
		return e.evalVarArgsOfRoll(n)
	}

	return fmt.Errorf("evalVarArgsOfVariableExpr not implemented: %s", node.Type())
}

// evalVarArgsOfRoll はダイスロールノードの引数を評価して整数に変換する。
func (e *Evaluator) evalVarArgsOfRoll(node ast.InfixExpression) error {
	leftObj, leftErr := e.Eval(node.Left())
	if leftErr != nil {
		return leftErr
	}

	rightObj, rightErr := e.Eval(node.Right())
	if rightErr != nil {
		return rightErr
	}

	evaluatedLeft := ast.NewInt(leftObj.(*object.Integer).Value)
	evaluatedRight := ast.NewInt(rightObj.(*object.Integer).Value)

	node.SetLeft(evaluatedLeft)
	node.SetRight(evaluatedRight)

	return nil
}

// evalVarArgsInBRollList はバラバラロールリスト内の可変ノードの引数を評価して整数に変換する。
func (e *Evaluator) evalVarArgsInBRollList(node *ast.BRollList) error {
	for _, b := range node.BRolls {
		err := e.evalVarArgsOfVariableExpr(b)
		if err != nil {
			return err
		}
	}

	return nil
}

// evalVarArgsInCommand はコマンドノード内の可変ノードの引数を評価して整数に変換する。
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

// evalVarArgsInCommand は前置式内の可変ノードの引数を評価して整数に変換する。
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

// evalVarArgsInCommand は中置式内の可変ノードの引数を評価して整数に変換する。
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
