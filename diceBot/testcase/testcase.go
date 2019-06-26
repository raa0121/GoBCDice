package testcase

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// サイコロを表す構造体
type Die struct {
	// 出目
	result int
	// サイコロの面の数
	sides int
}

// ダイスボットのテストケースを表す構造体
type DiceBotTestCase struct {
	// ゲーム識別子
	gameId string
	// テストケース番号
	index int
	// 入力文字列
	input []string
	// 出力文字列
	output string
	// 入力するダイス列
	dice []Die
}

var (
	sourceRe = regexp.MustCompile("(?s)\\Ainput:\n(.+)\noutput:(.*)\nrand:(.*)")
	diceRe   = regexp.MustCompile(`\A\s*(\d+)/(\d+)\s*\z`)
)

func Parse(source string, gameId string, index int) (*DiceBotTestCase, error) {
	matches := sourceRe.FindStringSubmatch(source)
	if matches == nil {
		return nil, fmt.Errorf("Parse: %s#%d: テストケース構文エラー", gameId, index)
	}

	dice, err := ParseDice(matches[3])
	if err != nil {
		return nil, err
	}

	input := strings.Split(matches[1], "\n")
	output := strings.TrimLeft(matches[2], "\n")

	return &DiceBotTestCase{
		gameId: gameId,
		index:  index,
		input:  input,
		output: output,
		dice:   dice,
	}, nil
}

func ParseDice(source string) ([]Die, error) {
	dice := []Die{}

	if source == "" {
		return dice, nil
	}

	diceStrs := strings.Split(source, ",")
	for i, diceStr := range diceStrs {
		matches := diceRe.FindStringSubmatch(diceStr)
		if matches == nil {
			return nil, fmt.Errorf("ParseDice: #%d: %q: ダイス構文エラー", i+1, diceStr)
		}

		result, _ := strconv.Atoi(matches[1])
		sides, _ := strconv.Atoi(matches[2])
		dice = append(dice, Die{result, sides})
	}

	return dice, nil
}

func ParseFile(filename string) ([]*DiceBotTestCase, error) {
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	basename := path.Base(filename)
	gameId := strings.TrimSuffix(basename, path.Ext(basename))

	content := strings.TrimRight(string(contentBytes), "\n")
	testCaseSources := strings.Split(content, "\n============================\n")
	testCases := []*DiceBotTestCase{}

	for i, source := range testCaseSources {
		index := i + 1

		testCase, err := Parse(source, gameId, index)
		if err != nil {
			return nil, err
		}

		testCases = append(testCases, testCase)
	}

	return testCases, nil
}
