package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalBRollList はバラバラロール列を評価する。
func (e *Evaluator) evalBRollList(node *ast.BRollList) (*object.Array, error) {
	elements := []object.Object{}

	for _, b := range node.BRolls {
		o, err := e.Eval(b)
		if err != nil {
			return nil, err
		}

		intObjs := o.(*object.Array)
		elements = append(elements, intObjs.Elements...)
	}

	return object.NewArrayByMove(elements), nil
}
