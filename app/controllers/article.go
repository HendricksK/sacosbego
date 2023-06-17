package controllers

import (
	// "log"
	"net/http"
	"github.com/labstack/echo/v4"
	// "strconv"

	models "github.com/HendricksK/sacosbego/app/models"
)

func GetArticle(c echo.Context) error {
	id := c.Param("id")
	response := models.GetArticle(id)
	httpStatusCode := http.StatusOK

	if response.Id == nil {
		httpStatusCode = http.StatusNotFound
	} 

	return c.JSON(httpStatusCode, response)
}

func GetArticles(c echo.Context) error {
	// For now, not enough data to stress about pagination

	response := models.GetArticles()
	httpStatusCode := http.StatusOK

	if len(response) == 0 {
		httpStatusCode = http.StatusNotFound
	} 

	return c.JSON(httpStatusCode, response)
}

func CreateArticle(c echo.Context) error {
	httpStatusCode := models.CreateArticle(c)
	return c.JSON(httpStatusCode, 0)
}


func PatchArticle(c echo.Context) error {
	httpStatusCode := models.UpdateArticle(c)
	return c.JSON(httpStatusCode, 0)
}

func DeleteArticle(c echo.Context) error {
	httpStatusCode := models.DeleteArticle(c)
	return c.JSON(httpStatusCode, 0)
}