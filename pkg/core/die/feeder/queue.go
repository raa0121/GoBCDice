package feeder

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/die"
)

// 指定したダイスを取り出せる、キュー型ダイス供給機の構造体。
type Queue struct {
	queue []die.Die
}

// QueueがFeederインターフェースを実装しているかの確認
var _ DieFeeder = (*Queue)(nil)

// NewQueue はキュー型ダイス供給機を返す。
//
// dice: 供給するダイスのスライス
func NewQueue(dice []die.Die) *Queue {
	f := &Queue{queue: []die.Die{}}
	f.Append(dice)

	return f
}

// NewEmptyQueue は、空のキュー型ダイス供給機を返す。
func NewEmptyQueue() *Queue {
	return NewQueue([]die.Die{})
}

// CanSpecifyDie は、供給されるダイスを指定できるかを返す。
// キュー型ダイス供給機ではtrueを返す。
func (f *Queue) CanSpecifyDie() bool {
	return true
}

// Dice は、現在のキューの内容をコピーして返す。
func (f *Queue) Dice() []die.Die {
	dice := []die.Die{}

	for _, d := range f.queue {
		dice = append(dice, d)
	}

	return dice
}

// Next はキューからダイスを1つ取り出して供給する。
// キューが空だった場合はエラーを返す。
func (f *Queue) Next(_ int) (die.Die, error) {
	if f.IsEmpty() {
		return die.Die{}, fmt.Errorf("取り出せるダイスがありません")
	}

	// キューからダイスを取り出す
	d := f.queue[0]
	f.queue = f.queue[1:]

	return d, nil
}

// Push はダイスをキューに追加する。
func (f *Queue) Push(d die.Die) {
	f.queue = append(f.queue, d)
}

// Append は複数のダイスをキューの末尾に追加する。
func (f *Queue) Append(dice []die.Die) {
	for _, d := range dice {
		f.Push(d)
	}
}

// Clear はキューを空にする。
func (f *Queue) Clear() {
	f.queue = []die.Die{}
}

// Set は指定されたダイスをキューに配置する。
func (f *Queue) Set(dice []die.Die) {
	f.Clear()
	f.Append(dice)
}

// Remaining は残りのダイスの数を返す。
func (f *Queue) Remaining() int {
	return len(f.queue)
}

// IsEmpty は、ダイスのキューが空ならばtrueを、空でなければfalseを返す。
func (f *Queue) IsEmpty() bool {
	return f.Remaining() == 0
}
