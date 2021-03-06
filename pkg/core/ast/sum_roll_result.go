package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"strings"
)

// 加算ロールの結果を表すノード。
// 一次式。
type SumRollResult struct {
	NodeImpl
	NonNilNode
	ConstNode

	// 振られたダイスの配列
	Dice []dice.Die
}

// SumRollResult がNodeを実装していることの確認。
var _ Node = (*SumRollResult)(nil)

// NewSumRollResult は新しい加算ロール結果のノードを返す。
//
// rolledDice: 振られたダイスのスライス。
func NewSumRollResult(rolledDice []dice.Die) *SumRollResult {
	r := &SumRollResult{
		NodeImpl: NodeImpl{
			nodeType:            SUM_ROLL_RESULT_NODE,
			isPrimaryExpression: true,
		},

		Dice: make([]dice.Die, len(rolledDice)),
	}

	copy(r.Dice, rolledDice)

	return r
}

// Value は出目の合計を返す。
func (n *SumRollResult) Value() int {
	sum := 0

	for _, d := range n.Dice {
		sum += d.Value
	}

	return sum
}

// SExp はノードのS式を返す。
func (n *SumRollResult) SExp() string {
	diceStrs := []string{}

	for _, d := range n.Dice {
		diceStrs = append(diceStrs, d.SExp())
	}

	return "(SumRollResult " + strings.Join(diceStrs, " ") + ")"
}
