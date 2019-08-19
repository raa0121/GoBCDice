package base

import (
	"github.com/labstack/echo"
)

const PREFIX_V1 = "/v1"

type BaseController struct {
	Router *echo.Router
}

func (controller *BaseController) AddRouteV1(method, path string, h echo.HandlerFunc) {
	controller.Router.Add(method, PREFIX_V1+path, h)
}
