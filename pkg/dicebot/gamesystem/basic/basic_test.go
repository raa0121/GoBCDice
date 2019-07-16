package basic_test

import (
	dicebottesting "github.com/raa0121/GoBCDice/pkg/dicebot/testing"
	"path/filepath"
	"testing"
)

func TestDiceBot(t *testing.T) {
	testDataPath := filepath.Join("testdata", "DiceBot.txt")
	dicebottesting.Run("DiceBot", testDataPath, t)
}
