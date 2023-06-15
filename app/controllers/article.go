package controllers

import (
	// "log"
	"net/http"
	"github.com/labstack/echo/v4"
	// "strconv"

	models "github.com/HendricksK/sacosbego/app/models"
)

func getHTTPStatus() {

}

func GetArticle(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, models.GetArticle(id))
}

func GetArticles(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetArticles())
}

func CreateArticle(c echo.Context) error {
	return c.JSON(http.StatusCreated, models.CreateArticle(c))
}
