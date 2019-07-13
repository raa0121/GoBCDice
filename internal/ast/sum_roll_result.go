package ast

import (
	"github.com/raa0121/GoBCDice/internal/die"
	"strings"
)

// 加算ロールの結果を表すノード
type SumRollResult struct {
	NodeImpl

	// 振られたダイスの配列
	Dice []die.Die
}

// SumRollResultがNodeを実装していることの確認
var _ Node = (*SumRollResult)(nil)

// Typeはノードの種類を返す
func (n *SumRollResult) Type() NodeType {
	return SUM_ROLL_RESULT_NODE
}

// Valueは出目の合計を返す
func (n *SumRollResult) Value() int {
	sum := 0

	for _, d := range n.Dice {
		sum += d.Value
	}

	return sum
}

func (n *SumRollResult) SExp() string {
	diceStrs := []string{}

	for _, d := range n.Dice {
		diceStrs = append(diceStrs, d.SExp())
	}

	return "(SumRollResult " + strings.Join(diceStrs, " ") + ")"
}

// IsPrimaryExpressionは一次式かどうかを返す
func (n *SumRollResult) IsPrimaryExpression() bool {
	return true
}

// IsVariableは可変ノードかどうかを返す。
func (n *SumRollResult) IsVariable() bool {
	return false
}

// NewSumRollResultは、新しい加算ロールの結果のノードを返す
func NewSumRollResult(dice []die.Die) *SumRollResult {
	r := &SumRollResult{
		Dice: make([]die.Die, len(dice)),
	}

	copy(r.Dice, dice)

	return r
}
