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

// MaxValue は出目のグループの最大値を返す。
func (r *URollExprResult) MaxValue() *Integer {
	max, _ := r.SumOfGroups().MaxInteger()

	return NewInteger(max.Value + r.Modifier.Value)
}

// SumOfValues は出目の合計を返す。
func (r *URollExprResult) SumOfValues() *Integer {
	sumOfGroups := r.SumOfGroups()
	sum, _ := r.SumOfGroups().SumOfIntegers()

	return NewInteger(sum.Value + len(sumOfGroups.Elements)*r.Modifier.Value)
}

// SumOfGroups は出目のグループごとの合計値の配列を返す。
func (r *URollExprResult) SumOfGroups() *Array {
	sums := make([]Object, 0, len(r.ValueGroups.Elements))

	for _, e := range r.ValueGroups.Elements {
		ea := e.(*Array)
		sum, ok := ea.SumOfIntegers()
		if !ok {
			panic("ValueGroups contain non-Integer elements")
		}

		sums = append(sums, sum)
	}

	return NewArrayByMove(sums)
}
