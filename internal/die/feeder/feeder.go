package feeder

import (
	"github.com/raa0121/GoBCDice/internal/die"
)

// DieFeederは、サイコロ供給機のインターフェース
type DieFeeder interface {
	// Nextはサイコロを1つ供給する
	//
	// sides: サイコロの面の数
	Next(sides int) (die.Die, error)
}
