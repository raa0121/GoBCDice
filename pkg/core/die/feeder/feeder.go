package feeder

import (
	"github.com/raa0121/GoBCDice/pkg/core/die"
)

// DieFeederは、ダイス供給機のインターフェース
type DieFeeder interface {
	// Nextはダイスを1つ供給する
	//
	// sides: ダイスの面の数
	Next(sides int) (die.Die, error)

	// CanSpecifyDieは、供給されるダイスを指定できるかを返す
	CanSpecifyDie() bool
}
