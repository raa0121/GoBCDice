package object

import (
	"bytes"
)

// URollExprResult は上方無限ロール式の結果を表すオブジェクト。
type URollExprResult struct {
	// 出目のグループの配列
	ValueGroups *Array
	// 修正値
	Modifier *Integer
}

// URollExprResultがObjectを実装していることを確認する
var _ Object = (*URollExprResult)(nil)

// Type はオブジェクトの種類を返す。
func (r *URollExprResult) Type() ObjectType {
	return U_ROLL_EXPR_RESULT_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (r *URollExprResult) Inspect() string {
	var out bytes.Buffer

	out.WriteString("<URollExprResult ValueGroups=")
	out.WriteString(r.ValueGroups.Inspect())
	out.WriteString(", Modifier=")
	out.WriteString(r.Modifier.Inspect())
	out.WriteString(">")

	return out.String()
}