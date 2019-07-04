package feeder

import (
	"github.com/raa0121/GoBCDice/internal/die"
	"reflect"
	"testing"
)

func TestQueue_Next(t *testing.T) {
	testcases := []struct {
		dice []die.Die
	}{
		{[]die.Die{{2, 6}}},
		{[]die.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}}},
		{[]die.Die{{2, 4}, {3, 6}, {5, 10}, {10, 20}}},
	}

	for i, test := range testcases {
		f := newQueue(test.dice)

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
