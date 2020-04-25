package parser

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"testing"
)

// 構文解析の例。
func Example() {
	// 構文解析する
	r, err := Parse("Example", []byte("(2*3-4)d6-1d4+1"))
	if err != nil {
		return
	}

	// 得られた抽象構文木のS式を出力する
	if node, ok := r.(ast.Node); ok {
		fmt.Println(node.SExp())
	}
	// Output:
	// (DRollExpr (+ (- (DRoll (- (* 2 3) 4) 6) (DRoll 1 4)) 1))
}

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
		// int_expr SLASH int_expr D
		{"C(1/2d)", "(Calc (/ 1 2))", false},
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
		{"C((1+2)/3d+4*5-6)", "(Calc (- (+ (/ (+ 1 2) 3) (* 4 5)) 6))", false},
		{"C((1+2)/3u+4*5-6)", "(Calc (- (+ (/U (+ 1 2) 3) (* 4 5)) 6))", false},
		{"C((1+2)/3r+4*5-6)", "(Calc (- (+ (/R (+ 1 2) 3) (* 4 5)) 6))", false},
		{"C(100/(1+2))", "(Calc (/ 100 (+ 1 2)))", false},
		{"C(100/(1+2)d)", "(Calc (/ 100 (+ 1 2)))", false},
		{"C(100/(1+2)u)", "(Calc (/U 100 (+ 1 2)))", false},
		{"C(100/(1+2)r)", "(Calc (/R 100 (+ 1 2)))", false},
		{"C(-(1+2))", "(Calc (- (+ 1 2)))", false},
		{"C(+(1+2))", "(Calc (+ 1 2))", false},
		{"CC(1)", "", true},
		{"C(10+5) mokekeke", "(Calc (+ 10 5))", false},

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
		// d_roll_expr SLASH int_expr D
		{"2d6/2d", "(DRollExpr (/ (DRoll 2 6) 2))", false},
		// d_roll_expr SLASH int_expr U
		{"2d6/2u", "(DRollExpr (/U (DRoll 2 6) 2))", false},
		// d_roll_expr SLASH int_expr R
		{"2d6/2r", "(DRollExpr (/R (DRoll 2 6) 2))", false},

		// int_expr SLASH d_roll_expr
		{"100/2d6+1", "(DRollExpr (+ (/ 100 (DRoll 2 6)) 1))", false},
		// int_expr SLASH d_roll_expr D
		{"100/2d6d+1", "(DRollExpr (+ (/ 100 (DRoll 2 6)) 1))", false},
		// int_expr SLASH d_roll_expr U
		{"100/2d6u+1", "(DRollExpr (+ (/U 100 (DRoll 2 6)) 1))", false},
		// int_expr SLASH d_roll_expr R
		{"100/2d6r+1", "(DRollExpr (+ (/R 100 (DRoll 2 6)) 1))", false},

		// int_expr SLASH d_roll_expr
		{"100/(2d6+1)+4*5", "(DRollExpr (+ (/ 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},
		// int_expr SLASH d_roll_expr D
		{"100/(2d6+1)d+4*5", "(DRollExpr (+ (/ 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},
		// int_expr SLASH d_roll_expr U
		{"100/(2d6+1)u+4*5", "(DRollExpr (+ (/U 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},
		// int_expr SLASH d_roll_expr R
		{"100/(2d6+1)r+4*5", "(DRollExpr (+ (/R 100 (+ (DRoll 2 6) 1)) (* 4 5)))", false},

		// d_roll_expr SLASH d_roll_expr
		{"4d10/2d6+1", "(DRollExpr (+ (/ (DRoll 4 10) (DRoll 2 6)) 1))", false},
		// d_roll_expr SLASH d_roll_expr D
		{"4d10/2d6d+1", "(DRollExpr (+ (/ (DRoll 4 10) (DRoll 2 6)) 1))", false},
		// d_roll_expr SLASH d_roll_expr U
		{"4d10/2d6u+1", "(DRollExpr (+ (/U (DRoll 4 10) (DRoll 2 6)) 1))", false},
		// d_roll_expr SLASH d_roll_expr R
		{"4d10/2d6r+1", "(DRollExpr (+ (/R (DRoll 4 10) (DRoll 2 6)) 1))", false},

		{"2d10+3-4", "(DRollExpr (- (+ (DRoll 2 10) 3) 4))", false},
		{"2d10+3*4", "(DRollExpr (+ (DRoll 2 10) (* 3 4)))", false},
		{"2d10/3+4*5-6", "(DRollExpr (- (+ (/ (DRoll 2 10) 3) (* 4 5)) 6))", false},
		{"2d10/3d+4*5-6", "(DRollExpr (- (+ (/ (DRoll 2 10) 3) (* 4 5)) 6))", false},
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
		{"[1...5]D6", "(DRollExpr (DRoll (RandomNumber 1 5) 6))", false},
		{"([2...4]+2)D10", "(DRollExpr (DRoll (+ (RandomNumber 2 4) 2) 10))", false},
		{"[(2+3)...8]D6", "(DRollExpr (DRoll (RandomNumber (+ 2 3) 8) 6))", false},
		{"[5...(7+1)]D6", "(DRollExpr (DRoll (RandomNumber 5 (+ 7 1)) 6))", false},
		{"2d[1...5]", "(DRollExpr (DRoll 2 (RandomNumber 1 5)))", false},
		{"2d([2...4]+2)", "(DRollExpr (DRoll 2 (+ (RandomNumber 2 4) 2)))", false},
		{"2d[(2+3)...8]", "(DRollExpr (DRoll 2 (RandomNumber (+ 2 3) 8)))", false},
		{"2d[5...(7+1)]", "(DRollExpr (DRoll 2 (RandomNumber 5 (+ 7 1))))", false},
		{"[1...5]d(2*3)", "(DRollExpr (DRoll (RandomNumber 1 5) (* 2 3)))", false},
		{"(1+1)d[1...5]", "(DRollExpr (DRoll (+ 1 1) (RandomNumber 1 5)))", false},
		{"([1...4]+1)d([2...4]+2)-1", "(DRollExpr (- (DRoll (+ (RandomNumber 1 4) 1) (+ (RandomNumber 2 4) 2)) 1))", false},

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
		{"[1...3]b6", "(BRollList (BRoll (RandomNumber 1 3) 6))", false},
		{"2b[4...6]", "(BRollList (BRoll 2 (RandomNumber 4 6)))", false},
		{"[1...3]b[4...6]", "(BRollList (BRoll (RandomNumber 1 3) (RandomNumber 4 6)))", false},
		{"(1*2)b6", "(BRollList (BRoll (* 1 2) 6))", false},
		{"([1...3]+1)b6", "(BRollList (BRoll (+ (RandomNumber 1 3) 1) 6))", false},
		{"2b(2+4)", "(BRollList (BRoll 2 (+ 2 4)))", false},
		{"2b([3...5]+1)", "(BRollList (BRoll 2 (+ (RandomNumber 3 5) 1)))", false},
		{"[1...5]b(2*3)", "(BRollList (BRoll (RandomNumber 1 5) (* 2 3)))", false},
		{"(1+1)b[1...5]", "(BRollList (BRoll (+ 1 1) (RandomNumber 1 5)))", false},
		{"(1*2)b(2+4)", "(BRollList (BRoll (* 1 2) (+ 2 4)))", false},
		{"2b6+4b10", "(BRollList (BRoll 2 6) (BRoll 4 10))", false},
		{"2b6+3b8+5b12", "(BRollList (BRoll 2 6) (BRoll 3 8) (BRoll 5 12))", false},
		{"2b6+1", "", true},
		{"1+2b6", "", true},

		// バラバラロールの成功数カウント
		{"2b6=3", "(BRollComp (= (BRollList (BRoll 2 6)) 3))", false},
		{"2b6<>3", "(BRollComp (<> (BRollList (BRoll 2 6)) 3))", false},
		{"2b6>3", "(BRollComp (> (BRollList (BRoll 2 6)) 3))", false},
		{"2b6<3", "(BRollComp (< (BRollList (BRoll 2 6)) 3))", false},
		{"2b6>=3", "(BRollComp (>= (BRollList (BRoll 2 6)) 3))", false},
		{"2b6<=3", "(BRollComp (<= (BRollList (BRoll 2 6)) 3))", false},
		{"2b6>4-1", "(BRollComp (> (BRollList (BRoll 2 6)) (- 4 1)))", false},
		{"2b6+4b10>4", "(BRollComp (> (BRollList (BRoll 2 6) (BRoll 4 10)) 4))", false},
		{"2b6>-(-1*3)", "(BRollComp (> (BRollList (BRoll 2 6)) (- (* (- 1) 3))))", false},
		{"2b6+1>3", "", true},
		{"1+2b6>3", "", true},
		{"3<2b6", "", true},
		{"1<2b6<5", "", true},
		{"2b6<4<5", "", true},
		{"1<2<2b6", "", true},

		// 個数振り足しロール
		{"3r6=4", "(RRollComp (= (RRollList nil (RRoll 3 6)) 4))", false},
		{"3r6<>4", "(RRollComp (<> (RRollList nil (RRoll 3 6)) 4))", false},
		{"3r6>4", "(RRollComp (> (RRollList nil (RRoll 3 6)) 4))", false},
		{"3r6<4", "(RRollComp (< (RRollList nil (RRoll 3 6)) 4))", false},
		{"3r6>=4", "(RRollComp (>= (RRollList nil (RRoll 3 6)) 4))", false},
		{"3r6<=4", "(RRollComp (<= (RRollList nil (RRoll 3 6)) 4))", false},
		{"3r6+2r6<=2", "(RRollComp (<= (RRollList nil (RRoll 3 6) (RRoll 2 6)) 2))", false},
		{"(3+2)r6>=5", "(RRollComp (>= (RRollList nil (RRoll (+ 3 2) 6)) 5))", false},
		{"1r(2*3)>=4", "(RRollComp (>= (RRollList nil (RRoll 1 (* 2 3))) 4))", false},
		{"3r6>1*4", "(RRollComp (> (RRollList nil (RRoll 3 6)) (* 1 4)))", false},
		{"2r6", "(RRollList nil (RRoll 2 6))", false},
		{"2r6[5]", "(RRollList 5 (RRoll 2 6))", false},
		{"3r6+2r6[2]", "(RRollList 2 (RRoll 3 6) (RRoll 2 6))", false},
		{"6R6[6]>=5", "(RRollComp (>= (RRollList 6 (RRoll 6 6)) 5))", false},
		{"6R6[2*3]>=5", "(RRollComp (>= (RRollList (* 2 3) (RRoll 6 6)) 5))", false},
		{"2r6+1>=4", "", true},
		{"1<3r6<4", "", true},

		// 上方無限ロール
		{"3u6", "(URollExpr (RRollList nil (URoll 3 6)))", false},
		{"(1*3)u6", "(URollExpr (RRollList nil (URoll (* 1 3) 6)))", false},
		{"3u(5+1)", "(URollExpr (RRollList nil (URoll 3 (+ 5 1))))", false},
		{"3u6[6]", "(URollExpr (RRollList 6 (URoll 3 6)))", false},
		{"3u6[2+4]", "(URollExpr (RRollList (+ 2 4) (URoll 3 6)))", false},
		{"3u6+5u6[6]", "(URollExpr (RRollList 6 (URoll 3 6) (URoll 5 6)))", false},
		{"3u6[6]+1", "(URollExpr (+ (RRollList 6 (URoll 3 6)) 1))", false},
		{"3u6[6]-1", "(URollExpr (- (RRollList 6 (URoll 3 6)) 1))", false},
		{"1U100[96]+3", "(URollExpr (+ (RRollList 96 (URoll 1 100)) 3))", false},
		{"3u6[6]=10", "(URollComp (= (URollExpr (RRollList 6 (URoll 3 6))) 10))", false},
		{"3u6[6]<>10", "(URollComp (<> (URollExpr (RRollList 6 (URoll 3 6))) 10))", false},
		{"3u6[6]>10", "(URollComp (> (URollExpr (RRollList 6 (URoll 3 6))) 10))", false},
		{"3u6[6]<10", "(URollComp (< (URollExpr (RRollList 6 (URoll 3 6))) 10))", false},
		{"3u6[6]>=10", "(URollComp (>= (URollExpr (RRollList 6 (URoll 3 6))) 10))", false},
		{"3u6[6]<=10", "(URollComp (<= (URollExpr (RRollList 6 (URoll 3 6))) 10))", false},
		{"3u6[6]>=2+8", "(URollComp (>= (URollExpr (RRollList 6 (URoll 3 6))) (+ 2 8)))", false},
		{"3u6[6]+1>=10", "(URollComp (>= (URollExpr (+ (RRollList 6 (URoll 3 6)) 1)) 10))", false},
		{"3u6+5u6[6]>=7", "(URollComp (>= (URollExpr (RRollList 6 (URoll 3 6) (URoll 5 6))) 7))", false},
		{"(5+6)u10[10]+5>=8", "(URollComp (>= (URollExpr (+ (RRollList 10 (URoll (+ 5 6) 10)) 5)) 8))", false},
		{"5<3u6[6]<10", "", true},

		// ランダム選択
		{"choice[A,B,C]どれにしよう", `(Choice "A" "B" "C")`, false},
		{"choice[A,B, ]", `(Choice "A" "B")`, false},
		{"Choice[ A, B,   C     ,D ]", `(Choice "A" "B" "C" "D")`, false},
		{
			input:        "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expectedSExp: `(Choice "Call of Cthulhu" "Sword World" "Double Cross")`,
			err:          false,
		},
		{
			input:        "CHOICE[日本語, でも,　だいじょうぶ]",
			expectedSExp: `(Choice "日本語" "でも" "だいじょうぶ")`,
			err:          false,
		},
		{"choice[1+2, (3*4), 5d6]", `(Choice "1+2" "(3*4)" "5d6")`, false},
		{"choice[forgetting R_BRACKET!", "", true},
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
