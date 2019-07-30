package object

import (
	"fmt"
	"testing"
)

func TestInteger_Type(t *testing.T) {
	obj := &Integer{}

	expected := INTEGER_OBJ
	actual := obj.Type()
	if actual != expected {
		t.Fatalf("got=%s, want=%s", actual, expected)
	}
}

func TestInteger_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *Integer
		expected string
	}{
		{&Integer{Value: 1}, "1"},
		{&Integer{Value: 0}, "0"},
		{&Integer{Value: -1}, "-1"},
		{&Integer{Value: 12}, "12"},
		{&Integer{Value: 12345}, "12345"},
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
