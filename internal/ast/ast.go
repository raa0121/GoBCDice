package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

type Node interface {
	Token() token.Token
	Type() string
	SExp() string
}

type Calc struct {
	// 最初のトークン
	Tok        token.Token
	Expression Node
}

func (n *Calc) Token() token.Token {
	return n.Tok
}

func (n *Calc) Type() string {
	return "Calc"
}

func (n *Calc) SExp() string {
	return fmt.Sprintf("(Calc %s)", n.Expression.SExp())
}

type DRoll struct {
	Tok   token.Token
	Num   Node
	Sides Node
}

type InfixExpression struct {
	Tok             token.Token
	Left            Node
	Operator        string
	OperatorForSExp string
	Right           Node
}

func (n *InfixExpression) Token() token.Token {
	return n.Tok
}

func (n *InfixExpression) Type() string {
	return "InfixExpression"
}

func (n *InfixExpression) SExp() string {
	return fmt.Sprintf("(%s %s %s)", n.OperatorForSExp, n.Left.SExp(), n.Right.SExp())
}

func NewDRoll(num Node, tok token.Token, sides Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            num,
		Operator:        "D",
		OperatorForSExp: "DRoll",
		Right:           sides,
	}
}

func NewInfixExpression(left Node, tok token.Token, right Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            left,
		Operator:        tok.Literal,
		OperatorForSExp: tok.Literal,
		Right:           right,
	}
}

type Int struct {
	Tok   token.Token
	Value int
}

func (n *Int) Token() token.Token {
	return n.Tok
}

func (n *Int) Type() string {
	return "Int"
}

func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}
