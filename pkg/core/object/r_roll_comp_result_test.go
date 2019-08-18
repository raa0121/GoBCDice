package object

import (
	"fmt"
	"testing"
)

func TestRRollCompResult_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *RRollCompResult
		expected string
	}{
		{
			obj: &RRollCompResult{
				ValueGroups: NewArray(
					NewArray(
						NewInteger(3),
						NewInteger(1),
					),
					NewArray(
						NewInteger(2),
					),
				),
				NumOfSuccesses: NewInteger(1),
			},
			expected: "<RRollCompResult Values=[[3, 1], [2]], NumOfSuccesses=1>",
		},
		{
			obj: &RRollCompResult{
				ValueGroups: NewArray(
					NewArray(
						NewInteger(6),
						NewInteger(1),
					),
					NewArray(
						NewInteger(3),
					),
					NewArray(
						NewInteger(1),
					),
				),
				NumOfSuccesses: NewInteger(3),
			},
			expected: "<RRollCompResult Values=[[6, 1], [3], [1]], NumOfSuccesses=3>",
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
