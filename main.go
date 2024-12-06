package main

import (
	"mindarc/mindarc-reporting/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	mainHandler := handler.MainHandler{}
	e.GET("/", mainHandler.HandleMain)
	e.GET("/t", mainHandler.HandleTest)

	e.Logger.Fatal(e.Start(":3000"))
}
