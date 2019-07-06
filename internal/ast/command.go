package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

// トップレベルにあるコマンドのノード
type Command struct {
	Tok        token.Token
	Expression Node
	Name       string
}

func (n *Command) Token() token.Token {
	return n.Tok
}

func (n *Command) Type() NodeType {
	return COMMAND_NODE
}

func (n *Command) SExp() string {
	return fmt.Sprintf("(%s %s)", n.Name, n.Expression.SExp())
}
