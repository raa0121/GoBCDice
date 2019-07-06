package evaluator

import (
	"github.com/raa0121/GoBCDice/internal/die"
)

// コマンド評価の環境を表す構造体
type Environment struct {
	rolledDice []die.Die
}

// NewEnvironmentは新しいコマンド評価環境を返す
func NewEnvironment() *Environment {
	return &Environment{
		rolledDice: []die.Die{},
	}
}

// RolledDiceは、記録されたダイスロール結果を返す
func (e *Environment) RolledDice() []die.Die {
	// ダイスロール結果のコピー先
	dice := []die.Die{}

	for _, d := range e.rolledDice {
		newDie := d
		dice = append(dice, newDie)
	}

	return dice
}

func (e *Environment) PushRolledDie(d die.Die) {
	e.rolledDice = append(e.rolledDice, d)
}

func (e *Environment) AppendRolledDice(dice []die.Die) {
	for _, d := range dice {
		e.PushRolledDie(d)
	}
}

// ClearRolledDiceは、記録されたダイスロール結果をクリアする
func (e *Environment) ClearRolledDice() {
	e.rolledDice = []die.Die{}
}
