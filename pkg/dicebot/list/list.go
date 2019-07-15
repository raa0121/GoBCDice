package list

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/dicebot"
	"github.com/raa0121/GoBCDice/pkg/dicebot/gamesystem/basic"
	"sort"
)

// Find は指定された識別子を持つゲームシステムのダイスボットのコンストラクタを返す。
// ゲームシステムが見つからなかった場合はエラーを返す。
func Find(gameID string) (dicebot.DiceBotConstructor, error) {
	if gameID == basic.GAME_ID {
		return basic.New, nil
	}

	constructor, found := gameIDToDiceBotConstructor[gameID]
	if !found {
		return nil, fmt.Errorf("unknown game system: %s", gameID)
	}

	return constructor, nil
}

// AvailableGameIDs は利用可能なゲームシステムの識別子のスライスを返す。
func AvailableGameIDs(includeBasicDiceBot bool) []string {
	gameIDs := make([]string, 0, len(gameIDToDiceBotConstructor))
	for k := range gameIDToDiceBotConstructor {
		gameIDs = append(gameIDs, k)
	}

	if includeBasicDiceBot {
		gameIDs = append(gameIDs, basic.GAME_ID)
	}

	gameIDsStrSlice := sort.StringSlice(gameIDs)
	sort.Sort(gameIDsStrSlice)

	return gameIDs
}

// ゲーム識別子とダイスボットのコンストラクタとの対応
var gameIDToDiceBotConstructor = map[string]dicebot.DiceBotConstructor{}
