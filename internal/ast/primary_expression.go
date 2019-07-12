package ast

// 一次式のインターフェース
type PrimaryExpression interface {
	// IsPrimaryExpressionは一次式かどうかを返す（ダミー関数）
	IsPrimaryExpression() bool
}

// 一次式のノードを表す構造体
type PrimaryExpressionImpl struct {
}

// PrimaryExpressionImplがPrimaryExpressionを実装していることの確認
var _ PrimaryExpression = (*PrimaryExpressionImpl)(nil)

// IsPrimaryExpressionは一次式かどうかを返す（ダミー関数）
func (n *PrimaryExpressionImpl) IsPrimaryExpression() bool {
	return true
}
