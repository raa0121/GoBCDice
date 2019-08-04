package basic_test

import (
	dicebottesting "github.com/raa0121/GoBCDice/pkg/dicebot/testing"
	"path/filepath"
	"testing"
)

func TestDiceBot(t *testing.T) {
	testDataFileBaseNames := []string{
		"calc.txt",
		"d_roll_expr.txt",
		"d_roll_comp.txt",
		"b_roll_list.txt",
		"b_roll_comp.txt",
	}

	testDataFiles := joinWithTestData(testDataFileBaseNames)
	dicebottesting.Run("DiceBot", t, testDataFiles...)
}

func joinWithTestData(basenames []string) []string {
	files := make([]string, 0, len(basenames))

	for _, b := range basenames {
		files = append(files, filepath.Join("testdata", b))
	}

	return files
}
