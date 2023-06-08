package controllers

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"

	models "github.com/HendricksK/sacosbego/app/models"
)

func GetArticle(c echo.Context) error {
	log.Println("this is an article")
	
	return c.JSON(http.StatusOK, models.GetArticle())
}