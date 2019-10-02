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
		{
			obj: NewArray(
				NewArray(
					NewInteger(1),
					NewInteger(2),
				),
				NewArray(
					NewInteger(3),
					NewInteger(4),
				),
			),
			expected: "[[1,2],[3,4]]",
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

func TestArray_MaxInteger(t *testing.T) {
	testcases := []struct {
		obj           *Array
		expectedValue int
		expectedOK    bool
	}{
		{
			obj:           NewArray(),
			expectedValue: 0,
			expectedOK:    true,
		},
		{
			obj: NewArray(
				NewInteger(1),
				NewString("str"),
			),
			expectedOK: false,
		},
		{
			obj: NewArray(
				NewInteger(1),
			),
			expectedValue: 1,
			expectedOK:    true,
		},
		{
			obj: NewArray(
				NewInteger(1),
				NewInteger(2),
			),
			expectedValue: 2,
			expectedOK:    true,
		},
		{
			obj: NewArray(
				NewInteger(2),
				NewInteger(1),
			),
			expectedValue: 2,
			expectedOK:    true,
		},
		{
			obj: NewArray(
				NewInteger(2),
				NewInteger(3),
				NewInteger(5),
				NewInteger(8),
				NewInteger(13),
			),
			expectedValue: 13,
			expectedOK:    true,
		},
		{
			obj: NewArray(
				NewInteger(3),
				NewInteger(2),
				NewInteger(8),
				NewInteger(13),
				NewInteger(5),
			),
			expectedValue: 13,
			expectedOK:    true,
		},
	}

	for _, test := range testcases {
		t.Run(test.obj.InspectWithoutSpaces(), func(t *testing.T) {
			actualInt, actualOK := test.obj.MaxInteger()
			if !actualOK {
				if test.expectedOK {
					t.Fatal("not OK")
				}

				return
			}

			if !test.expectedOK {
				t.Fatal("unexpected OK")
				return
			}

			if actualInt.Value != test.expectedValue {
				t.Errorf("got=%d, want=%d", actualInt.Value, test.expectedValue)
			}
		})
	}
}
