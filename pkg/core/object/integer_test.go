package object

import (
	"fmt"
	"testing"
)

func TestInteger_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *Integer
		expected string
	}{
		{NewInteger(1), "1"},
		{NewInteger(0), "0"},
		{NewInteger(-1), "-1"},
		{NewInteger(12), "12"},
		{NewInteger(12345), "12345"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%d", test.obj.Value), func(t *testing.T) {
			actual := test.obj.Inspect()
			if actual != test.expected {
				t.Fatalf("got=%q, want=%q", actual, test.expected)
			}
		})
	}
}

func TestInteger_Add(t *testing.T) {
	testcases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{3, 5, 8},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%d+%d", test.a, test.b), func(t *testing.T) {
			aObj := NewInteger(test.a)
			bObj := NewInteger(test.b)

			actual := aObj.Add(bObj).Value
			if actual != test.expected {
				t.Fatalf("got=%d, want=%d", actual, test.expected)
			}
		})
	}
}
