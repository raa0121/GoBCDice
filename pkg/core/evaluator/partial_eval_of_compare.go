package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// EvalCompareLeft は比較式の左辺を評価する。
func (e *Evaluator) EvalCompareLeft(node *ast.BasicInfixExpression) (object.Object, error) {
	leftObj, leftEvalErr := e.Eval(node.Left())
	if leftEvalErr != nil {
		return nil, leftEvalErr
	}

	node.SetLeft(objectToIntNode(leftObj))

	return leftObj, nil
}
