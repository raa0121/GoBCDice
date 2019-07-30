package object

import (
	"fmt"
	"testing"
)

func TestBoolean_Type(t *testing.T) {
	obj := &Boolean{}

	expected := BOOLEAN_OBJ
	actual := obj.Type()
	if actual != expected {
		t.Fatalf("got=%s, want=%s", actual, expected)
	}
}

func TestBoolean_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *Boolean
		expected string
	}{
		{&Boolean{Value: true}, "true"},
		{&Boolean{Value: false}, "false"},
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
