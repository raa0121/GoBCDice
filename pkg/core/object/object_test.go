package object

import (
	"testing"
)

func TestObject_Type(t *testing.T) {
	testcases := []struct {
		obj      Object
		expected string
	}{
		{&Integer{}, "INTEGER"},
		{&Boolean{}, "BOOLEAN"},
		{&String{}, "STRING"},
		{&Array{}, "ARRAY"},
		{&BRollCompResult{}, "B_ROLL_COMP_RESULT"},
		{&RRollCompResult{}, "R_ROLL_COMP_RESULT"},
		{&URollExprResult{}, "U_ROLL_EXPR_RESULT"},
		{&URollCompResult{}, "U_ROLL_COMP_RESULT"},
	}

	for _, test := range testcases {
		t.Run(test.expected, func(t *testing.T) {
			actual := test.obj.Type().String()
			if actual != test.expected {
				t.Errorf("got: %q, want: %q", actual, test.expected)
			}
		})
	}
}
