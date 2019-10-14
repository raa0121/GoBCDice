package object

import (
	"fmt"
	"testing"
)

func TestURollCompResult_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *URollCompResult
		expected string
	}{
		{
			obj: NewURollCompResult(
				NewURollExprResult(
					NewArray(
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
					NewInteger(1),
				),
				NewInteger(2),
			),
			expected: "<URollCompResult RollResult=<URollExprResult ValueGroups=[[6, 1], [3], [1]], Modifier=1>, NumOfSuccesses=2>",
		},
		{
			obj: NewURollCompResult(
				NewURollExprResult(
					NewArray(
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
					NewInteger(1),
				),
				NewInteger(0),
			),
			expected: "<URollCompResult RollResult=<URollExprResult ValueGroups=[[6, 1], [3], [1]], Modifier=1>, NumOfSuccesses=0>",
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.expected), func(t *testing.T) {
			actual := test.obj.Inspect()
			if actual != test.expected {
				t.Fatalf("got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
