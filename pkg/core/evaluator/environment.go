package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice"
)

// コマンド評価の環境を表す構造体。
type Environment struct {
	rolledDice []dice.Die
}

// NewEnvironment は新しいコマンド評価環境を返す。
func NewEnvironment() *Environment {
	return &Environment{
		rolledDice: []dice.Die{},
	}
}

// RolledDice は記録されたダイスロール結果を返す。
func (e *Environment) RolledDice() []dice.Die {
	// ダイスロール結果のコピー先
	dice := []dice.Die{}

	for _, d := range e.rolledDice {
		newDie := d
		dice = append(dice, newDie)
	}

	return dice
}

// PushRolledDie は振られたダイスを記録に追加する。
func (e *Environment) PushRolledDie(d dice.Die) {
	e.rolledDice = append(e.rolledDice, d)
}

// AppendRolledDice は振られたダイスの列を記録に追加する。
func (e *Environment) AppendRolledDice(dice []dice.Die) {
	for _, d := range dice {
		e.PushRolledDie(d)
	}
}

// ClearRolledDice は記録されたダイスロール結果をクリアする。
func (e *Environment) ClearRolledDice() {
	e.rolledDice = []dice.Die{}
}
