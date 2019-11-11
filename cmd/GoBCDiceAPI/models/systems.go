package models

import (
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/pkg/dicebot/list"
)

type Systems struct {
	Names []string
}

func NewSystems() *Systems {
	return &Systems{Names: list.AvailableGameIDs(true)}
}

func (s *Systems) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"systems": s.Names,
	}
}
