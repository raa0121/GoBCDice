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
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 1},
					&Integer{Value: 2},
				},
			},
			expected: "[1, 2]",
		},
		{
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 2},
					&Integer{Value: 3},
					&Integer{Value: 5},
					&Integer{Value: 8},
					&Integer{Value: 13},
				},
			},
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
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 1},
					&Integer{Value: 2},
				},
			},
			expected: "[1,2]",
		},
		{
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 2},
					&Integer{Value: 3},
					&Integer{Value: 5},
					&Integer{Value: 8},
					&Integer{Value: 13},
				},
			},
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
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 1},
					&Integer{Value: 2},
				},
			},
			sep:      ",",
			expected: "1,2",
		},
		{
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 2},
					&Integer{Value: 3},
					&Integer{Value: 5},
					&Integer{Value: 8},
					&Integer{Value: 13},
				},
			},
			sep:      ",",
			expected: "2,3,5,8,13",
		},
		{
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 1},
					&Integer{Value: 2},
				},
			},
			sep:      "/",
			expected: "1/2",
		},
		{
			obj: &Array{
				Elements: []Object{
					&Integer{Value: 2},
					&Integer{Value: 3},
					&Integer{Value: 5},
					&Integer{Value: 8},
					&Integer{Value: 13},
				},
			},
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
