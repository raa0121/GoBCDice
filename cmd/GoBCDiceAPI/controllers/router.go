package controllers

import (
	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/controllers/v1"
)

// Setup はすべてのコントローラの初期設定を行う。
func Setup(e *echo.Echo) {
	root := RootController{Router: e.Router()}
	root.Setup()

	gV1 := e.Group("/v1")
	setupV1(gV1)
}

// setupV1 は v1/ 以下のコントローラの初期設定を行う。
func setupV1(g *echo.Group) {
	version := v1.NewVersionController(g)
	version.Setup()
}
