package feeder

import (
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

func TestQueue_Next(t *testing.T) {
	testcases := []struct {
		dice []die.Die
	}{
		{[]die.Die{{2, 6}}},
		{[]die.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}},
		{[]die.Die{{2, 4}, {3, 6}, {5, 10}, {10, 20}}},
	}

	for i, test := range testcases {
		f := NewQueue(test.dice)

		allDiceTaken := true

		for j, expectedDie := range test.dice {
			actualDie, err := f.Next(expectedDie.Sides)
			if err != nil {
				t.Errorf("#%d-%d: got err: %s", i, j, err)

				allDiceTaken = false
				break
			}

			if !reflect.DeepEqual(actualDie, expectedDie) {
				t.Errorf("#%d-%d: wrong die: got %v, want %v",
					i, j, actualDie, expectedDie)
			}
		}

		if allDiceTaken {
			if !f.IsEmpty() {
				t.Errorf("#%d: %d dice remain", i, f.Remaining())
			}
		}
	}
}

func TestQueue_Append(t *testing.T) {
	dice := []die.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}

	f := NewQueue(dice[0:2])
	if f.Remaining() != 2 {
		t.Fatalf("キューに正しくダイスが入っていません")
	}

	f.Append(dice[2:6])

	for i, expected := range dice {
		actual, err := f.Next(6)
		if err != nil {
			t.Fatalf("#%d: エラー: %s", i, err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("#%d: 異なるダイス: got %s, want %s", i, actual, expected)
		}
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
