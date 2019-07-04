package die

import (
	"fmt"
)

// サイコロを表す構造体
type Die struct {
	// 出目
	Value int
	// サイコロの面の数
	Sides int
}

func (d Die) String() string {
	return fmt.Sprintf("<Die %d/%d>", d.Value, d.Sides)
}
