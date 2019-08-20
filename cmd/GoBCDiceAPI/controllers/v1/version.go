package v1

import (
	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/models"
)

// VersionController はGoBCDiceAPIのバージョン情報を返すコントローラ。
type VersionController struct {
	Group *echo.Group
}

// NewVersionController は新しいVersionControllerを返す。
func NewVersionController(g *echo.Group) *VersionController {
	return &VersionController{
		Group: g,
	}
}

// getVersion はGoBCDiceAPIのバージョン情報を返す。
func (controller *VersionController) getVersion(c echo.Context) error {
	version := models.NewVersion()

	return helpers.JSONResponseObject(c, 200, version)
}

// Setup はコントローラの初期設定を行う。
func (controller *VersionController) Setup() {
	controller.Group.Add("GET", "/version", controller.getVersion)
}
