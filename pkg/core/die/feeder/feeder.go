/*
1個1個のダイスの供給の仕方を定義するパッケージ。
このパッケージに含まれる構造体を利用することで、ダイスの値をランダムにするか、指定したものにするかを切り替えることができる。

ダイスの値をランダムにする場合は、MT19937を使用する。
ダイスの値を指定したものにする場合は、Queueを使用する。
*/
package feeder

import (
	"github.com/raa0121/GoBCDice/pkg/core/die"
)

// DieFeeder は、ダイス供給機のインターフェース。
type DieFeeder interface {
	// Next はダイスを1つ供給する。
	//
	// sides: ダイスの面数
	Next(sides int) (die.Die, error)

	// CanSpecifyDie は、供給されるダイスを指定できるかを返す。
	CanSpecifyDie() bool
}
