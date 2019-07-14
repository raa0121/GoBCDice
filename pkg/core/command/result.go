package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"strings"
)

// コマンドの実行結果の構造体
type Result struct {
	// ゲーム識別子
	GameId string
	// メッセージの部分の配列
	MessageParts []string
	// 振られたダイス
	RolledDice []dice.Die
}

// JoinedMessageParts は、メッセージの部分を結合したものを返す。
func (r *Result) JoinedMessageParts() string {
	return strings.Join(r.MessageParts, " ＞ ")
}

// Message はコマンドの応答メッセージを返す。
func (r *Result) Message() string {
	return r.GameId + " : " + r.JoinedMessageParts()
}

// appendMessagePart はメッセージの部分を追加する。
func (r *Result) appendMessagePart(message string) {
	r.MessageParts = append(r.MessageParts, message)
}
