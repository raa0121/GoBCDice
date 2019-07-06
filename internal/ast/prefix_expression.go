package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

type PrefixExpression struct {
	Tok             token.Token
	Operator        string
	OperatorForSExp string
	Right           Node
}

func (n *PrefixExpression) Token() token.Token {
	return n.Tok
}

func (n *PrefixExpression) Type() NodeType {
	return PREFIX_EXPRESSION_NODE
}

func (n *PrefixExpression) SExp() string {
	return fmt.Sprintf("(%s %s)", n.OperatorForSExp, n.Right.SExp())
}

func NewPrefixExpression(tok token.Token, right Node) *PrefixExpression {
	return &PrefixExpression{
		Tok:             tok,
		Operator:        tok.Literal,
		OperatorForSExp: tok.Literal,
		Right:           right,
	}
}
