package feeder

import (
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"github.com/seehuhn/mt19937"
	"math/rand"
	"time"
)

// ランダムにダイスを取り出すダイス供給機の構造体
// Ruby版BCDiceと同様にメルセンヌ・ツイスタを使用する
type MT19937 struct {
	seed int64
	rng  *rand.Rand
}

// MT19937がFeederインターフェースを実装しているかの確認
var _ DieFeeder = (*MT19937)(nil)

// NewMT19937は、シードを指定したMT19937ダイス供給機を返す
func NewMT19937(seed int64) *MT19937 {
	f := &MT19937{
		seed: seed,
		rng:  rand.New(mt19937.New()),
	}

	f.rng.Seed(seed)

	return f
}

// NewMT19937WithSeedFromTimeは、現在の時刻をシードとした
// MT19937ダイス供給機を返す
func NewMT19937WithSeedFromTime() *MT19937 {
	return NewMT19937(time.Now().UnixNano())
}

// CanSpecifyDieは、供給されるダイスを指定できるかを返す
func (f *MT19937) CanSpecifyDie() bool {
	return false
}

// Seedは設定されているシードを返す
func (f *MT19937) Seed() int64 {
	return f.seed
}

// Nextはランダムな値のダイスを1つ供給する
//
// sides: ダイスの面の数
func (f *MT19937) Next(sides int) (die.Die, error) {
	d := die.Die{
		Sides: sides,
		Value: 1 + f.rng.Intn(sides),
	}
	return d, nil
}
