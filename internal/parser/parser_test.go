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
		{"C(1+(2-3))", "(Calc (+ 1 (- 2 3)))", false},
		{"C((1+2)*3)", "(Calc (* (+ 1 2) 3))", false},
		{"C((1+2)/3)", "(Calc (/ (+ 1 2) 3))", false},
		{"CC(1)", "", true},
		{"C([1...3])", "(Calc (Rand 1 3))", false},
		{"C([1...3]*2)", "(Calc (* (Rand 1 3) 2))", false},
		{"C([(1+2)...(4+5)])", "(Calc (Rand (+ 1 2) (+ 4 5)))", false},
		{"C([1+2...4-5])", "", true},
		{"C([1...2...3])", "", true},
		{"2D6", "(DRollExpr (DRoll 2 6))", false},
		{"2D6D6", "", true},
		{"12D60", "(DRollExpr (DRoll 12 60))", false},
		{"1", "", true},
		{"2D6+1", "(DRollExpr (+ (DRoll 2 6) 1))", false},
		{"1+2D6", "(DRollExpr (+ 1 (DRoll 2 6)))", false},
		{"2d6+1-1-2-3-4", "(DRollExpr (- (- (- (- (+ (DRoll 2 6) 1) 1) 2) 3) 4))", false},
		{"2D6+4D10", "(DRollExpr (+ (DRoll 2 6) (DRoll 4 10)))", false},
		{"(2D6)", "(DRollExpr (DRoll 2 6))", false},
		{"(1)", "", true},
		{"2d6*3", "(DRollExpr (* (DRoll 2 6) 3))", false},
		{"2d10+3-4", "(DRollExpr (- (+ (DRoll 2 10) 3) 4))", false},
		{"2d10+3*4", "(DRollExpr (+ (DRoll 2 10) (* 3 4)))", false},
		{"2d6*3-1d6+1", "(DRollExpr (+ (- (* (DRoll 2 6) 3) (DRoll 1 6)) 1))", false},
		{"(2+3)d6-1+3d6+2", "(DRollExpr (+ (+ (- (DRoll (+ 2 3) 6) 1) (DRoll 3 6)) 2))", false},
		{"(2*3-4)d6-1d4+1", "(DRollExpr (+ (- (DRoll (- (* 2 3) 4) 6) (DRoll 1 4)) 1))", false},
		{"((2+3)*4/3)d6*2+5", "(DRollExpr (+ (* (DRoll (/ (* (+ 2 3) 4) 3) 6) 2) 5))", false},
		{"2d(1+5)", "(DRollExpr (DRoll 2 (+ 1 5)))", false},
		{"(8/2)D(4+6)", "(DRollExpr (DRoll (/ 8 2) (+ 4 6)))", false},
		{"(2-1)d(8/2)*(1+1)d(3*4/2)+2*3", "(DRollExpr (+ (* (DRoll (- 2 1) (/ 8 2)) (DRoll (+ 1 1) (/ (* 3 4) 2))) (* 2 3)))", false},
		{"[1...5]D6", "(DRollExpr (DRoll (Rand 1 5) 6))", false},
		{"([2...4]+2)D10", "(DRollExpr (DRoll (+ (Rand 2 4) 2) 10))", false},
		{"[(2+3)...8]D6", "(DRollExpr (DRoll (Rand (+ 2 3) 8) 6))", false},
		{"[5...(7+1)]D6", "(DRollExpr (DRoll (Rand 5 (+ 7 1)) 6))", false},
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
			t.Errorf("#%d (%q): wrong SExp: got: %q, want: %q",
				i, test.input, actualSExp, test.expectedSExp)
		}
	}
}
