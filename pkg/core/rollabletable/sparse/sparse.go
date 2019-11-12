package sparse

import (
	"sort"

	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
)

// Item は表の項目の構造体。
type Item struct {
	// Max は該当する出目の最大値。
	Max int
	// Content は項目の内容。
	Content string
}

// Table は項目定義で最大値を指定する表の構造体。
type Table struct {
	// Name は表の名前。
	name string
	// NumOfDice は振るダイス数。
	numOfDice int
	// SidesOfDie はダイスの面数。
	sidesOfDie int
	// Items は表の項目のスライス。
	items []Item
}

type RollResult struct {
	Sum          int
	SelectedItem Item
}

// New は新しい表を返す。
func New(name string, numOfDice int, sidesOfDie int, items ...Item) *Table {
	if numOfDice < 1 {
		panic("numOfDice must be positive")
	}

	if sidesOfDie < 1 {
		panic("sidesOfDie must be positive")
	}

	if len(items) < 1 {
		panic("items must not be empty")
	}

	// 元の項目のスライスを破壊しないようにコピーする
	copiedItems := make([]Item, len(items))
	copy(copiedItems, items)

	sort.Slice(copiedItems, func(i int, j int) bool {
		return copiedItems[i].Max < copiedItems[j].Max
	})

	return &Table{
		name:       name,
		numOfDice:  numOfDice,
		sidesOfDie: sidesOfDie,
		items:      copiedItems,
	}
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) Get(value int) Item {
	for _, item := range t.items {
		if value <= item.Max {
			return item
		}
	}

	return t.items[len(t.items)-1]
}

func (t *Table) Roll(e *evaluator.Evaluator) (RollResult, error) {
	rolledDice, err := e.RollDice(t.numOfDice, t.sidesOfDie)
	if err != nil {
		return RollResult{}, err
	}

	sum := 0
	for _, d := range rolledDice {
		sum += d.Value
	}

	return RollResult{
		Sum:          sum,
		SelectedItem: t.Get(sum),
	}, nil
}
