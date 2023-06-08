package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"


	controllers "github.com/HendricksK/sacosbego/app/controllers"
)

func Routes() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", controllers.Ping)
	e.GET("/healthz", controllers.Healthz)
	e.GET("/article/:id", controllers.GetArticle)

	e.Logger.Fatal(e.Start("localhost:9001"))
}