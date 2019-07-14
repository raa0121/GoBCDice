package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"github.com/raa0121/GoBCDice/pkg/core/die/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/die/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"reflect"
	"testing"
)

// 抽象構文木を評価し、値のオブジェクトに変換する例。
func ExampleEvaluator_Eval() {
	// 構文解析する
	ast, parseErr := parser.Parse("(2*3-4)d6-1d4+1")
	if parseErr != nil {
		return
	}

	fmt.Println("抽象構文木: " + ast.SExp())

	// ノードを評価する
	dieFeeder := feeder.NewQueue([]die.Die{{6, 6}, {2, 6}, {3, 6}})
	evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

	obj, evalErr := evaluator.Eval(ast)
	if evalErr != nil {
		return
	}

	fmt.Println("ダイスロール結果: " + die.FormatDice(evaluator.RolledDice()))
	fmt.Println("評価結果: " + obj.Inspect())
	// Output:
	// 抽象構文木: (DRollExpr (+ (- (DRoll (- (* 2 3) 4) 6) (DRoll 1 4)) 1))
	// ダイスロール結果: 6/6, 2/6, 3/6
	// 評価結果: 6
}

func TestEvalCalc(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{"C(5)", 5},
		{"C(10)", 10},
		{"C(42)", 42},
		{"C(65535)", 65535},
		{"C(-5)", -5},
		{"C(-10)", -10},
		{"C(-42)", -42},
		{"C(-65535)", -65535},
		{"C(+5)", 5},
		{"C(+10)", 10},
		{"C(+42)", 42},
		{"C(+65535)", 65535},
		{"C(1+2)", 3},
		{"C(1-2)", -1},
		{"C(1*2)", 2},
		{"C(1/2)", 0},
		{"C(-1+2)", 1},
		{"C(+1+2)", 3},
		{"C(1+2-3)", 0},
		{"C(1*2+3)", 5},
		{"C(1/2+3)", 3},
		{"C(1+2*3)", 7},
		{"C(1+2/3)", 1},
		{"C(1+(2-3))", 0},
		{"C((1+2)*3)", 9},
		{"C((1+2)/3)", 1},
		{"C(-(1+2))", -3},
		{"C(+(1+2))", 3},
		{"C(1+50/3-2)", 15},
		{"C(1+50/3u-2)", 16},
		{"C(1+50/3r-2)", 16},
		{"C(1+100/3-2)", 32},
		{"C(1+100/3u-2)", 33},
		{"C(1+100/3r-2)", 32},
	}

	for _, test := range testcases {
		t.Run(test.input, func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Errorf("構文解析エラー: %s", parseErr)
				return
			}

			// ノードを評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evaluated, evalErr := evaluator.Eval(ast)
			if evalErr != nil {
				t.Errorf("評価エラー: %s", evalErr)
				return
			}

			if evaluated == nil {
				t.Error("Evalの対象外 (nil)")
				return
			}

			// 型が合っているか？
			obj, typeMatched := evaluated.(*object.Integer)
			if !typeMatched {
				t.Errorf("整数オブジェクトでない: %T (%+v)", obj, obj)
				return
			}

			if obj.Value != test.expected {
				t.Errorf("異なる値: got=%d, want=%d", obj.Value, test.expected)
			}
		})
	}
}

func TestEvalDRollExpr(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
		dice     []die.Die
	}{
		{
			input:    "2D6",
			expected: 8,
			dice:     []die.Die{{5, 6}, {3, 6}},
		},
		{
			input:    "2D4",
			expected: 3,
			dice:     []die.Die{{1, 4}, {2, 4}},
		},
		{
			input:    "2D6+1",
			expected: 9,
			dice:     []die.Die{{2, 6}, {6, 6}},
		},
		{
			input:    "1+2D6",
			expected: 8,
			dice:     []die.Die{{4, 6}, {3, 6}},
		},
		{
			input:    "2d6+1-1-2-3-4",
			expected: -2,
			dice:     []die.Die{{1, 6}, {6, 6}},
		},
		{
			input:    "2D6+4D10",
			expected: 30,
			dice:     []die.Die{{5, 6}, {4, 6}, {1, 10}, {9, 10}, {7, 10}, {4, 10}},
		},
		{
			input:    "2d6*3",
			expected: 18,
			dice:     []die.Die{{2, 6}, {4, 6}},
		},
		{
			input:    "2d10+3-4",
			expected: 7,
			dice:     []die.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d10+3*4",
			expected: 20,
			dice:     []die.Die{{3, 10}, {5, 10}},
		},
		{
			input:    "2d6*3-1d6+1",
			expected: 22,
			dice:     []die.Die{{6, 6}, {2, 6}, {3, 6}},
		},
		{
			input:    "(2+3)d6-1+3d6+2",
			expected: 31,
			dice:     []die.Die{{2, 6}, {3, 6}, {1, 6}, {5, 6}, {6, 6}, {5, 6}, {4, 6}, {4, 6}},
		},
		{
			input:    "(2*3-4)d6-1d4+1",
			expected: 10,
			dice:     []die.Die{{6, 6}, {5, 6}, {2, 4}},
		},
		{
			input:    "((2+3)*4/3)d6*2+5",
			expected: 53,
			dice:     []die.Die{{6, 6}, {5, 6}, {6, 6}, {2, 6}, {1, 6}, {4, 6}},
		},
		{
			input:    "1D6/2",
			expected: 0,
			dice:     []die.Die{{1, 6}},
		},
		{
			input:    "3D6/2",
			expected: 3,
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "3D6/2+1D6",
			expected: 8,
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2",
			expected: 9,
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "3D6+1D6/2U",
			expected: 10,
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}, {5, 6}},
		},
		{
			input:    "5D6/10",
			expected: 2,
			dice:     []die.Die{{6, 6}, {6, 6}, {6, 6}, {6, 6}, {5, 6}},
		},
		{
			input:    "3D6/2U",
			expected: 4,
			dice:     []die.Die{{1, 6}, {2, 6}, {4, 6}},
		},
		{
			input:    "5D6/10u",
			expected: 3,
			dice:     []die.Die{{6, 6}, {6, 6}, {6, 6}, {2, 6}, {1, 6}},
		},
		{
			input:    "1D100/10R",
			expected: 6,
			dice:     []die.Die{{55, 100}},
		},
		{
			input:    "1D100/10r",
			expected: 5,
			dice:     []die.Die{{54, 100}},
		},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			// ノードを評価する
			dieFeeder := feeder.NewQueue(test.dice)
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evaluated, evalErr := evaluator.Eval(ast)
			if evalErr != nil {
				t.Fatalf("評価エラー: %s", evalErr)
				return
			}

			if evaluated == nil {
				t.Fatalf("Evalの対象外 (nil)")
				return
			}

			// 型が合っているか？
			obj, typeMatched := evaluated.(*object.Integer)
			if !typeMatched {
				t.Fatalf("整数オブジェクトでない: %T (%+v)", obj, obj)
				return
			}

			if obj.Value != test.expected {
				t.Errorf("異なる評価結果: got=%d, want=%d", obj.Value, test.expected)
			}

			rolledDice := evaluator.RolledDice()
			if !reflect.DeepEqual(rolledDice, test.dice) {
				t.Errorf("異なるダイスロール結果記録: got=%v, want=%v",
					rolledDice, test.dice)
			}
		})
	}
}
