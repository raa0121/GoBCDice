/*
ダイスボットの一覧を管理するパッケージ。

このパッケージを使用することで、指定したゲーム名のダイスボットを取得することができるようになる。
*/
package list

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/dicebot"
	"github.com/raa0121/GoBCDice/pkg/dicebot/gamesystem/basic"
	"sort"
)

type Names struct {
	Name []SystemInfo
}

type SystemInfo struct {
	System  string `json:"system"`
	Name    string `json:"name"`
	SortKey string `json:"sort_key"`
}

// Find は指定された識別子を持つゲームシステムのダイスボットのコンストラクタを返す。
// ゲームシステムが見つからなかった場合はエラーを返す。
func Find(gameID string) (dicebot.DiceBotConstructor, error) {
	if gameID == basic.BasicInfo().GameID {
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

	gameIDsStrSlice := sort.StringSlice(gameIDs)
	sort.Sort(gameIDsStrSlice)

	if !includeBasicDiceBot {
		return gameIDs
	}

	gameIDsWithBasicDiceBot := make([]string, 1, len(gameIDs)+1)
	gameIDsWithBasicDiceBot[0] = basic.BasicInfo().GameID
	gameIDsWithBasicDiceBot = append(gameIDsWithBasicDiceBot, gameIDs...)

	return gameIDsWithBasicDiceBot
}

// AvailableGameInfos は利用可能なゲームシステムの情報のスライスを返す。
func AvailableGameInfos(includeBasicDiceBot bool) Names {
	games := make([]SystemInfo, 0, len(gameIDToDiceBotConstructor))
	for _, k := range gameIDToDiceBotConstructor {
		info := SystemInfo{
			System: k().GameID(),
			Name: k().GameName(),
			SortKey: k().SortKey(),
		}
		games = append(games, info)
	}
	sort.Slice(games, func (i, j int) bool { return games[i].SortKey < games[j].SortKey })

	if !includeBasicDiceBot {
		return Names{Name: games}
	}
	gamesWithBasicDiceBot := make([]SystemInfo, 1, len(games)+1)
	info := SystemInfo{
		System: basic.BasicInfo().GameID,
		Name: basic.BasicInfo().GameName,
		SortKey: basic.BasicInfo().SortKey,
	}
	gamesWithBasicDiceBot[0] = info
	gamesWithBasicDiceBot = append(gamesWithBasicDiceBot, games...)
	return Names{Name: gamesWithBasicDiceBot}
}

// ゲーム識別子とダイスボットのコンストラクタとの対応
var gameIDToDiceBotConstructor = map[string]dicebot.DiceBotConstructor{}
