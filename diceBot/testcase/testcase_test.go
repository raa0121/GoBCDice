package testcase

import (
	"reflect"
	"testing"
)

var parseTestCases = []struct {
	// テストケースのソース
	source string
	// ゲーム識別子
	gameId string
	// テストケース番号
	index int
	// 期待する値
	expected DiceBotTestCase
	// エラーを期待するか
	err bool
}{
	{
		source: "",
		err:    true,
	},
	{
		source: "input:",
		err:    true,
	},
	{
		source: `input:
2d6+1-1-2-3-4
output:
DiceBot : (2D6+1-1-2-3-4) ＞ 5[4,1]+1-1-2-3-4 ＞ -4
rand:4/6,1/6`,
		gameId: "DiceBot",
		index:  1,
		expected: DiceBotTestCase{
			gameId: "DiceBot",
			index:  1,
			input:  []string{"2d6+1-1-2-3-4"},
			output: "DiceBot : (2D6+1-1-2-3-4) ＞ 5[4,1]+1-1-2-3-4 ＞ -4",
			dice:   []Die{{4, 6}, {1, 6}},
		},
		err: false,
	},
	{
		source: `input:
S2d6
output:
DiceBot : (2D6) ＞ 5[4,1] ＞ 5###secret dice###
rand:4/6,1/6`,
		gameId: "DiceBot",
		index:  2,
		expected: DiceBotTestCase{
			gameId: "DiceBot",
			index:  2,
			input:  []string{"S2d6"},
			output: "DiceBot : (2D6) ＞ 5[4,1] ＞ 5###secret dice###",
			dice:   []Die{{4, 6}, {1, 6}},
		},
		err: false,
	},
	{
		source: `input:
GETSST
output:
Satasupe : サタスペ作成：ベース部品：「大型の金属製の筒」  アクセサリ部品：「ガスボンベや殺虫剤」
部品効果：「命中：8、ダメージ：5、耐久度3、両手」「爆発3」
完成品：サタスペ  （ダメージ＋5・命中8・射撃、「両手」「爆発3」「サタスペ1」「耐久度3」）
rand:6/6,6/6,6/6`,
		gameId: "Satasupe",
		index:  1,
		expected: DiceBotTestCase{
			gameId: "Satasupe",
			index:  1,
			input:  []string{"GETSST"},
			output: `Satasupe : サタスペ作成：ベース部品：「大型の金属製の筒」  アクセサリ部品：「ガスボンベや殺虫剤」
部品効果：「命中：8、ダメージ：5、耐久度3、両手」「爆発3」
完成品：サタスペ  （ダメージ＋5・命中8・射撃、「両手」「爆発3」「サタスペ1」「耐久度3」）`,
			dice: []Die{{6, 6}, {6, 6}, {6, 6}},
		},
		err: false,
	},
	{
		source: `input:
CCT
output:
GranCrest : 国特徴・文化表(13) ＞ 禁欲的
あなたの国民は、道徳を重んじ、常に自分の欲望を制限することが理想的だと考えている。
食料＋４、資金－１
rand:1/6,3/6`,
		gameId: "GranCrest",
		index:  1,
		expected: DiceBotTestCase{
			gameId: "GranCrest",
			index:  1,
			input:  []string{"CCT"},
			output: `GranCrest : 国特徴・文化表(13) ＞ 禁欲的
あなたの国民は、道徳を重んじ、常に自分の欲望を制限することが理想的だと考えている。
食料＋４、資金－１`,
			dice: []Die{{1, 6}, {3, 6}},
		},
		err: false,
	},
}

func TestParse(t *testing.T) {
	for i, test := range parseTestCases {
		actual, err := Parse(test.source, test.gameId, test.index)

		if err != nil {
			if !test.err {
				t.Errorf("#%d:\ngot err: %v", i, err)
			}
			continue
		}

		if test.err {
			t.Errorf("#%d:\nshould err", i)
			continue
		}

		if !reflect.DeepEqual(*actual, test.expected) {
			t.Errorf("#%d:\ngot:  %v\nwant: %v", i, *actual, test.expected)
		}
	}
}

var parseDiceTestCases = []struct {
	source   string
	expected []Die
	err      bool
}{
	{
		source: "1",
		err:    true,
	},
	{
		source: "1/6,1",
		err:    true,
	},
	{
		source: "a1/6",
		err:    true,
	},
	{
		source:   "",
		expected: []Die{},
		err:      false,
	},
	{
		source:   "1/6",
		expected: []Die{{1, 6}},
		err:      false,
	},
	{
		source:   "1/6,2/6,3/6",
		expected: []Die{{1, 6}, {2, 6}, {3, 6}},
		err:      false,
	},
	{
		source:   "1/6, 2/6, 3/6",
		expected: []Die{{1, 6}, {2, 6}, {3, 6}},
		err:      false,
	},
}

func TestParseDice(t *testing.T) {
	for i, test := range parseDiceTestCases {
		actual, err := ParseDice(test.source)

		if err != nil {
			if !test.err {
				t.Errorf("#%d: %s\ngot err: %v", i, test.source, err)
			}
			continue
		}

		if test.err {
			t.Errorf("#%d: %s\nshould err", i, test.source)
			continue
		}

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("#%d: %s\ngot:  %v\nwant: %v", i, test.source, actual, test.expected)
		}
	}
}
