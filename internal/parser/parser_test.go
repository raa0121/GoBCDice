package parser

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	expected := IntNode{value: 42}

	actual, err := Parse("42")

	if err != nil {
		t.Fatalf("got err: %s", err)
	}

	actualInt := actual.(IntNode)

	if actualInt.value != expected.value {
		t.Errorf("wrong value: got: %v, want: %v", actualInt.value, expected.value)
	}
}
