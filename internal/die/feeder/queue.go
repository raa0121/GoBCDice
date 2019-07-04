package feeder

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/die"
)

// 指定したサイコロを取り出せる、キュー型サイコロ供給機の構造体
type Queue struct {
	queue []die.Die
}

// QueueがDieFeederインターフェースを実装しているかの確認
var _ DieFeeder = (*Queue)(nil)

// newQueueはキュー型サイコロ供給機を返す
//
// dice: 供給するサイコロのスライス
func newQueue(dice []die.Die) *Queue {
	f := &Queue{queue: make([]die.Die, 0)}

	for _, d := range dice {
		f.Push(d)
	}

	return f
}

// Nextはリストからサイコロを1つ取り出して供給する
func (f *Queue) Next(_ int) (die.Die, error) {
	if f.IsEmpty() {
		return die.Die{}, fmt.Errorf("no die to take")
	}

	// キューからサイコロを取り出す
	d := f.queue[0]
	f.queue = f.queue[1:]

	return d, nil
}

// Pushはサイコロdをキューに追加する
func (f *Queue) Push(d die.Die) {
	f.queue = append(f.queue, d)
}

// Remainingは残りのサイコロの数を返す
func (f *Queue) Remaining() int {
	return len(f.queue)
}

// IsEmptyは、サイコロのキューが空ならばtrueを、空でなければfalseを返す
func (f *Queue) IsEmpty() bool {
	return f.Remaining() == 0
}
