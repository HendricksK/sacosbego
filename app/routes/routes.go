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

	// HEALTH
	e.GET("/ping", controllers.Ping)
	e.GET("/healthz", controllers.Healthz)

	// ARTICLES
	e.GET("/article/:id", controllers.GetArticle)
	e.GET("/articles", controllers.GetArticles)
	e.POST("/article", controllers.CreateArticle)



	e.Logger.Fatal(e.Start("localhost:9002"))
}