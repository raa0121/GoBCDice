package object

import (
	"fmt"
	"testing"
)

func TestURollExprResult_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *URollExprResult
		expected string
	}{
		{
			obj: &URollExprResult{
				ValueGroups: NewArray(
					NewArray(
						NewInteger(3),
						NewInteger(1),
					),
					NewArray(
						NewInteger(2),
					),
				),
				Modifier: NewInteger(0),
			},
			expected: "<URollExprResult ValueGroups=[[3, 1], [2]], Modifier=0>",
		},
		{
			obj: &URollExprResult{
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
				Modifier: NewInteger(1),
			},
			expected: "<URollExprResult ValueGroups=[[6, 1], [3], [1]], Modifier=1>",
		},
		{
			obj: &URollExprResult{
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
				Modifier: NewInteger(-1),
			},
			expected: "<URollExprResult ValueGroups=[[6, 1], [3], [1]], Modifier=-1>",
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

func TestURollExprResult_MaxValue(t *testing.T) {
	testcases := []struct {
		obj      *URollExprResult
		expected int
	}{
		{
			obj: &URollExprResult{
				ValueGroups: NewArray(
					NewArray(
						NewInteger(3),
						NewInteger(1),
					),
					NewArray(
						NewInteger(2),
					),
				),
				Modifier: NewInteger(0),
			},
			expected: 4,
		},
		{
			obj: &URollExprResult{
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
				Modifier: NewInteger(1),
			},
			expected: 8,
		},
		{
			obj: &URollExprResult{
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
				Modifier: NewInteger(-1),
			},
			expected: 6,
		},
		{
			obj: &URollExprResult{
				// 3,5,3,10[6,4],1,15[6,6,3],5,1
				ValueGroups: NewArray(
					NewArray(
						NewInteger(3),
					),
					NewArray(
						NewInteger(5),
					),
					NewArray(
						NewInteger(3),
					),
					NewArray(
						NewInteger(6),
						NewInteger(4),
					),
					NewArray(
						NewInteger(1),
					),
					NewArray(
						NewInteger(6),
						NewInteger(6),
						NewInteger(3),
					),
					NewArray(
						NewInteger(5),
					),
					NewArray(
						NewInteger(1),
					),
				),
				Modifier: NewInteger(0),
			},
			expected: 15,
		},
	}

	for _, test := range testcases {
		t.Run(test.obj.ValueGroups.InspectWithoutSpaces(), func(t *testing.T) {
			actual := test.obj.MaxValue()
			if actual.Value != test.expected {
				t.Fatalf("got=%d, want=%d", actual.Value, test.expected)
			}
		})
	}
}
