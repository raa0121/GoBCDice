package models

import (
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/pkg/dicebot/list"
)

type Names struct {
	list.Names
}

func NewNames() *Names {
	names := Names{
		Names: list.AvailableGameInfos(true),
	}
	return &names
}

func (s *Names) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"names": s.Names.Name,
	}
}
