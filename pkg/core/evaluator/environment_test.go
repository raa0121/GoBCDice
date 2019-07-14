package evaluator

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"reflect"
	"testing"
)

func TestEnvironment_RolledDice_ShouldCopyDice(t *testing.T) {
	testcases := [][]dice.Die{
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
	}

	for _, ds := range testcases {
		t.Run("["+dice.FormatDiceWithoutSpaces(ds)+"]", func(t *testing.T) {
			env := NewEnvironment()
			env.AppendRolledDice(ds)

			diceFromRolledDice := env.RolledDice()
			diceFromRolledDice[0] = dice.Die{99, 100}

			for i, d := range env.RolledDice() {
				if !reflect.DeepEqual(d, ds[i]) {
					t.Errorf("RolledDice() の結果が変わった: got %v, want %v",
						env.RolledDice(), ds)
					continue
				}
			}
		})
	}
}

func TestEnvironment_ClearRolledDice(t *testing.T) {
	testcases := [][]dice.Die{
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
	}

	for _, ds := range testcases {
		t.Run("["+dice.FormatDiceWithoutSpaces(ds)+"]", func(t *testing.T) {
			env := NewEnvironment()

			env.AppendRolledDice(ds)
			env.ClearRolledDice()

			if len(env.RolledDice()) > 0 {
				t.Error("記録されたダイスロール結果がクリアされていない")
			}
		})
	}
}
