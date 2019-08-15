package ast

// トップレベルにあるコマンドのインターフェース。
type Command interface {
	Node

	// Expression はコマンドの引数である式のノードを返す。
	Expression() Node
	// SetExpression は式のノードを設定する。
	SetExpression(Node)
}

// Command が共通して持つ要素。
type CommandImpl struct {
	// コマンドの引数である式のノード
	expr Node
}

// Expression はコマンドの引数である式のノードを返す。
func (n *CommandImpl) Expression() Node {
	return n.expr
}

// SetExpression は式のノードを設定する。
func (n *CommandImpl) SetExpression(e Node) {
	n.expr = e
}

// IsPrimaryExpression は一次式かどうかを返す。
// コマンドではfalseを返す。
func (n *CommandImpl) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// コマンドでは、引数の式が可変ノードならばtrueを返す。
// 引数の式が可変ノードでない場合はfalseを返す。
func (n *CommandImpl) IsVariable() bool {
	return n.Expression().IsVariable()
}
