package dicebot

import (
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
