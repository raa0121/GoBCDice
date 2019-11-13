package table

import (
	"sort"

	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
)

// SparseTableItem は疎らな表の項目の構造体。
type SparseTableItem struct {
	// Max は該当する出目の最大値。
	Max int
	// Content は項目の内容。
	Content string
}

// SparseTable は項目定義で最大値を指定する疎らな表の構造体。
type SparseTable struct {
	// name は表の名前。
	name string
	// numOfDice は振るダイス数。
	numOfDice int
	// sidesOfDie はダイスの面数。
	sidesOfDie int
	// items は表の項目のスライス。
	items []SparseTableItem
}

// SparseTableRollResult は疎らな表から引いた結果の構造体。
type SparseTableRollResult struct {
	// Sum は出目の合計。
	Sum int
	// SelectedItem は引かれた項目。
	SelectedItem SparseTableItem
}

// NewSparseTable は新しい疎らな表を返す。
func NewSparseTable(
	name string,
	numOfDice int,
	sidesOfDie int,
	items ...SparseTableItem,
) *SparseTable {
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
	copiedItems := make([]SparseTableItem, len(items))
	copy(copiedItems, items)

	sort.Slice(copiedItems, func(i int, j int) bool {
		return copiedItems[i].Max < copiedItems[j].Max
	})

	return &SparseTable{
		name:       name,
		numOfDice:  numOfDice,
		sidesOfDie: sidesOfDie,
		items:      copiedItems,
	}
}

// Name は表の名前を返す。
func (t *SparseTable) Name() string {
	return t.name
}

// Get は指定された値に対応する項目を返す。
//
// 最大値が指定された値以上で、かつ最も近い項目を返す。
// 最大値が指定された値未満の項目しかなかった場合、最大値が最も大きい項目を返す。
func (t *SparseTable) Get(value int) SparseTableItem {
	for _, item := range t.items {
		if value <= item.Max {
			return item
		}
	}

	return t.items[len(t.items)-1]
}

// Roll は、ダイスを振って表から項目を引く。
func (t *SparseTable) Roll(e *evaluator.Evaluator) (SparseTableRollResult, error) {
	rolledDice, err := e.RollDice(t.numOfDice, t.sidesOfDie)
	if err != nil {
		return SparseTableRollResult{}, err
	}

	sum := 0
	for _, d := range rolledDice {
		sum += d.Value
	}

	return SparseTableRollResult{
		Sum:          sum,
		SelectedItem: t.Get(sum),
	}, nil
}
