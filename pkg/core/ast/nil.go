package ast

// NonNilNode はnilでないノードの型。
type NonNilNode struct{}

// IsNil はnilかどうかを返す。
func (n *NonNilNode) IsNil() bool {
	return false
}

// Nil はnilを表すノード。
// 一次式。
type Nil struct {
	NodeImpl
	ConstNode
}

// Nil がNodeを実装していることの確認。
var _ Node = (*Nil)(nil)

// nilInstance はnilの唯一のインスタンス。
var nilInstance = &Nil{
	NodeImpl: NodeImpl{
		nodeType:            NIL_NODE,
		isPrimaryExpression: true,
	},
}

// NilInstance はnilのインスタンスを返す。
func NilInstance() *Nil {
	return nilInstance
}

// IsNil はnilかどうかを返す。
func (n *Nil) IsNil() bool {
	return true
}

// SExp はノードのS式を返す。
func (n *Nil) SExp() string {
	return "nil"
}
