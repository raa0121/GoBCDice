package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"strings"
)

// 成功判定の結果の型
type SuccessCheckResultType int

const (
	// 成功判定結果：未指定（既定値）
	SUCCESS_CHECK_UNSPECIFIED SuccessCheckResultType = iota
	// 成功判定結果：成功
	SUCCESS_CHECK_SUCCESS
	// 成功判定結果：失敗
	SUCCESS_CHECK_FAILURE
)

// 成功判定結果の文字列表現
var successCheckResultString = map[SuccessCheckResultType]string{
	SUCCESS_CHECK_UNSPECIFIED: "UNSPECIFIED",
	SUCCESS_CHECK_SUCCESS:     "SUCCESS",
	SUCCESS_CHECK_FAILURE:     "FAILURE",
}

// String は成功判定結果を文字列として返す。
func (r SuccessCheckResultType) String() string {
	if s, found := successCheckResultString[r]; found {
		return s
	}

	return "UNKNOWN"
}

// コマンドの実行結果の構造体
type Result struct {
	// ゲーム識別子
	GameID string
	// メッセージの部分の配列
	MessageParts []string
	// 振られたダイス
	RolledDice []dice.Die
	// 成功判定の結果
	SuccessCheckResult SuccessCheckResultType
}

// JoinedMessageParts は、メッセージの部分を結合したものを返す。
func (r *Result) JoinedMessageParts() string {
	return strings.Join(r.MessageParts, " ＞ ")
}

// Message はコマンドの応答メッセージを返す。
func (r *Result) Message() string {
	return r.GameID + " : " + r.JoinedMessageParts()
}

// appendMessagePart はメッセージの部分を追加する。
func (r *Result) appendMessagePart(message string) {
	r.MessageParts = append(r.MessageParts, message)
}
