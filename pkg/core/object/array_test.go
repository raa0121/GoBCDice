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

func (a *Array) TestArray_Length(t *testing.T) {
	testcases := []struct {
		array    *Array
		expected int
	}{
		{
			array:    NewArray(),
			expected: 0,
		},
		{
			array: NewArray(
				NewInteger(1),
				NewInteger(2),
			),
			expected: 2,
		},
		{
			array: NewArray(
				NewInteger(1),
				NewInteger(2),
				NewInteger(3),
				NewInteger(4),
				NewInteger(5),
			),
			expected: 5,
		},
	}

	for _, test := range testcases {
		t.Run(test.array.InspectWithoutSpaces(), func(t *testing.T) {
			actual := test.array.Length()
			if actual != test.expected {
				t.Fatalf("got=%d, want=%d", actual, test.expected)
			}
		})
	}
}

func (a *Array) TestArray_At(t *testing.T) {
	array := NewArray(
		NewInteger(2),
		NewInteger(3),
		NewInteger(5),
		NewInteger(8),
		NewInteger(13),
	)

	x1 := 1
	x2 := 2
	for i := 0; i < 5; i++ {
		actual := array.At(i).(*Integer)
		if actual.Value != x2 {
			t.Fatalf("[%d]: got=%d, want=%d", i, actual, x2)
		}

		temp := x1
		x1 = x2
		x2 += temp
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

func TestArray_SumOfIntegers(t *testing.T) {
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
			expectedValue: 3,
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
			expectedValue: 31,
			expectedOK:    true,
		},
	}

	for _, test := range testcases {
		t.Run(test.obj.InspectWithoutSpaces(), func(t *testing.T) {
			actualInt, actualOK := test.obj.SumOfIntegers()
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
