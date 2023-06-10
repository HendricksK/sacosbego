package controllers

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"

	models "github.com/HendricksK/sacosbego/app/models"
)

func GetArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, models.GetArticle(id))
}

func GetArticles(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetArticles())
}