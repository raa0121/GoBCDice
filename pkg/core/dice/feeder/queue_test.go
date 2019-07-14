package feeder

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"reflect"
	"testing"
)

// 供給するダイスを指定する場合の例。
func Example_queue() {
	// 供給するダイスを指定する
	dieFeeder := NewQueue([]dice.Die{{1, 6}, {3, 6}, {5, 6}})

	// 6面ダイスを1個振る
	d, err := dieFeeder.Next(6)
	if err != nil {
		return
	}

	fmt.Println(d.String())
	// Output: <Die 1/6>
}

func TestNewEmptyQueue(t *testing.T) {
	f := NewEmptyQueue()

	if !f.IsEmpty() {
		t.Fatalf("空のキューが作れない")
	}
}

func TestQueue_CanSpecifyDie(t *testing.T) {
	f := NewQueue([]dice.Die{})

	if !f.CanSpecifyDie() {
		t.Fatalf("Queueはダイスを指定できてなければならない")
	}
}

func TestQueue_Dice_ShouldCopyDice(t *testing.T) {
	testcases := [][]dice.Die{
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
	}

	for _, ds := range testcases {
		t.Run("["+dice.FormatDiceWithoutSpaces(ds)+"]", func(t *testing.T) {
			f := NewQueue(ds)

			diceFromQueue := f.Dice()
			diceFromQueue[0] = dice.Die{99, 100}

			for i, d := range f.Dice() {
				if !reflect.DeepEqual(d, ds[i]) {
					t.Errorf("Dice() の結果が変わった: got %v, want %v",
						f.Dice(), ds)
					continue
				}
			}
		})
	}
}

func TestQueue_Next(t *testing.T) {
	testcases := [][]dice.Die{
		{{2, 6}},
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{2, 4}, {3, 6}, {5, 10}, {10, 20}},
	}

	for _, ds := range testcases {
		t.Run(fmt.Sprintf("[%s]", dice.FormatDiceWithoutSpaces(ds)), func(t *testing.T) {
			f := NewQueue(ds)

			gotErr := false
			for _, expectedDie := range ds {
				if gotErr {
					return
				}

				t.Run(expectedDie.String(), func(t *testing.T) {
					actualDie, err := f.Next(expectedDie.Sides)
					if err != nil {
						t.Errorf("got err: %s", err)
						gotErr = true

						return
					}

					if !reflect.DeepEqual(actualDie, expectedDie) {
						t.Errorf("wrong die: got %v, want %v",
							actualDie, expectedDie)
					}
				})
			}

			if !f.IsEmpty() {
				t.Errorf("%d dice remain", f.Remaining())
			}
		})
	}
}

func TestQueue_Append(t *testing.T) {
	ds := []dice.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}

	f := NewQueue(ds[0:2])
	if f.Remaining() != 2 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Append(ds[2:6])

	for _, expected := range ds {
		t.Run(expected.String(), func(t *testing.T) {
			actual, err := f.Next(6)
			if err != nil {
				t.Fatalf("エラー: %s", err)
				return
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("異なるダイス: got %s, want %s", actual, expected)
			}
		})
	}

	if !f.IsEmpty() {
		t.Fatalf("キューにダイスが残っています")
	}
}

func TestQueue_Clear(t *testing.T) {
	f := NewQueue([]dice.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}})

	if f.Remaining() != 6 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Clear()

	if !f.IsEmpty() {
		t.Fatalf("キューを空にできませんでした (残り %d ダイス)", f.Remaining())
	}
}

func TestQueue_Set(t *testing.T) {
	ds := []dice.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}

	f := NewQueue(ds[0:2])
	if f.Remaining() != 2 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Set(ds[2:6])

	for _, expected := range ds[2:6] {
		t.Run(expected.String(), func(t *testing.T) {
			actual, err := f.Next(6)
			if err != nil {
				t.Fatalf("エラー: %s", err)
				return
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("異なるダイス: got %s, want %s", actual, expected)
			}
		})
	}

	if !f.IsEmpty() {
		t.Fatalf("キューにダイスが残っています")
	}
}
