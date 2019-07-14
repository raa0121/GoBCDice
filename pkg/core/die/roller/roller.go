/*
ダイスを振る処理のパッケージ。
DieFeederを指定できるため、ダイスの供給方法に依らずに複数個のダイスを振る処理を実行できる。
*/
package roller

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"github.com/raa0121/GoBCDice/pkg/core/die/feeder"
)

// ダイスローラーを表す構造体。
type DiceRoller struct {
	// feederが実際にダイスを供給する
	feeder feeder.DieFeeder
}

// New は指定したDieFeederを使うDiceRollerを構築して返す。
func New(f feeder.DieFeeder) *DiceRoller {
	dr := &DiceRoller{
		feeder: f,
	}

	return dr
}

// DieFeeder は指定したDieFeederを返す。
func (dr *DiceRoller) DieFeeder() feeder.DieFeeder {
	return dr.feeder
}

// RollDice は、sides個の面を持つダイスをnum個振り、その結果を返す。
//
// num、sidesともに正の整数でなければならない。
// この条件が満たされていなかった場合は、エラーを返す。
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
