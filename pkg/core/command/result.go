package command

import (
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"strings"
)

// コマンドの実行結果の構造体
type Result struct {
	// ゲーム識別子
	GameId string
	// メッセージの部分の配列
	MessageParts []string
	// 振られたダイス
	RolledDice []die.Die
}

// JoinedMessagePartsは、メッセージの部分を結合したものを返す
func (r *Result) JoinedMessageParts() string {
	return strings.Join(r.MessageParts, " ＞ ")
}

// Messageはコマンドの応答メッセージを返す
func (r *Result) Message() string {
	return r.GameId + " : " + r.JoinedMessageParts()
}

// appendMessagePartはメッセージの部分を追加する
func (r *Result) appendMessagePart(message string) {
	r.MessageParts = append(r.MessageParts, message)
}
