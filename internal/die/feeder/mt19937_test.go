package feeder

import (
	"github.com/raa0121/GoBCDice/internal/die"
	"reflect"
	"testing"
)

func TestMT19937_Seed(t *testing.T) {
	testcases := []int64{1, 2, 20190401}

	for i, expected := range testcases {
		f := newMT19937(expected)

		if actual := f.Seed(); actual != expected {
			t.Errorf("#%d: wrong seed: got %d, want %d", i, actual, expected)
		}
	}
}

func TestMT19937_Next(t *testing.T) {
	// テストケース
	// シードを固定値に設定するため、必ずこの順番で出るはず
	testcases := [][]die.Die{
		{{4, 6}, {1, 6}, {2, 6}, {6, 6}, {5, 6}, {3, 6}},
		{{2, 2}, {1, 4}, {2, 6}, {4, 10}, {3, 20}},
	}

	for i, test := range testcases {
		f := newMT19937(1)

		for j, expectedDie := range test {
			actualDie, err := f.Next(expectedDie.Sides)
			if err != nil {
				t.Errorf("#%d-%d: got err: %s", i, j, err)
				break
			}

			if !reflect.DeepEqual(actualDie, expectedDie) {
				t.Errorf("#%d-%d: wrong die: got %s, want %s",
					i, j, actualDie, expectedDie)
			}
		}
	}
}
