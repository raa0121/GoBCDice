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
	case *ast.RRollList:
		return e.evalVarArgsInRRollList(n)
	case *ast.URollExpr:
		return e.evalVarArgsInURollExpr(n)
	case *ast.Command:
		return e.evalVarArgsInCommand(n)
	case ast.PrefixExpression:
		return e.evalVarArgsInPrefixExpression(n)
	case ast.InfixExpression:
		if n.Type() == ast.COMPARE_NODE {
			return e.evalVarArgsInCompare(n)
		} else {
			return e.evalVarArgsInInfixExpression(n)
		}
	}

	return fmt.Errorf("EvalVarArgs not implemented: %s", node.Type())
}

// evalVarArgsOfVariableExpr は可変ノードの引数を評価して整数に変換する。
//
// このメソッドには、DRollなど、実際に引数を評価して整数に変換する必要がある、可変一次式のノードを渡す。
// このメソッドは、ノードの型に合わせて処理を振り分ける。
func (e *Evaluator) evalVarArgsOfVariableExpr(node ast.Node) error {
	switch node.Type() {
	case ast.D_ROLL_NODE, ast.B_ROLL_NODE, ast.R_ROLL_NODE, ast.U_ROLL_NODE:
		return e.evalVarArgsOfRoll(node.(ast.InfixExpression))
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

	node.SetLeft(objectToIntNode(leftObj))
	node.SetRight(objectToIntNode(rightObj))

	return nil
}

// evalVarArgsInBRollList はバラバラロール列内の可変ノードの引数を評価して整数に変換する。
func (e *Evaluator) evalVarArgsInBRollList(node *ast.BRollList) error {
	for _, b := range node.BRolls {
		err := e.evalVarArgsOfVariableExpr(b)
		if err != nil {
			return err
		}
	}

	return nil
}

// evalVarArgsInRRollList は個数振り足しロール列内の可変ノードの引数を評価して整数に変換する。
func (e *Evaluator) evalVarArgsInRRollList(node *ast.RRollList) error {
	// 振り足しの閾値を評価する
	if !node.Threshold.IsNil() {
		thresholdObj, err := e.Eval(node.Threshold)
		if err != nil {
			return err
		}

		node.Threshold = objectToIntNode(thresholdObj)
	}

	// 個数振り足しロールの引数を評価して整数に変換する
	for _, r := range node.RRolls {
		err := e.evalVarArgsOfVariableExpr(r)
		if err != nil {
			return err
		}
	}

	return nil
}

// evalVarArgsInURollExpr は上方無限ロール式内の可変ノードの引数を評価して整数に変換する。
func (e *Evaluator) evalVarArgsInURollExpr(node *ast.URollExpr) error {
	uRollList := node.URollList

	// 振り足しの閾値を評価する
	if !uRollList.Threshold.IsNil() {
		thresholdObj, err := e.Eval(uRollList.Threshold)
		if err != nil {
			return err
		}

		uRollList.Threshold = objectToIntNode(thresholdObj)
	}

	// 上方無限ロールの引数を評価して整数に変換する
	for _, r := range uRollList.RRolls {
		err := e.evalVarArgsOfVariableExpr(r)
		if err != nil {
			return err
		}
	}

	// ボーナスを評価して整数に変換する
	if node.Bonus != nil {
		bonusObj, err := e.evalInfixExpression(node.Bonus)
		if err != nil {
			return err
		}

		bonusValue := bonusObj.(*object.Integer).Value
		var newBonus ast.InfixExpression
		if bonusValue == 0 {
			newBonus = nil
		} else if bonusValue > 0 {
			newBonus = ast.NewAdd(ast.NewInt(0), ast.NewInt(bonusValue))
		} else {
			newBonus = ast.NewSubtract(ast.NewInt(0), ast.NewInt(-bonusValue))
		}

		node.Bonus = newBonus
	}

	return nil
}

// evalVarArgsInCommand はコマンドノード内の可変ノードの引数を評価して整数に変換する。
func (e *Evaluator) evalVarArgsInCommand(node *ast.Command) error {
	expr := node.Expression
	if expr.IsPrimaryExpression() {
		if expr.IsVariable() {
			return e.evalVarArgsOfVariableExpr(expr)
		}

		return nil
	}

	return e.EvalVarArgs(expr)
}

// evalVarArgsInCompare は、比較式の左辺の可変ノードの引数および右辺を評価する。
func (e *Evaluator) evalVarArgsInCompare(node ast.InfixExpression) error {
	// 左辺の可変ノードの引数を評価して整数に変換する
	leftEvalErr := e.EvalVarArgs(node.Left())
	if leftEvalErr != nil {
		return leftEvalErr
	}

	// 右辺（目標値）を評価して整数に変換する
	rightObj, rightEvalErr := e.Eval(node.Right())
	if rightEvalErr != nil {
		return rightEvalErr
	}

	node.SetRight(objectToIntNode(rightObj))

	return nil
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
