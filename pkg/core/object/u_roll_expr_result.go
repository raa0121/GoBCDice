package object

import (
	"bytes"
)

// URollExprResult は上方無限ロール式の結果を表すオブジェクト。
//
// 計算回数を少なくするため、読み出し専用としている。
type URollExprResult struct {
	// 出目のグループの配列
	valueGroups *Array
	// 修正値
	modifier *Integer

	// 出目のグループごとの合計値の配列
	sumOfGroups *Array
}

// URollExprResultがObjectを実装していることを確認する
var _ Object = (*URollExprResult)(nil)

// NewURollExprResult は、上方無限ロール式の評価結果を表す新しいオブジェクトを返す。
func NewURollExprResult(valueGroups *Array, modifier *Integer) *URollExprResult {
	return &URollExprResult{
		valueGroups: valueGroups,
		modifier:    modifier,
	}
}

// Type はオブジェクトの種類を返す。
func (r *URollExprResult) Type() ObjectType {
	return U_ROLL_EXPR_RESULT_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (r *URollExprResult) Inspect() string {
	var out bytes.Buffer

	out.WriteString("<URollExprResult ValueGroups=")
	out.WriteString(r.valueGroups.Inspect())
	out.WriteString(", Modifier=")
	out.WriteString(r.modifier.Inspect())
	out.WriteString(">")

	return out.String()
}

// ValueGroups は出目のグループの配列を返す。
func (r *URollExprResult) ValueGroups() *Array {
	return r.valueGroups
}

// Modifier は修正値を返す。
func (r *URollExprResult) Modifier() *Integer {
	return r.modifier
}

// SumOfGroups は出目のグループごとの合計値の配列を返す。
func (r *URollExprResult) SumOfGroups() *Array {
	if r.sumOfGroups == nil {
		// 呼び出し初回に配列を求め、キャッシュする
		r.calculateSumOfGroups()
	}

	return r.sumOfGroups
}

// MaxValue は出目のグループの最大値を返す。
func (r *URollExprResult) MaxValue() *Integer {
	max, _ := r.SumOfGroups().MaxInteger()

	return max.Add(r.modifier)
}

// SumOfValues は出目の合計を返す。
func (r *URollExprResult) SumOfValues() *Integer {
	sum, _ := r.SumOfGroups().SumOfIntegers()

	return sum.Add(r.modifier)
}

// calculateSumOfGroups は出目のグループごとの合計値を計算する。
func (r *URollExprResult) calculateSumOfGroups() {
	valueGroups := r.valueGroups
	sums := make([]Object, 0, valueGroups.Length())

	for _, e := range valueGroups.Elements {
		ea := e.(*Array)
		sum, ok := ea.SumOfIntegers()
		if !ok {
			panic("valueGroups contain non-Integer elements")
		}

		sums = append(sums, sum)
	}

	r.sumOfGroups = NewArrayByMove(sums)
}
