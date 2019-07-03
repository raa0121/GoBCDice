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

func (n *DRoll) Token() token.Token {
	return n.Tok
}

func (n *DRoll) Type() string {
	return "DRoll"
}

func (n *DRoll) SExp() string {
	return fmt.Sprintf("(DRoll %s %s)", n.Num.SExp(), n.Sides.SExp())
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
