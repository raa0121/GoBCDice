package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/config"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceAPI/controllers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	return port
}

func main() {
	e := echo.New()

	config.Setup(e)
	controllers.Setup(e)

	err := e.Start(":" + getPort())
	if err != nil {
		panic(err)
	}
}
