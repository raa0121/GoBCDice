package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// Environment is the environment for the application
	Environment = os.Getenv("ECHO_ENV")
)

func Setup(e *echo.Echo) {
	if Environment == "" {
		Environment = "development"
	}

	if Environment == "production" {
		tmpdir := filepath.Join(os.TempDir(), "GoBCDiceAPI")
		os.MkdirAll(tmpdir, 0700)
		logFileName := filepath.Join(tmpdir, fmt.Sprintf("%s-%d.log", Environment, time.Now().Unix()))
		f, err := os.Create(logFileName)
		if err != nil {
			panic(err)
		}
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: f,
		}))
	} else {
		e.Use(middleware.Logger())
	}
}

// vi:syntax=go
