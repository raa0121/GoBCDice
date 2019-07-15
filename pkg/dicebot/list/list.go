package list

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/dicebot"
	"github.com/raa0121/GoBCDice/pkg/dicebot/gamesystem/basic"
)

// Find は指定された識別子を持つゲームシステムのダイスボットのコンストラクタを返す。
// ゲームシステムが見つからなかった場合はエラーを返す。
func Find(gameID string) (dicebot.DiceBotConstructor, error) {
	constructor, found := gameIDToDiceBotConstructor[gameID]
	if !found {
		return nil, fmt.Errorf("unknown game system: %s", gameID)
	}

	return constructor, nil
}

// ゲーム識別子とダイスボットのコンストラクタとの対応
var gameIDToDiceBotConstructor = map[string]dicebot.DiceBotConstructor{
	"DiceBot": basic.New,
}
