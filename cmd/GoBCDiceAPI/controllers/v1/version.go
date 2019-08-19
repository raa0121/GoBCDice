package v1

import (
	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/controllers/base"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/models"
)

type VersionController struct {
	Router *base.BaseController
}

func (controller *VersionController) getVersion(c echo.Context) error {
	c.Request().ParseForm()

	version := models.NewVersion()
	versionResponse := version.ToResponseMap()

	return helpers.JSONResponse(c, 200, versionResponse)
}

func (controller *VersionController) Setup() {
	controller.Router.AddRouteV1("GET", "/version", controller.getVersion)
}
