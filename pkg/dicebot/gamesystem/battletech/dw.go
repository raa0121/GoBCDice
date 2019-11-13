package battletech

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/command"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/table"
)

// dwTable は転倒後の向き表。
var dwTable = table.NewTable(
	"転倒後の向き表",
	1,
	6,
	"同じ（前面から転倒） 正面／背面",
	"1ヘクスサイド右（側面から転倒） 右側面",
	"2ヘクスサイド右（側面から転倒） 右側面",
	"180度逆（背面から転倒） 正面／背面",
	"2ヘクスサイド左（側面から転倒） 左側面",
	"1ヘクスサイド左（側面から転倒） 左側面",
)

func (b *BattleTech) executeDW(e *evaluator.Evaluator) (*command.Result, error) {
	result := &command.Result{
		GameID: b.GameID(),
	}

	rollResult, err := dwTable.Roll(e)
	if err != nil {
		result.AppendMessagePart(err.Error())
		return result, nil
	}

	result.RolledDice = e.RolledDice()

	result.AppendMessagePart(fmt.Sprintf("%d", rollResult.Sum))
	result.AppendMessagePart(rollResult.SelectedItem)

	return result, nil
}
