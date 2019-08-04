package object

import (
	"fmt"
	"testing"
)

func TestBoolean_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *Boolean
		expected string
	}{
		{NewBoolean(true), "true"},
		{NewBoolean(false), "false"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%t", test.obj.Value), func(t *testing.T) {
			actual := test.obj.Inspect()
			if actual != test.expected {
				t.Fatalf("got=%s, want=%s", actual, test.expected)
			}
		})
	}
}
