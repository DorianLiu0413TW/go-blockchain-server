package main

import (
	"time"

	"go-blockchain-api/apierror"
	"go-blockchain-api/config"
	"go-blockchain-api/database"
	"go-blockchain-api/log"
	"go-blockchain-api/util"
	
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.Global.AllowOrigins,
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders: []string{
			"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization",
			"X-DeviceToken", "X-ActivateCode", "X-MFAToken",
		},
		ExposeHeaders:    []string{"Content-Length", "Pagination-Count", "Pagination-Limit", "Pagination-Page"},
		AllowCredentials: true,
	}))

	err = database.InitDB()
	if err != nil {
		log.Fatal(log.LabelStartup, "Failed to start. Database ", err)
	}
	
	e.Use(middleware.Gzip());
	util.SetCurrentTimeFunc(func() time.Time {
		return time.Now().In(config.Global.TimeZone)
	})

	e.HTTPErrorHandler = apierror.HTTPErrorHandler

}