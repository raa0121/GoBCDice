/*
バトルテック
*/
package battletech

import (
	"fmt"
	"strings"

	"github.com/raa0121/GoBCDice/pkg/core/command"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/dicebot"
)

// basicInfo はダイスボットの基本情報。
var basicInfo = dicebot.DiceBotBasicInfo{
	GameID:   "BattleTech",
	GameName: "バトルテック",
	Usage: `・判定方法
　(回数)BT(ダメージ)(部位)+(基本値)>=(目標値)
　回数は省略時 1固定。
　部位はC（正面）R（右）、L（左）。省略時はC（正面）固定
　U（上半身）、L（下半身）を組み合わせ CU/RU/LU/CL/RL/LLも指定可能
　例）BT3+2>=4
　　正面からダメージ3の攻撃を技能ベース2目標値4で1回判定
　例）2BT3RL+5>=8
　　右下半身にダメージ3の攻撃を技能ベース5目標値8で2回判定
　ミサイルによるダメージは BT(ダメージ)の変わりに SRM2/4/6, LRM5/10/15/20を指定
　例）3SRM6LU+5>=8
　　左上半身にSRM6連を技能ベース5目標値8で3回判定
・CT：致命的命中表
・DW：転倒後の向き表
・CDx：メック戦士意識維持表。ダメージ値xで判定　例）CD3`,
}

// BasicInfo はダイスボットの基本情報を返す。
func BasicInfo() *dicebot.DiceBotBasicInfo {
	return &basicInfo
}

// バトルテックのダイスボット。
type BattleTech struct {
	dicebot.DiceBotImpl
}

// New は新しいダイスボットを構築する。
func New() dicebot.DiceBot {
	return &BattleTech{
		DiceBotImpl: dicebot.DiceBotImpl{
			BasicInfo: BasicInfo(),
		},
	}
}

// ExecuteCommand は指定されたコマンドを実行する。
//
// 基本のダイスボットには特別なコマンドが存在しないため、必ずエラーを返す。
func (b *BattleTech) ExecuteCommand(
	command string,
	evaluator *evaluator.Evaluator,
) (*command.Result, error) {
	switch strings.ToUpper(command) {
	case "CT":
		return b.executeCT(evaluator)
	default:
		return nil, fmt.Errorf("not implemented")
	}
}
