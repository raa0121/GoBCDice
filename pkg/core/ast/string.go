package ast

import (
	"fmt"
)

// 文字列のノード。
type String struct {
	NonNilNode

	// 文字列
	Value string
}

// String がNodeを実装していることの確認。
var _Node = (*String)(nil)

// Type はノードの種類を返す。
func (n *String) Type() NodeType {
	return STRING_NODE
}

// SExp はノードのS式を返す。
func (n *String) SExp() string {
	return fmt.Sprintf("%q", n.Value)
}

// IsPrimaryExpression は一次式かどうかを返す。
// Stringではtrueを返す。
func (n *String) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// Stringではfalseを返す。
func (n *String) IsVariable() bool {
	return false
}

// NewString は新しい文字列のノードを返す。
//
// value: 文字列,
func NewString(value string) *String {
	return &String{
		Value: value,
	}
}
