package v1

import (
	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/models"
)

type SystemsController struct {
	Group *echo.Group
}

func NewSystemsController(g *echo.Group) *SystemsController {
	return &SystemsController{
		Group: g,
	}
}

func (controller *SystemsController) getSystems(c echo.Context) error {
	systems := models.NewSystems()

	return helpers.JSONResponseObject(c, 200, systems)
}

// Setup はコントローラの初期設定を行う。
func (controller *SystemsController) Setup() {
	controller.Group.Add("GET", "/systems", controller.getSystems)
}
