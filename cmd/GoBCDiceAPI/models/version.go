package models

import (
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/pkg/version"
)

type Version struct {
	Api    string
	Bcdice string
}

func NewVersion() *Version {
	return &Version{
		Api:    version.API_VERSION,
		Bcdice: version.BCDICE_VERSION,
	}
}

func (v *Version) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"api":    v.Api,
		"bcdice": v.Bcdice,
	}
}
