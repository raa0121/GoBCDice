package object

import (
	"fmt"
	"testing"
)

func TestArray_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *Array
		expected string
	}{
		{
			obj:      NewArray(),
			expected: "[]",
		},
		{
			obj: NewArray(
				NewInteger(1),
				NewInteger(2),
			),
			expected: "[1, 2]",
		},
		{
			obj: NewArray(
				NewInteger(2),
				NewInteger(3),
				NewInteger(5),
				NewInteger(8),
				NewInteger(13),
			),
			expected: "[2, 3, 5, 8, 13]",
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

func TestArray_InspectWithoutSpaces(t *testing.T) {
	testcases := []struct {
		obj      *Array
		expected string
	}{
		{
			obj:      NewArray(),
			expected: "[]",
		},
		{
			obj: NewArray(
				NewInteger(1),
				NewInteger(2),
			),
			expected: "[1,2]",
		},
		{
			obj: NewArray(
				NewInteger(2),
				NewInteger(3),
				NewInteger(5),
				NewInteger(8),
				NewInteger(13),
			),
			expected: "[2,3,5,8,13]",
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.expected), func(t *testing.T) {
			actual := test.obj.InspectWithoutSpaces()
			if actual != test.expected {
				t.Fatalf("got=%q, want=%q", actual, test.expected)
			}
		})
	}
}

func TestArray_JoinedElements(t *testing.T) {
	testcases := []struct {
		obj      *Array
		sep      string
		expected string
	}{
		{
			obj:      NewArray(),
			expected: "",
		},
		{
			obj: NewArray(
				NewInteger(1),
				NewInteger(2),
			),
			sep:      ",",
			expected: "1,2",
		},
		{
			obj: NewArray(
				NewInteger(2),
				NewInteger(3),
				NewInteger(5),
				NewInteger(8),
				NewInteger(13),
			),
			sep:      ",",
			expected: "2,3,5,8,13",
		},
		{
			obj: NewArray(
				NewInteger(1),
				NewInteger(2),
			),
			sep:      "/",
			expected: "1/2",
		},
		{
			obj: NewArray(
				NewInteger(2),
				NewInteger(3),
				NewInteger(5),
				NewInteger(8),
				NewInteger(13),
			),
			sep:      "/",
			expected: "2/3/5/8/13",
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.expected), func(t *testing.T) {
			actual := test.obj.JoinedElements(test.sep)
			if actual != test.expected {
				t.Fatalf("got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
