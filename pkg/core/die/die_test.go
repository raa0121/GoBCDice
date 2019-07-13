package die

import (
	"fmt"
	"testing"
)

func TestFormatDice(t *testing.T) {
	testcases := []struct {
		dice     []Die
		expected string
	}{
		{
			dice:     []Die{{2, 6}},
			expected: "2/6",
		},
		{
			dice:     []Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
			expected: "1/6, 2/6, 3/6, 4/6, 5/6, 6/6",
		},
		{
			dice:     []Die{{2, 4}, {3, 6}, {5, 10}, {10, 20}},
			expected: "2/4, 3/6, 5/10, 10/20",
		},
	}

	for _, test := range testcases {
		dice := test.dice
		expected := test.expected

		t.Run(fmt.Sprintf("%v", dice), func(t *testing.T) {
			actual := FormatDice(dice)

			if actual != expected {
				t.Fatalf("got %v, want %v", actual, expected)
			}
		})
	}
}

func TestFormatDiceWithoutSpace(t *testing.T) {
	testcases := []struct {
		dice     []Die
		expected string
	}{
		{
			dice:     []Die{{2, 6}},
			expected: "2/6",
		},
		{
			dice:     []Die{{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}},
			expected: "1/6,2/6,3/6,4/6,5/6,6/6",
		},
		{
			dice:     []Die{{2, 4}, {3, 6}, {5, 10}, {10, 20}},
			expected: "2/4,3/6,5/10,10/20",
		},
	}

	for _, test := range testcases {
		dice := test.dice
		expected := test.expected

		t.Run(fmt.Sprintf("%v", dice), func(t *testing.T) {
			actual := FormatDiceWithoutSpaces(dice)

			if actual != expected {
				t.Fatalf("got %v, want %v", actual, expected)
			}
		})
	}
}
