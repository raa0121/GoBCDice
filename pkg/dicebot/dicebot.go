/*
ダイスボットのインターフェースを定義するパッケージ。
*/
package dicebot

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/command"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
)

// ダイスボットを構築する関数の型。
type DiceBotConstructor func() DiceBot

// ダイスボットのインターフェース。
type DiceBot interface {
	// GameID はゲーム識別子を返す。
	GameID() string
	// GameName はゲームシステム名を返す。
	GameName() string
	// Usage はダイスボットの使用法の説明を返す。
	Usage() string
	// ExecuteCommand は指定されたコマンドを実行する。
	ExecuteCommand(command string, ev *evaluator.Evaluator) (*command.Result, error)
}

// DiceBotBasicInfo はダイスボットの基本情報を表す構造体。
type DiceBotBasicInfo struct {
	// GameID はゲーム識別子。
	GameID string
	// GameName はゲームシステム名。
	GameName string
	// Usage はダイスボットの使用法の説明。
	Usage string
}

// DiceBotImpl はダイスボットの実装のベースとなる構造体。
type DiceBotImpl struct {
	// BasicInfo はダイスボットの基本情報。
	BasicInfo *DiceBotBasicInfo
}

// DiceBotImpl がDiceBotを実装していることを確認する。
var _ DiceBot = (*DiceBotImpl)(nil)

// GameID はゲーム識別子を返す。
func (d *DiceBotImpl) GameID() string {
	return d.BasicInfo.GameID
}

// GameName はゲームシステム名を返す。
func (d *DiceBotImpl) GameName() string {
	return d.BasicInfo.GameName
}

// Usage はダイスボットの使用法の説明を返す。
func (d *DiceBotImpl) Usage() string {
	return d.BasicInfo.Usage
}

// ExecuteCommand は指定されたコマンドを実行する。
//
// 基本のダイスボットには特別なコマンドが存在しないため、必ずエラーを返す。
func (b *DiceBotImpl) ExecuteCommand(
	_ string,
	_ *evaluator.Evaluator,
) (*command.Result, error) {
	return nil, fmt.Errorf("no game-system-specific command")
}
