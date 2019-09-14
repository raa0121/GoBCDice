package models

import (
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
)

type RootResponse struct {
	Message string
}

func NewRootResponse() *RootResponse {
	return &RootResponse{
		Message: "Hello. This is GoBCDiceAPI.",
	}
}

func (r *RootResponse) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"message": r.Message,
	}
}
