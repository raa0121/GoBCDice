package evaluator

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/object"
)

// evalRRollList は個数振り足しロール列を評価する。
//
// TODO: シャドウラン4版のグリッチに対応する。
func (e *Evaluator) evalRRollList(node *ast.RRollList) (*object.Array, error) {
	if node.Threshold.IsNil() {
		return nil, fmt.Errorf("evalRRollList: threshold is nil")
	}

	// 閾値を評価する
	thresholdObj, evalThresholdErr := e.Eval(node.Threshold)
	if evalThresholdErr != nil {
		return nil, evalThresholdErr
	}
	thresholdInt := thresholdObj.(*object.Integer)
	threshold := thresholdInt.Value

	// ダイスロールのキュー。
	// ダイスロール後、条件を満たす出目が得られるたびに、このキューに
	// そのダイスロールを追加する。
	rollQueue := make([]*ast.VariableInfixExpression, len(node.RRolls))
	copy(rollQueue, node.RRolls)

	// ダイスロール結果を格納する配列
	valueGroups := []object.Object{}
	for i := 0; i < e.MaxRerolls && len(rollQueue) > 0; i++ {
		// キューの最初のダイスロールを取り出す
		rRoll := rollQueue[0]
		if len(rollQueue) < 2 {
			rollQueue = nil
		} else {
			rollQueue = rollQueue[1:len(rollQueue)]
		}

		// ダイスロールを行う
		o, err := e.Eval(rRoll)
		if err != nil {
			return nil, err
		}

		// 出目を結果の配列に格納する
		values := o.(*object.Array)
		valueGroups = append(valueGroups, values)

		// 成功数を数える
		numOfSuccesses := 0
		for _, v := range values.Elements {
			// TODO: 演算子 ">=" 以外も振り足しの条件に設定できるようにする
			if vi := v.(*object.Integer); vi.Value >= threshold {
				numOfSuccesses++
			}
		}

		// 成功した分だけ追加のダイスロールをキューに追加する
		if numOfSuccesses > 0 {
			numNode := ast.NewInt(numOfSuccesses)
			sidesNode := rRoll.Right()
			newRRoll := ast.NewRRoll(numNode, sidesNode)

			rollQueue = append(rollQueue, newRRoll)
		}
	}

	return object.NewArrayByMove(valueGroups), nil
}
