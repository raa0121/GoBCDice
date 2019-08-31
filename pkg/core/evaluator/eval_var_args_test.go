package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"testing"
)

// 可変ノードの引数を評価して整数に変換する例。
func ExampleEvaluator_EvalVarArgs() {
	// 構文解析する
	r, parseErr := parser.Parse("ExampleEvaluator_EvalVarArgs", []byte("(2*3-4)d6-1d4+1"))
	if parseErr != nil {
		return
	}

	node := r.(ast.Node)

	fmt.Println("構文解析直後の抽象構文木: " + node.SExp())

	// 可変ノードの引数を評価して整数に変換する
	dieFeeder := feeder.NewMT19937WithSeedFromTime()
	evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

	evalErr := evaluator.EvalVarArgs(node)
	if evalErr != nil {
		return
	}

	fmt.Println("引数評価後の抽象構文木: " + node.SExp())
	// Output:
	// 構文解析直後の抽象構文木: (DRollExpr (+ (- (DRoll (- (* 2 3) 4) 6) (DRoll 1 4)) 1))
	// 引数評価後の抽象構文木: (DRollExpr (+ (- (DRoll 2 6) (DRoll 1 4)) 1))
}

func TestEvalVarArgs(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		dice     []dice.Die
	}{
		// 加算ロール
		{
			input:    "(1+2)d6",
			expected: "(DRollExpr (DRoll 3 6))",
		},
		{
			input:    "4d(3*2)",
			expected: "(DRollExpr (DRoll 4 6))",
		},
		{
			input:    "(8/2)D(4+6)",
			expected: "(DRollExpr (DRoll 4 10))",
		},
		{
			input:    "-(1+2)d6",
			expected: "(DRollExpr (- (DRoll 3 6)))",
		},
		{
			input:    "(2+3)d6-1+3d6+2",
			expected: "(DRollExpr (+ (+ (- (DRoll 5 6) 1) (DRoll 3 6)) 2))",
		},
		{
			input:    "(2*3-4)d6-1d4+1",
			expected: "(DRollExpr (+ (- (DRoll 2 6) (DRoll 1 4)) 1))",
		},
		{
			input:    "((2+3)*4/3)d6*2+5",
			expected: "(DRollExpr (+ (* (DRoll 6 6) 2) 5))",
		},
		{
			input:    "(2-1)d(8/2)*(1+1)d(3*4/2)+2*3",
			expected: "(DRollExpr (+ (* (DRoll 1 4) (DRoll 2 6)) (* 2 3)))",
		},

		// ランダム数値埋め込み
		{
			input:    "[1...5]D6",
			expected: "(DRollExpr (DRoll 4 6))",
			dice:     []dice.Die{{4, 5}},
		},
		{
			input:    "([2...4]+2)D10",
			expected: "(DRollExpr (DRoll 6 10))",
			dice:     []dice.Die{{3, 3}},
		},
		{
			input:    "[(2+3)...8]D6",
			expected: "(DRollExpr (DRoll 5 6))",
			dice:     []dice.Die{{1, 4}},
		},
		{
			input:    "[5...(7+1)]D6",
			expected: "(DRollExpr (DRoll 5 6))",
			dice:     []dice.Die{{1, 4}},
		},
		{
			input:    "2d[1...5]",
			expected: "(DRollExpr (DRoll 2 2))",
			dice:     []dice.Die{{2, 5}},
		},
		{
			input:    "2d([2...4]+2)",
			expected: "(DRollExpr (DRoll 2 5))",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "2d[(2+3)...8]",
			expected: "(DRollExpr (DRoll 2 8))",
			dice:     []dice.Die{{4, 4}},
		},
		{
			input:    "2d[5...(7+1)]",
			expected: "(DRollExpr (DRoll 2 8))",
			dice:     []dice.Die{{4, 4}},
		},
		{
			input:    "([1...4]+1)d([2...4]+2)-1",
			expected: "(DRollExpr (- (DRoll 3 6) 1))",
			dice:     []dice.Die{{2, 4}, {3, 3}},
		},

		// 加算ロール式の成功判定
		{
			input:    "(1*2)D6+1>7",
			expected: "(DRollComp (> (+ (DRoll 2 6) 1) 7))",
		},
		{
			input:    "2D(3*2)+1>7",
			expected: "(DRollComp (> (+ (DRoll 2 6) 1) 7))",
		},
		{
			input:    "(1*2)d6+3d(3+1)>15",
			expected: "(DRollComp (> (+ (DRoll 2 6) (DRoll 3 4)) 15))",
		},
		{
			input:    "2d6+1>3+4",
			expected: "(DRollComp (> (+ (DRoll 2 6) 1) 7))",
		},

		// バラバラロール
		{
			input:    "[1...3]b6",
			expected: "(BRollList (BRoll 1 6))",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "2b[4...6]",
			expected: "(BRollList (BRoll 2 4))",
			dice:     []dice.Die{{1, 3}},
		},
		{
			input:    "[1...3]b[4...6]",
			expected: "(BRollList (BRoll 3 6))",
			dice:     []dice.Die{{3, 3}, {3, 3}},
		},
		{
			input:    "(1*2)b6",
			expected: "(BRollList (BRoll 2 6))",
		},
		{
			input:    "([1...3]+1)b6",
			expected: "(BRollList (BRoll 3 6))",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "2b(2+4)",
			expected: "(BRollList (BRoll 2 6))",
		},
		{
			input:    "2b([3...5]+1)",
			expected: "(BRollList (BRoll 2 5))",
			dice:     []dice.Die{{2, 3}},
		},
		{
			input:    "(1*2)b(2+4)",
			expected: "(BRollList (BRoll 2 6))",
		},
		{
			input:    "(1*2)b(2+4)+(8/2)b(101/10R)",
			expected: "(BRollList (BRoll 2 6) (BRoll 4 10))",
		},

		// バラバラロールの成功数カウント
		{
			input:    "2b6+4b10>10/3",
			expected: "(BRollComp (> (BRollList (BRoll 2 6) (BRoll 4 10)) 3))",
		},

		// 個数振り足しロール
		{
			input:    "(3+2)r6",
			expected: "(RRollList nil (RRoll 5 6))",
		},
		{
			input:    "1r(2*3)",
			expected: "(RRollList nil (RRoll 1 6))",
		},
		{
			input:    "(1*3)r6+2r(5+1)",
			expected: "(RRollList nil (RRoll 3 6) (RRoll 2 6))",
		},
		{
			input:    "6R6[2*3]",
			expected: "(RRollList 6 (RRoll 6 6))",
		},

		// 個数振り足しロールの成功数カウント
		{
			input:    "3r6>1*4",
			expected: "(RRollComp (> (RRollList nil (RRoll 3 6)) 4))",
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(ast.Node)

			// 可変ノードの引数を評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evalErr := evaluator.EvalVarArgs(node)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			actual := node.SExp()
			if actual != test.expected {
				t.Errorf("異なる評価結果: got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
