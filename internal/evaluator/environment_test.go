package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/die"
	"reflect"
	"testing"
)

func TestEnvironment_RolledDice_ShouldCopyDice(t *testing.T) {
	testcases := [][]die.Die{
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
	}

	for _, dice := range testcases {
		t.Run(fmt.Sprintf("%v", dice), func(t *testing.T) {
			env := NewEnvironment()
			env.AppendRolledDice(dice)

			diceFromRolledDice := env.RolledDice()
			diceFromRolledDice[0] = die.Die{99, 100}

			for i, d := range env.RolledDice() {
				if !reflect.DeepEqual(d, dice[i]) {
					t.Errorf("RolledDice() の結果が変わった: got %v, want %v",
						env.RolledDice(), dice)
					continue
				}
			}
		})
	}
}

func TestEnvironment_ClearRolledDice(t *testing.T) {
	testcases := [][]die.Die{
		{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
		{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
	}

	for _, dice := range testcases {
		t.Run(fmt.Sprintf("%v", dice), func(t *testing.T) {
			env := NewEnvironment()

			env.AppendRolledDice(dice)
			env.ClearRolledDice()

			if len(env.RolledDice()) > 0 {
				t.Error("記録されたダイスロール結果がクリアされていない")
			}
		})
	}
}
