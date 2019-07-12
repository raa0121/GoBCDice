package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

type Int struct {
	NodeImpl
	PrimaryExpressionImpl
	Value int
}

// IntがNodeを実装していることの確認
var _ Node = (*Int)(nil)

// IntがPrimaryExpressionを実装していることの確認
var _ PrimaryExpression = (*Int)(nil)

func (n *Int) Type() NodeType {
	return INT_NODE
}

func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}

func (n *Int) InfixNotation() string {
	return fmt.Sprintf("%d", n.Value)
}

func NewInt(value int, tok token.Token) *Int {
	return &Int{
		NodeImpl: NodeImpl{
			tok: tok,
		},
		Value: value,
	}
}
