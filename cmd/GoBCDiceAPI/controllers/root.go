package controllers

import (
	"github.com/labstack/echo"

	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/models"
)

// RootController is the Root controller.
type RootController struct {
	Router *echo.Router
}

func (controller *RootController) getRoot(c echo.Context) error {
	r := models.NewRootResponse()

	return helpers.JSONResponseObject(c, 200, r)
}

// Setup sets up routes for the Root controller.
func (controller *RootController) Setup() {
	controller.Router.Add("GET", "/", controller.getRoot)
}
