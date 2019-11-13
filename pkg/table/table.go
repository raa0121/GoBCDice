package table

import (
	"fmt"

	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
)

type Table struct {
	// name は表の名前。
	name string
	// numOfDice は振るダイス数。
	numOfDice int
	// sidesOfDie はダイスの面数。
	sidesOfDie int
	// items は表の項目のスライス。
	items []string
}

type TableRollResult struct {
	// Sum は出目の合計。
	Sum int
	// SelectedItem は引かれた項目。
	SelectedItem string
}

func NewTable(name string, numOfDice int, sidesOfDie int, items ...string) *Table {
	if numOfDice < 1 {
		panic("numOfDice must be positive")
	}

	if sidesOfDie < 1 {
		panic("sidesOfDie must be positive")
	}

	expectedItemsLength := (numOfDice-1)*sidesOfDie + 1
	if len(items) != expectedItemsLength {
		message := fmt.Sprintf(
			"wrong length of items: got=%d, want=%d",
			len(items),
			expectedItemsLength,
		)
		panic(message)
	}

	// 元の項目のスライスを破壊しないようにコピーする
	copiedItems := make([]string, len(items))
	copy(copiedItems, items)

	return &Table{
		name:       name,
		numOfDice:  numOfDice,
		sidesOfDie: sidesOfDie,
		items:      copiedItems,
	}
}

// Name は表の名前を返す。
func (t *Table) Name() string {
	return t.name
}

// Get は指定された値に対応する項目を返す。
//
// 最大値が指定された値以上で、かつ最も近い項目を返す。
// 最大値が指定された値未満の項目しかなかった場合、最大値が最も大きい項目を返す。
func (t *Table) Get(value int) string {
	return t.items[value-t.numOfDice]
}

// Roll は、ダイスを振って表から項目を引く。
func (t *Table) Roll(e *evaluator.Evaluator) (TableRollResult, error) {
	rolledDice, err := e.RollDice(t.numOfDice, t.sidesOfDie)
	if err != nil {
		return TableRollResult{}, err
	}

	sum := 0
	for _, d := range rolledDice {
		sum += d.Value
	}

	return TableRollResult{
		Sum:          sum,
		SelectedItem: t.Get(sum),
	}, nil
}
