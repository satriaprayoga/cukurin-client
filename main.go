package main

import (
	"cukurin-client/pkg/dbpostgres"
	"cukurin-client/pkg/setting"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	setting.Setup()
	dbpostgres.Setup()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	sPort := fmt.Sprintf(":%d", setting.FileConfigSetting.Server.HTTPPort)
	e.Logger.Fatal(e.Start(sPort))
}
