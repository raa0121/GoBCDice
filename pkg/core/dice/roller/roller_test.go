package roller

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"reflect"
	"testing"
)

func ExampleDiceRoller_RollDice_mT19937() ([]dice.Die, error) {
	dieFeeder := feeder.NewMT19937WithSeedFromTime()
	dieRoller := New(dieFeeder)

	// 6面ダイスを2個振る
	rolledDice, err := dieRoller.RollDice(2, 6)
	if err != nil {
		return nil, err
	}

	return rolledDice, nil
}

func ExampleDiceRoller_RollDice_queue() ([]dice.Die, error) {
	dieFeeder := feeder.NewQueue([]dice.Die{{1, 6}, {3, 6}, {5, 6}})
	dieRoller := New(dieFeeder)

	// 6面ダイスを3個振る
	rolledDice, err := dieRoller.RollDice(3, 6)
	if err != nil {
		return nil, err
	}

	return rolledDice, nil
}

func TestDiceRoller_RollDice_Queue(t *testing.T) {
	testcases := []struct {
		dice  []dice.Die
		num   int
		sides int
		err   bool
	}{
		{
			dice:  []dice.Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
			num:   6,
			sides: 6,
			err:   false,
		},
		{
			dice:  []dice.Die{{1, 4}, {2, 4}, {3, 4}, {4, 4}},
			num:   3,
			sides: 4,
			err:   false,
		},
		{
			dice:  []dice.Die{{1, 6}, {2, 6}, {3, 6}},
			num:   4,
			sides: 6,
			err:   true,
		},
		{
			dice:  []dice.Die{},
			num:   0,
			sides: 6,
			err:   true,
		},
		{
			dice:  []dice.Die{},
			num:   -1,
			sides: 6,
			err:   true,
		},
		{
			dice:  []dice.Die{{1, -6}},
			num:   1,
			sides: -6,
			err:   true,
		},

		// 指定した面数と取り出されたダイスの面数が一致しなければエラーとする
		// （どどんとふのテストとの互換性維持のため）
		{
			dice:  []dice.Die{{3, 6}, {4, 8}},
			num:   2,
			sides: 6,
			err:   true,
		},
	}

	for i, test := range testcases {
		f := feeder.NewQueue(test.dice)
		dr := New(f)

		actualDice, err := dr.RollDice(test.num, test.sides)
		if err != nil {
			if !test.err {
				t.Errorf("#%d: got err: %s", i, err)
			}

			continue
		}

		if test.err {
			t.Errorf("#%d: should err", i)
			continue
		}

		expectedDice := test.dice[0:test.num]
		expectedRemaining := len(test.dice) - test.num

		if !reflect.DeepEqual(actualDice, expectedDice) {
			t.Errorf("#%d: wrong dice: got %v, want %v", i, actualDice, expectedDice)
		}

		if f.Remaining() != expectedRemaining {
			t.Errorf("#%d: wrong number of remaining dice: got %d, want %d",
				i, f.Remaining(), expectedRemaining)
		}
	}
}

func TestDiceRoller_RollDice_MT19937(t *testing.T) {
	nums := []int{1, 1, 2, 3, 5, 8, 13, 100, 89, 55, 34}

	f := feeder.NewMT19937(1)
	dr := New(f)

	for i, num := range nums {
		rolledDice, err := dr.RollDice(num, 6)
		if err != nil {
			t.Errorf("#%d: got err: %s", i, err)
			continue
		}

		if actual := len(rolledDice); actual != num {
			t.Errorf("#%d: wrong number of dice: got %d dice, want %d dice",
				i, actual, num)
		}
	}
}
