package roller

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"github.com/raa0121/GoBCDice/pkg/core/die/feeder"
)

// ダイスローラーを表す構造体
type DiceRoller struct {
	// feederが実際にダイスを供給する
	feeder feeder.DieFeeder
}

// Newは指定したDieDieFeederを使うDiceRollerを構築して返す
func New(f feeder.DieFeeder) *DiceRoller {
	dr := &DiceRoller{
		feeder: f,
	}

	return dr
}

// DieFeederは指定したDieFeederを返す
func (dr *DiceRoller) DieFeeder() feeder.DieFeeder {
	return dr.feeder
}

// RollDiceは、sides個の面を持つダイスをnum個振り、その結果を返す
func (dr *DiceRoller) RollDice(num int, sides int) ([]die.Die, error) {
	if sides < 1 {
		return nil, fmt.Errorf("ダイスの面数が少なすぎます: %d", sides)
	}

	if num < 1 {
		return nil, fmt.Errorf("振るダイス数が少なすぎます: %d", num)
	}

	dice := []die.Die{}

	for i := 0; i < num; i++ {
		d, err := dr.feeder.Next(sides)
		if err != nil {
			return nil, err
		}

		dice = append(dice, d)
	}

	return dice, nil
}
