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

		// 計算コマンド内でのランダム数値は無効にする
		{"C([1...3])", "", true},
		{"C([1...3]*2)", "", true},
		{"C([(1+2)...(4+5)])", "", true},

		// ランダム数値取り出しの構文エラー
		{"C([1+2...4-5])", "", true},
		{"C([1...2...3])", "", true},

		// 加算ロール
		{"2D6", "(DRollExpr (DRoll 2 6))", false},
		{"2D6D6", "", true},
		{"12D60", "(DRollExpr (DRoll 12 60))", false},
		{"1", "", true},
		{"-2D6", "(DRollExpr (- (DRoll 2 6)))", false},
		{"+2D6", "(DRollExpr (DRoll 2 6))", false},
		{"2D6+1", "(DRollExpr (+ (DRoll 2 6) 1))", false},
		{"1+2D6", "(DRollExpr (+ 1 (DRoll 2 6)))", false},
		{"1+2D6+2", "(DRollExpr (+ (+ 1 (DRoll 2 6)) 2))", false},
		{"-2D6+1", "(DRollExpr (+ (- (DRoll 2 6)) 1))", false},
		{"+2D6+1", "(DRollExpr (+ (DRoll 2 6) 1))", false},
		{"2d6+1-1-2-3-4", "(DRollExpr (- (- (- (- (+ (DRoll 2 6) 1) 1) 2) 3) 4))", false},
		{"2D6+4D10", "(DRollExpr (+ (DRoll 2 6) (DRoll 4 10)))", false},
		{"(2D6)", "(DRollExpr (DRoll 2 6))", false},
		{"-(2D6)", "(DRollExpr (- (DRoll 2 6)))", false},
		{"+(2D6)", "(DRollExpr (DRoll 2 6))", false},
		{"(1)", "", true},
		{"2d6*3", "(DRollExpr (* (DRoll 2 6) 3))", false},

		// d_roll_expr SLASH int_expr
		{"2d6/2", "(DRollExpr (/ (DRoll 2 6) 2))", false},
		// d_roll_expr SLASH int_expr U
		{"2d6/2u", "(DRollExpr (/U (DRoll 2 6) 2))", false},
		// d_roll_expr SLASH int_expr R
		{"2d6/2r", "(DRollExpr (/R (DRoll 2 6) 2))", false},

		// int_expr SLASH d_roll_expr
		{"100/2d6+1", "(DRollExpr (+ (/ 100 (DRoll 2 6)) 1))", false},
		// int_expr SLASH d_roll_expr U
		{"100/2d6u+1", "(DRollExpr (+ (/U 100 (DRoll 2 6)) 1))", false},
		// int_expr SLASH d_roll_expr R
		{"100/2d6r+1", "(DRollExpr (+ (/R 100 (DRoll 2 6)) 1))", false},

		// int_expr SLASH d_roll_expr
		{"100/(2d6+1)+4*5", "(DRollExpr (+ (/ 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},
		// int_expr SLASH d_roll_expr U
		{"100/(2d6+1)u+4*5", "(DRollExpr (+ (/U 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},
		// int_expr SLASH d_roll_expr R
		{"100/(2d6+1)r+4*5", "(DRollExpr (+ (/R 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},

		// d_roll_expr SLASH d_roll_expr
		{"4d10/2d6+1", "(DRollExpr (+ (/ (DRoll 4 10) (DRoll 2 6)) 1))", false},
		// d_roll_expr SLASH d_roll_expr U
		{"4d10/2d6u+1", "(DRollExpr (+ (/U (DRoll 4 10) (DRoll 2 6)) 1))", false},
		// d_roll_expr SLASH d_roll_expr R
		{"4d10/2d6r+1", "(DRollExpr (+ (/R (DRoll 4 10) (DRoll 2 6)) 1))", false},

		{"2d10+3-4", "(DRollExpr (- (+ (DRoll 2 10) 3) 4))", false},
		{"2d10+3*4", "(DRollExpr (+ (DRoll 2 10) (* 3 4)))", false},
		{"2d10/3+4*5-6", "(DRollExpr (- (+ (/ (DRoll 2 10) 3) (* 4 5)) 6))", false},
		{"2d10/3u+4*5-6", "(DRollExpr (- (+ (/U (DRoll 2 10) 3) (* 4 5)) 6))", false},
		{"2d10/3r+4*5-6", "(DRollExpr (- (+ (/R (DRoll 2 10) 3) (* 4 5)) 6))", false},
		{"2d6*3-1d6+1", "(DRollExpr (+ (- (* (DRoll 2 6) 3) (DRoll 1 6)) 1))", false},
		{"(2+3)d6-1+3d6+2", "(DRollExpr (+ (+ (- (DRoll (+ 2 3) 6) 1) (DRoll 3 6)) 2))", false},
		{"(2*3-4)d6-1d4+1", "(DRollExpr (+ (- (DRoll (- (* 2 3) 4) 6) (DRoll 1 4)) 1))", false},
		{"((2+3)*4/3)d6*2+5", "(DRollExpr (+ (* (DRoll (/ (* (+ 2 3) 4) 3) 6) 2) 5))", false},
		{"2d(1+5)", "(DRollExpr (DRoll 2 (+ 1 5)))", false},
		{"(8/2)D(4+6)", "(DRollExpr (DRoll (/ 8 2) (+ 4 6)))", false},
		{"(2-1)d(8/2)*(1+1)d(3*4/2)+2*3", "(DRollExpr (+ (* (DRoll (- 2 1) (/ 8 2)) (DRoll (+ 1 1) (/ (* 3 4) 2))) (* 2 3)))", false},

		// ランダム数値取り出しを含む加算ロール
		{"[1...5]D6", "(DRollExpr (DRoll (Rand 1 5) 6))", false},
		{"([2...4]+2)D10", "(DRollExpr (DRoll (+ (Rand 2 4) 2) 10))", false},
		{"[(2+3)...8]D6", "(DRollExpr (DRoll (Rand (+ 2 3) 8) 6))", false},
		{"[5...(7+1)]D6", "(DRollExpr (DRoll (Rand 5 (+ 7 1)) 6))", false},
		{"2d[1...5]", "(DRollExpr (DRoll 2 (Rand 1 5)))", false},
		{"2d([2...4]+2)", "(DRollExpr (DRoll 2 (+ (Rand 2 4) 2)))", false},
		{"2d[(2+3)...8]", "(DRollExpr (DRoll 2 (Rand (+ 2 3) 8)))", false},
		{"2d[5...(7+1)]", "(DRollExpr (DRoll 2 (Rand 5 (+ 7 1))))", false},
		{"[1...5]d(2*3)", "(DRollExpr (DRoll (Rand 1 5) (* 2 3)))", false},
		{"(1+1)d[1...5]", "(DRollExpr (DRoll (+ 1 1) (Rand 1 5)))", false},
		{"([1...4]+1)d([2...4]+2)-1", "(DRollExpr (- (DRoll (+ (Rand 1 4) 1) (+ (Rand 2 4) 2)) 1))", false},

		// 加算ロール式の成功判定
		{"2d6=7", "(DRollComp (= (DRoll 2 6) 7))", false},
		{"2d6<>7", "(DRollComp (<> (DRoll 2 6) 7))", false},
		{"2d6>7", "(DRollComp (> (DRoll 2 6) 7))", false},
		{"2d6<7", "(DRollComp (< (DRoll 2 6) 7))", false},
		{"2d6>=7", "(DRollComp (>= (DRoll 2 6) 7))", false},
		{"2d6<=7", "(DRollComp (<= (DRoll 2 6) 7))", false},
		{"2d6>=5+3", "(DRollComp (>= (DRoll 2 6) (+ 5 3)))", false},
		{"2d6+1>=3+4", "(DRollComp (>= (+ (DRoll 2 6) 1) (+ 3 4)))", false},
		{"1+2d6>=3+4", "(DRollComp (>= (+ 1 (DRoll 2 6)) (+ 3 4)))", false},
		{"2*(2d6+1)/3<7", "(DRollComp (< (/ (* 2 (+ (DRoll 2 6) 1)) 3) 7))", false},
		{"7<2d6", "", true},
		{"2d6<7<8", "", true},
		{"1<=2d6<=12", "", true},
		{"1<2<2d6", "", true},

		// バラバラロール
		{"2b6", "(BRollList (BRoll 2 6))", false},
		{"[1...3]b6", "(BRollList (BRoll (Rand 1 3) 6))", false},
		{"2b[4...6]", "(BRollList (BRoll 2 (Rand 4 6)))", false},
		{"[1...3]b[4...6]", "(BRollList (BRoll (Rand 1 3) (Rand 4 6)))", false},
		{"(1*2)b6", "(BRollList (BRoll (* 1 2) 6))", false},
		{"([1...3]+1)b6", "(BRollList (BRoll (+ (Rand 1 3) 1) 6))", false},
		{"2b(2+4)", "(BRollList (BRoll 2 (+ 2 4)))", false},
		{"2b([3...5]+1)", "(BRollList (BRoll 2 (+ (Rand 3 5) 1)))", false},
		{"[1...5]b(2*3)", "(BRollList (BRoll (Rand 1 5) (* 2 3)))", false},
		{"(1+1)b[1...5]", "(BRollList (BRoll (+ 1 1) (Rand 1 5)))", false},
		{"(1*2)b(2+4)", "(BRollList (BRoll (* 1 2) (+ 2 4)))", false},
		{"2b6+4b10", "(BRollList (BRoll 2 6) (BRoll 4 10))", false},
		{"2b6+3b8+5b12", "(BRollList (BRoll 2 6) (BRoll 3 8) (BRoll 5 12))", false},
		{"2b6+1", "", true},
		{"1+2b6", "", true},
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
				t.Fatalf("not returned a node: %s", r)
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
