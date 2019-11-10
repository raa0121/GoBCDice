package controllers

import (
	"github.com/labstack/echo"

	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/helpers"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/models"
)

// RootController はルート階層のコントローラ。
type RootController struct {
	Router *echo.Router
}

// NewRootController は新しいRootControllerを返す。
func NewRootController(router *echo.Router) *RootController {
	return &RootController{
		Router: router,
	}
}

// getRoot は正常に動作していることを表す応答を返す。
func (controller *RootController) getRoot(c echo.Context) error {
	r := models.NewRootResponse()

	return helpers.JSONResponseObject(c, 200, r)
}

// Setup はコントローラの初期設定を行う。
func (controller *RootController) Setup() {
	controller.Router.Add("GET", "/", controller.getRoot)
}
