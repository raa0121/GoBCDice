package ast

// トップレベルにあるコマンドのインターフェース
type Command interface {
	Node
	Expression() Node
}

type CommandImpl struct {
	NodeImpl
	expr Node
}

func (n *CommandImpl) Expression() Node {
	return n.expr
}
