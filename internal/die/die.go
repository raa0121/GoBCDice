package die

import (
	"fmt"
	"strings"
)

// ダイスを表す構造体
type Die struct {
	// 出目
	Value int
	// ダイスの面の数
	Sides int
}

// Stringはダイスの文字列表現を返す
func (d Die) String() string {
	return fmt.Sprintf("<Die %d/%d>", d.Value, d.Sides)
}

// FormatDiceはダイス列を文字列として整形して返す。
// 結果の文字列は "値/面数, 値/面数, ..." という形式。
func FormatDice(dice []Die) string {
	dieStrs := []string{}
	for _, d := range dice {
		dieStr := fmt.Sprintf("%d/%d", d.Value, d.Sides)
		dieStrs = append(dieStrs, dieStr)
	}

	return strings.Join(dieStrs, ", ")
}

// FormatDiceWithoutSpacesはダイス列を文字列として整形して返す。
// 結果の文字列は "値/面数,値/面数,..." という形式。
// 空白を出力しないので、テストケースなどで使うとよい。
func FormatDiceWithoutSpaces(dice []Die) string {
	dieStrs := []string{}
	for _, d := range dice {
		dieStr := fmt.Sprintf("%d/%d", d.Value, d.Sides)
		dieStrs = append(dieStrs, dieStr)
	}

	return strings.Join(dieStrs, ",")
}
