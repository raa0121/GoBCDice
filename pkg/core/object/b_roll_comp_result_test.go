package object

import (
	"fmt"
	"testing"
)

func TestBRollCompResult_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *BRollCompResult
		expected string
	}{
		{
			obj: &BRollCompResult{
				Values: NewArray(
					NewInteger(3),
					NewInteger(4),
				),
				NumOfSuccesses: NewInteger(1),
			},
			expected: "<BRollCompResult Values=[3, 4], NumOfSuccesses=1>",
		},
		{
			obj: &BRollCompResult{
				Values: NewArray(
					NewInteger(3),
					NewInteger(4),
					NewInteger(9),
					NewInteger(7),
					NewInteger(1),
					NewInteger(5),
				),
				NumOfSuccesses: NewInteger(3),
			},
			expected: "<BRollCompResult Values=[3, 4, 9, 7, 1, 5], NumOfSuccesses=3>",
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.expected), func(t *testing.T) {
			actual := test.obj.Inspect()
			if actual != test.expected {
				t.Fatalf("got=%s, want=%s", actual, test.expected)
			}
		})
	}
}
