package ast

// VariableNode は常に可変であるノードの型。
type VariableNode struct{}

// IsVariable は可変ノードかどうかを返す。
// VariableNodeでは常にtrueを返す。
func (n *VariableNode) IsVariable() bool {
	return true
}

// ConstNode は常に変化しないノードの型。
type ConstNode struct{}

// IsVariable は可変ノードかどうかを返す。
// ConstNodeでは常にfalseを返す。
func (n *ConstNode) IsVariable() bool {
	return false
}
