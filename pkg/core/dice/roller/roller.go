/*
ダイスを振る処理のパッケージ。
DieFeederを指定できるため、ダイスの供給方法に依らずに複数個のダイスを振る処理を実行できる。
*/
package roller

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
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
func (dr *DiceRoller) RollDice(num int, sides int) ([]dice.Die, error) {
	if sides < 1 {
		return nil, fmt.Errorf(
			"RollDice(num: %d, sides: %d): ダイスの面数が少なすぎます",
			num,
			sides,
		)
	}

	if num < 1 {
		return nil, fmt.Errorf(
			"RollDice(num: %d, sides: %d): 振るダイス数が少なすぎます",
			num,
			sides,
		)
	}

	// 結果のスライスの領域をnum個分確保する
	rolledDice := make([]dice.Die, 0, num)

	for i := 0; i < num; i++ {
		d, err := dr.feeder.Next(sides)
		if err != nil {
			return nil, err
		}

		if d.Sides != sides {
			return nil, fmt.Errorf(
				"RollDice(num: %d, sides: %d) -> %d/%d: ダイスの面数が指定と一致しません",
				num,
				sides,
				d.Value,
				d.Sides,
			)
		}

		rolledDice = append(rolledDice, d)
	}

	return rolledDice, nil
}
