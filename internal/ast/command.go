package ast

// トップレベルにあるコマンドのインターフェース
type Command interface {
	Node
	// Expressionは、コマンドの引数である式のノードを返す
	Expression() Node
	// SetExpressionは、式のノードをeに設定する
	SetExpression(Node)
}

type CommandImpl struct {
	NodeImpl

	// コマンドの引数である式のノード
	expr Node
}

// Expressionは、コマンドの引数である式のノードを返す
func (n *CommandImpl) Expression() Node {
	return n.expr
}

// SetExpressionは、式のノードをeに設定する
func (n *CommandImpl) SetExpression(e Node) {
	n.expr = e
}

// IsVariableは可変ノードかどうかを返す。
func (n *CommandImpl) IsVariable() bool {
	return n.Expression().IsVariable()
}
