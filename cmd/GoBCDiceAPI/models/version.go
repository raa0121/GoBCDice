package models

import "github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"

type Version struct {
	Api    string
	Bcdice string
}

func NewVersion() *Version {
	return &Version{Api: "0.0.1", Bcdice: "0.0.0"}
}

func (v *Version) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"api":    v.Api,
		"bcdice": v.Bcdice,
	}
}
