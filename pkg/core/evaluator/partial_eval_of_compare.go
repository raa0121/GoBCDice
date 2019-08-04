package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// EvalCompareVarArgsAndRight は、比較式の左辺の可変ノードの引数および右辺を評価する。
func (e *Evaluator) EvalCompareVarArgsAndRight(node *ast.Compare) error {
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

	evaluatedRight := ast.NewInt(rightObj.(*object.Integer).Value, token.Token{})
	node.SetRight(evaluatedRight)

	return nil
}

// EvalCompareLeft は比較式の左辺を評価する。
func (e *Evaluator) EvalCompareLeft(node *ast.Compare) (object.Object, error) {
	leftObj, leftEvalErr := e.Eval(node.Left())
	if leftEvalErr != nil {
		return nil, leftEvalErr
	}

	evaluatedLeft := ast.NewInt(leftObj.(*object.Integer).Value, token.Token{})
	node.SetLeft(evaluatedLeft)

	return leftObj, nil
}
