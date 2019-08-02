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
				t.Fatalf("got=%s, want=%s", actual, test.expected)
			}
		})
	}
}
