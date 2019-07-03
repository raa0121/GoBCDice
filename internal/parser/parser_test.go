package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input        string
		expectedSExp string
		err          bool
	}{
		{"C(1)", "(Calc 1)", false},
		{"C(42)", "(Calc 42)", false},
		{"C(1+2)", "(Calc (+ 1 2))", false},
		{"C(1-2)", "(Calc (- 1 2))", false},
		{"C(1*2)", "(Calc (* 1 2))", false},
		{"C(1/2)", "(Calc (/ 1 2))", false},
		{"C(1+2-3)", "(Calc (- (+ 1 2) 3))", false},
		{"C(1*2+3)", "(Calc (+ (* 1 2) 3))", false},
		{"C(1/2+3)", "(Calc (+ (/ 1 2) 3))", false},
		{"C(1+2*3)", "(Calc (+ 1 (* 2 3)))", false},
		{"C(1+2/3)", "(Calc (+ 1 (/ 2 3)))", false},
		{"CC(1)", "", true},
		{"2D6", "(DRoll 2 6)", false},
		{"2D6D6", "", true},
		{"12D60", "(DRoll 12 60)", false},
	}

	for i, test := range testCases {
		actual, err := Parse(test.input)

		if err != nil {
			// エラーが発生した！

			if !test.err {
				// エラーを想定していなかった場合
				t.Errorf("#%d (%q): got err: %s", i, test.input, err)
			}

			// エラーを想定していたので、この挙動は正常

			continue
		}

		if test.err {
			// エラーを想定していたのに、発生していない
			t.Errorf("#%d (%q): should err", i, test.input)

			continue
		}

		actualSExp := actual.SExp()

		if actualSExp != test.expectedSExp {
			t.Errorf("#%d (%q): wrong value: got: %q, want: %q",
				i, test.input, actualSExp, test.expectedSExp)
		}
	}
}
