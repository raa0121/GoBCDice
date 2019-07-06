package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

type Int struct {
	Tok   token.Token
	Value int
}

func (n *Int) Token() token.Token {
	return n.Tok
}

func (n *Int) Type() NodeType {
	return INT_NODE
}

func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}
