package testing

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// ダイスボットのテストケース。
type DiceBotTestCase struct {
	// ゲーム識別子
	GameID string
	// テストケース番号
	Index int
	// 入力文字列
	Input []string
	// 出力文字列
	Output string
	// 入力するダイス列
	Dice []dice.Die
}

var (
	// テストケースのソースコードを表す正規表現
	sourceRe = regexp.MustCompile("(?s)\\Ainput:\n(.+)\noutput:(.*)\nrand:(.*)")
	// テストケースのソースコード内のダイス表記を表す正規表現
	diceRe = regexp.MustCompile(`\A\s*(\d+)/(\d+)\s*\z`)
)

// Parse はテストケースのソースコードを構文解析し、その内容のDiceBotTestCaseを構築して返す。
// 失敗するとnilを返す。
//
// gameID: にはゲーム識別子,
// index: テストケース番号。
func Parse(source string, gameID string, index int) (*DiceBotTestCase, error) {
	matches := sourceRe.FindStringSubmatch(source)
	if matches == nil {
		return nil, fmt.Errorf("Parse: %s#%d: テストケース構文エラー", gameID, index)
	}

	ds, err := ParseDice(matches[3])
	if err != nil {
		return nil, err
	}

	input := strings.Split(matches[1], "\n")
	output := strings.TrimLeft(matches[2], "\n")

	return &DiceBotTestCase{
		GameID: gameID,
		Index:  index,
		Input:  input,
		Output: output,
		Dice:   ds,
	}, nil
}

// ParseDice はテストケースのダイス表記を解析し、振られたダイスのスライスを返す。
func ParseDice(source string) ([]dice.Die, error) {
	rolledDice := []dice.Die{}

	if source == "" {
		return rolledDice, nil
	}

	diceStrs := strings.Split(source, ",")
	for i, diceStr := range diceStrs {
		matches := diceRe.FindStringSubmatch(diceStr)
		if matches == nil {
			return nil, fmt.Errorf("ParseDice: #%d: %q: ダイス構文エラー", i+1, diceStr)
		}

		Value, _ := strconv.Atoi(matches[1])
		Sides, _ := strconv.Atoi(matches[2])
		rolledDice = append(rolledDice, dice.Die{Value, Sides})
	}

	return rolledDice, nil
}

// ParseFile はテストデータファイルを解析し、テストケースのスライスを返す。
//
// filename: テストデータファイルのパス,
// gameID: ゲーム識別子。
func ParseFile(filename string, gameID string) ([]*DiceBotTestCase, error) {
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := strings.TrimRight(string(contentBytes), "\n")
	testCaseSources := strings.Split(content, "\n============================\n")
	testCases := []*DiceBotTestCase{}

	for i, source := range testCaseSources {
		index := i + 1

		testCase, err := Parse(source, gameID, index)
		if err != nil {
			return nil, err
		}

		testCases = append(testCases, testCase)
	}

	return testCases, nil
}
