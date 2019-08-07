package peg_parser

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input        string
		expectedSExp string
		err          bool
	}{
		// 計算コマンド
		{"C(1)", "(Calc 1)", false},
		{"C(42)", "(Calc 42)", false},
		{"C(-1)", "(Calc (- 1))", false},
		{"C(+1)", "(Calc 1)", false},
		{"C(1+2)", "(Calc (+ 1 2))", false},
		{"C(1-2)", "(Calc (- 1 2))", false},
		{"C(1*2)", "(Calc (* 1 2))", false},

		// int_expr SLASH int_expr
		{"C(1/2)", "(Calc (/ 1 2))", false},
		// int_expr SLASH int_expr U
		{"C(1/2u)", "(Calc (/U 1 2))", false},
		// int_expr SLASH int_expr R
		{"C(1/2r)", "(Calc (/R 1 2))", false},

		{"C(-1+2)", "(Calc (+ (- 1) 2))", false},
		{"C(+1+2)", "(Calc (+ 1 2))", false},
		{"C(1+2-3)", "(Calc (- (+ 1 2) 3))", false},
		{"C(1*2+3)", "(Calc (+ (* 1 2) 3))", false},
		{"C(1/2+3)", "(Calc (+ (/ 1 2) 3))", false},
		{"C(1+2*3)", "(Calc (+ 1 (* 2 3)))", false},
		{"C(1+2/3)", "(Calc (+ 1 (/ 2 3)))", false},
		{"C(1+(2-3))", "(Calc (+ 1 (- 2 3)))", false},
		{"C((1+2)*3)", "(Calc (* (+ 1 2) 3))", false},
		{"C((1+2)/3)", "(Calc (/ (+ 1 2) 3))", false},
		{"C((1+2)/3+4*5-6)", "(Calc (- (+ (/ (+ 1 2) 3) (* 4 5)) 6))", false},
		{"C((1+2)/3u+4*5-6)", "(Calc (- (+ (/U (+ 1 2) 3) (* 4 5)) 6))", false},
		{"C((1+2)/3r+4*5-6)", "(Calc (- (+ (/R (+ 1 2) 3) (* 4 5)) 6))", false},
		{"C(100/(1+2))", "(Calc (/ 100 (+ 1 2)))", false},
		{"C(100/(1+2)u)", "(Calc (/U 100 (+ 1 2)))", false},
		{"C(100/(1+2)r)", "(Calc (/R 100 (+ 1 2)))", false},
		{"C(-(1+2))", "(Calc (- (+ 1 2)))", false},
		{"C(+(1+2))", "(Calc (+ 1 2))", false},
		{"CC(1)", "", true},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, err := Parse("test", []byte(test.input))

			if err != nil {
				// エラーが発生した！

				if !test.err {
					// エラーを想定していなかった場合
					t.Fatalf("got err: %s", err)
				}

				// エラーを想定していたので、この挙動は正常

				return
			}

			if test.err {
				// エラーを想定していたのに、発生していない
				t.Fatal("should err")
				return
			}

			node, ok := r.(ast.Node)

			if !ok {
				t.Fatal("not returned a node")
				return
			}

			if node == nil {
				t.Fatal("got nil node")
				return
			}

			actualSExp := node.SExp()

			if actualSExp != test.expectedSExp {
				t.Errorf("wrong SExp: got: %q, want: %q",
					actualSExp, test.expectedSExp)
			}
		})
	}
}
