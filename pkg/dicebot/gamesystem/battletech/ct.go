package battletech

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/command"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	table "github.com/raa0121/GoBCDice/pkg/core/rollabletable/sparse"
)

var ctTable = table.New(
	"致命的命中表",
	2,
	6,
	table.Item{Max: 7, Content: "致命的命中はなかった"},
	table.Item{Max: 9, Content: "1箇所の致命的命中"},
	table.Item{Max: 11, Content: "2箇所の致命的命中"},
	table.Item{Max: 12, Content: "その部位が吹き飛ぶ（腕、脚、頭）または3箇所の致命的命中（胴）"},
)

// executeCT はCTコマンドを実行する。
func (b *BattleTech) executeCT(e *evaluator.Evaluator) (*command.Result, error) {
	result := &command.Result{
		GameID: b.GameID(),
	}

	rollResult, err := ctTable.Roll(e)
	if err != nil {
		result.AppendMessagePart(err.Error())
		return result, nil
	}

	result.AppendMessagePart(fmt.Sprintf("%d", rollResult.Sum))
	result.AppendMessagePart(rollResult.SelectedItem.Content)

	return result, nil
}
