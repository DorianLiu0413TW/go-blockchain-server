package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	err := config.InitConfig()
	if err!= nil {
		log.Fatal(log.LabelStartup, "Failed to start!", err)
	}

	e.Use(middleware.Gzip());
}