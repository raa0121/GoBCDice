package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalChoice はランダム選択を評価する。
func (e *Evaluator) evalChoice(node *ast.Choice) (*object.String, error) {
	rolledDice, err := e.RollDice(1, len(node.Items))
	if err != nil {
		return nil, err
	}

	index := rolledDice[0].Value - 1
	value := node.Items[index].Value

	return object.NewString(value), nil
}
