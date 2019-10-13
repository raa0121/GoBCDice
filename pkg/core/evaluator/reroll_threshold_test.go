package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"testing"
)

func TestSetRerollThreshold(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"3r6<=4", "(RRollComp (<= (RRollList 4 (RRoll 3 6)) 4))"},
		{"3r6+2r4<=2", "(RRollComp (<= (RRollList 2 (RRoll 3 6) (RRoll 2 4)) 2))"},
		{"(3+2)r6>=5", "(RRollComp (>= (RRollList 5 (RRoll (+ 3 2) 6)) 5))"},
		{"1r(2*3)>=4", "(RRollComp (>= (RRollList 4 (RRoll 1 (* 2 3))) 4))"},
		{"3r6>1*4", "(RRollComp (> (RRollList 4 (RRoll 3 6)) (* 1 4)))"},
		{"6R6[6]>=5", "(RRollComp (>= (RRollList 6 (RRoll 6 6)) 5))"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(*ast.Command)
			if node.Type() != ast.R_ROLL_COMP_NODE {
				t.Fatal("RRollCompでない")
				return
			}

			// 可変ノードの引数を評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			evalErr := evaluator.SetRerollThreshold(node)
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

func TestCheckRRollThreshold(t *testing.T) {
	testcases := []struct {
		input      string
		err        bool
		errMessage string
	}{
		{"2r6[4]", false, ""},
		{"6R6[6]", false, ""},
		{"2r6", true, "2R6>=5 あるいは 2R6[5] のように振り足し目標値を指定してください"},
		{"2r6[0]", true, "振り足し目標値として2以上の整数を指定してください"},
		{"2r6+3r4[1]", true, "振り足し目標値として2以上の整数を指定してください"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(*ast.RRollList)

			// 可変ノードの引数を評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			checkErr := evaluator.CheckRRollThreshold(node)
			if checkErr != nil {
				if !test.err {
					t.Fatalf("閾値チェックエラー: %s", checkErr)
					return
				}

				if checkErr.Error() != test.errMessage {
					t.Fatalf("異なるエラーメッセージ: got=%q, want=%q",
						checkErr.Error(), test.errMessage)
					return
				}

				return
			}

			if test.err {
				t.Fatal("エラーが発生しない")
				return
			}
		})
	}
}

func TestCheckURollThreshold(t *testing.T) {
	testcases := []struct {
		input      string
		err        bool
		errMessage string
	}{
		{"2u6[4]", false, ""},
		{"6U6[6]", false, ""},
		{"2u6", true, "2U6[5] のように振り足し目標値を指定してください"},
		{"2u6[0]", true, "振り足し目標値として2以上の整数を指定してください"},
		{"2u6+3u4[1]", true, "振り足し目標値として2以上の整数を指定してください"},
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(*ast.URollExpr)

			// 可変ノードの引数を評価する
			dieFeeder := feeder.NewEmptyQueue()
			evaluator := NewEvaluator(roller.New(dieFeeder), NewEnvironment())

			checkErr := evaluator.CheckURollThreshold(node.URollList)
			if checkErr != nil {
				if !test.err {
					t.Fatalf("閾値チェックエラー: %s", checkErr)
					return
				}

				if checkErr.Error() != test.errMessage {
					t.Fatalf("異なるエラーメッセージ: got=%q, want=%q",
						checkErr.Error(), test.errMessage)
					return
				}

				return
			}

			if test.err {
				t.Fatal("エラーが発生しない")
				return
			}
		})
	}
}
