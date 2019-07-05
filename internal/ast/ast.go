package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

type NodeType int

func (t NodeType) String() string {
	if str, ok := nodeTypeString[t]; ok {
		return str
	}

	return nodeTypeString[UNKNOWN_NODE]
}

const (
	UNKNOWN_NODE NodeType = iota

	COMMAND_NODE
	PREFIX_EXPRESSION_NODE
	INFIX_EXPRESSION_NODE
	INT_NODE
)

var nodeTypeString = map[NodeType]string{
	UNKNOWN_NODE: "UNKNOWN",

	COMMAND_NODE:           "Command",
	PREFIX_EXPRESSION_NODE: "PrefixExpression",
	INFIX_EXPRESSION_NODE:  "InfixExpression",
	INT_NODE:               "Int",
}

type Node interface {
	Token() token.Token
	Type() NodeType
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

func (n *Command) Type() NodeType {
	return COMMAND_NODE
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

func (n *InfixExpression) Type() NodeType {
	return INFIX_EXPRESSION_NODE
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

func (n *Int) Type() NodeType {
	return INT_NODE
}

func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}
