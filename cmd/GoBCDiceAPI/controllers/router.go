package controllers

import (
	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/controllers/base"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/controllers/v1"
)

var (
	VERSION_1_PREFIX = "/v1"
)

// Setup sets up all controllers.
func Setup(router *echo.Router) {
	version := v1.VersionController{&base.BaseController{router}}
	version.Setup()
}

// vi:syntax=go
