package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input        string
		expectedSExp string
	}{
		{"C(1)", "(Calc 1)"},
		{"C(42)", "(Calc 42)"},
		{"C(1+2)", "(Calc (+ 1 2))"},
		{"C(1-2)", "(Calc (- 1 2))"},
		{"C(1*2)", "(Calc (* 1 2))"},
		{"C(1/2)", "(Calc (/ 1 2))"},
		{"C(1+2-3)", "(Calc (- (+ 1 2) 3))"},
		{"C(1*2+3)", "(Calc (+ (* 1 2) 3))"},
		{"C(1/2+3)", "(Calc (+ (/ 1 2) 3))"},
		{"C(1+2*3)", "(Calc (+ 1 (* 2 3)))"},
		{"C(1+2/3)", "(Calc (+ 1 (/ 2 3)))"},
		{"2D6", "(DRoll 2 6)"},
		{"12D60", "(DRoll 12 60)"},
	}

	for i, test := range testCases {
		actual, err := Parse(test.input)

		if err != nil {
			t.Errorf("#%d (%q): got err: %s", i, test.input, err)
			continue
		}

		actualSExp := actual.SExp()

		if actualSExp != test.expectedSExp {
			t.Errorf("#%d (%q): wrong value: got: %q, want: %q",
				i, test.input, actualSExp, test.expectedSExp)
		}
	}
}
