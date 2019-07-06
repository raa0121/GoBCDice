package feeder

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/die"
	"reflect"
	"testing"
)

func TestMT19937_CanSpecifyDie(t *testing.T) {
	f := NewMT19937(1)

	if f.CanSpecifyDie() {
		t.Fatalf("MT19937はダイスを指定できてはならない")
	}
}

func TestMT19937_Seed(t *testing.T) {
	testcases := []int64{1, 2, 20190401}

	for _, expected := range testcases {
		t.Run(fmt.Sprintf("%d", expected), func(t *testing.T) {
			f := NewMT19937(expected)

			if actual := f.Seed(); actual != expected {
				t.Errorf("wrong seed: got %d, want %d", actual, expected)
			}
		})
	}
}

func TestMT19937_Next(t *testing.T) {
	// テストケース
	// シードを固定値に設定するため、必ずこの順番で出るはず
	testcases := [][]die.Die{
		{{4, 6}, {1, 6}, {2, 6}, {6, 6}, {5, 6}, {3, 6}},
		{{2, 2}, {1, 4}, {2, 6}, {4, 10}, {3, 20}},
	}

	for _, dice := range testcases {
		t.Run(fmt.Sprintf("[%s]", die.FormatDiceWithoutSpaces(dice)), func(t *testing.T) {
			f := NewMT19937(1)

			gotErr := false
			for _, expectedDie := range dice {
				if gotErr {
					return
				}

				t.Run(expectedDie.String(), func(t *testing.T) {
					actualDie, err := f.Next(expectedDie.Sides)
					if err != nil {
						t.Fatalf("got err: %s", err)
						return
					}

					if !reflect.DeepEqual(actualDie, expectedDie) {
						t.Errorf("wrong die: got %s, want %s",
							actualDie, expectedDie)
					}
				})
			}
		})
	}
}
