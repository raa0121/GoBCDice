package ast

// Nil はnilを表すノード。
// 一次式。
type Nil struct{}

// nilInstance はnilの唯一のインスタンス。
var nilInstance = &Nil{}

// Nil がNodeを実装していることの確認。
var _ Node = (*Nil)(nil)

// Type はノードの種類を返す。
func (n *Nil) Type() NodeType {
	return NIL_NODE
}

// IsNil はnilかどうかを返す。
func (n *Nil) IsNil() bool {
	return true
}

// SExp はノードのS式を返す。
func (n *Nil) SExp() string {
	return "nil"
}

// IsPrimaryExpression は一次式かどうかを返す。
// Nilではtrueを返す。
func (n *Nil) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// Nilではfalseを返す。
func (n *Nil) IsVariable() bool {
	return false
}

// NilInstance はnilのインスタンスを返す。
func NilInstance() *Nil {
	return nilInstance
}
