package feeder

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/die"
	"reflect"
	"testing"
)

func TestNewEmptyQueue(t *testing.T) {
	f := NewEmptyQueue()

	if !f.IsEmpty() {
		t.Fatalf("空のキューが作れない")
	}
}

func TestQueue_CanSpecifyDie(t *testing.T) {
	f := NewQueue([]die.Die{})

	if !f.CanSpecifyDie() {
		t.Fatalf("Queueはダイスを指定できてなければならない")
	}
}

func TestQueue_Dice_ShouldCopyDice(t *testing.T) {
	testcases := [][]die.Die{
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
	}

	for _, dice := range testcases {
		t.Run(fmt.Sprintf("[%v]", die.FormatDiceWithoutSpaces(dice)), func(t *testing.T) {
			f := NewQueue(dice)

			diceFromQueue := f.Dice()
			diceFromQueue[0] = die.Die{99, 100}

			for i, d := range f.Dice() {
				if !reflect.DeepEqual(d, dice[i]) {
					t.Errorf("Dice() の結果が変わった: got %v, want %v",
						f.Dice(), dice)
					continue
				}
			}
		})
	}
}

func TestQueue_Next(t *testing.T) {
	testcases := [][]die.Die{
		{{2, 6}},
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{2, 4}, {3, 6}, {5, 10}, {10, 20}},
	}

	for _, dice := range testcases {
		t.Run(fmt.Sprintf("[%s]", die.FormatDiceWithoutSpaces(dice)), func(t *testing.T) {
			f := NewQueue(dice)

			gotErr := false
			for _, expectedDie := range dice {
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
	dice := []die.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}

	f := NewQueue(dice[0:2])
	if f.Remaining() != 2 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Append(dice[2:6])

	for _, expected := range dice {
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
	f := NewQueue([]die.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}})

	if f.Remaining() != 6 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Clear()

	if !f.IsEmpty() {
		t.Fatalf("キューを空にできませんでした (残り %d ダイス)", f.Remaining())
	}
}

func TestQueue_Set(t *testing.T) {
	dice := []die.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}

	f := NewQueue(dice[0:2])
	if f.Remaining() != 2 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Set(dice[2:6])

	for _, expected := range dice[2:6] {
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
