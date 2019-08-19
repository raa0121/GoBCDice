package controllers

import (
	"github.com/labstack/echo"
)

// Setup sets up all controllers.
func Setup(router *echo.Router) {
	cRoot := RootController{Router: router}
	cRoot.Setup()
}
