package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

type Node interface {
	Token() token.Token
	SExp() string
}

type Command struct {
	Tok        token.Token
	Expression Node
	Name       string
}

func (n *Command) Token() token.Token {
	return n.Tok
}

func (n *Command) Type() string {
	return "Command"
}

func (n *Command) SExp() string {
	return fmt.Sprintf("(%s %s)", n.Name, n.Expression.SExp())
}

type PrefixExpression struct {
	Tok             token.Token
	Operator        string
	OperatorForSExp string
	Right           Node
}

func (n *PrefixExpression) Token() token.Token {
	return n.Tok
}

func (n *PrefixExpression) Type() string {
	return "PrefixExpression"
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

func NewRand(min Node, tok token.Token, max Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            min,
		Operator:        "...",
		OperatorForSExp: "Rand",
		Right:           max,
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
