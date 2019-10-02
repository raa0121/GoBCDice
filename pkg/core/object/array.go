package object

import (
	"bytes"
	"strings"
)

// 配列オブジェクトの構造体
type Array struct {
	// 要素
	Elements []Object
}

// NewArray は新しい配列オブジェクトを返す。
func NewArray(elements ...Object) *Array {
	a := &Array{
		Elements: make([]Object, len(elements)),
	}

	if len(elements) > 0 {
		copy(a.Elements, elements)
	}

	return a
}

// NewArray は、スライスelementsを参照する新しい配列オブジェクトを返す。
// 要素のコピーが必要ない場合に使用する。
func NewArrayByMove(elements []Object) *Array {
	return &Array{
		Elements: elements,
	}
}

// Type はオブジェクトの種類を返す。
func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (a *Array) Inspect() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(a.JoinedElements(", "))
	out.WriteString("]")

	return out.String()
}

// InspectWithoutSpaces は、Inspectの区切り文字に空白が含まれないもの。
func (a *Array) InspectWithoutSpaces() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(a.JoinedElementsWithoutSpaces(","))
	out.WriteString("]")

	return out.String()
}

// JoinedElements は要素の内容を区切り文字sepを使って結合した文字列を返す。
func (a *Array) JoinedElements(sep string) string {
	elements := make([]string, 0, len(a.Elements))
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	return strings.Join(elements, sep)
}

// JoinedElementsWithoutSpaces は要素の内容を区切り文字sepを使って結合した文字列を返す。
// 配列を含む場合、内部の配列の区切り文字にも空白を含めない。
func (a *Array) JoinedElementsWithoutSpaces(sep string) string {
	elements := make([]string, 0, len(a.Elements))
	for _, e := range a.Elements {
		ea, eIsArray := e.(*Array)
		if eIsArray {
			elements = append(elements, ea.InspectWithoutSpaces())
		} else {
			elements = append(elements, e.Inspect())
		}
	}

	return strings.Join(elements, sep)
}
