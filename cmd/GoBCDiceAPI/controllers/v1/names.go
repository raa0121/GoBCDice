package v1

import (
	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/models"
)

type NamesController struct {
	Group *echo.Group
}

func NewNamesController(g *echo.Group) *NamesController {
	return &NamesController{
		Group: g,
	}
}

func (controller *NamesController) getNames(c echo.Context) error {
	names := models.NewNames()

	return helpers.JSONResponseObject(c, 200, names)
}

// Setup はコントローラの初期設定を行う。
func (controller *NamesController) Setup() {
	controller.Group.Add("GET", "/names", controller.getNames)
}

