package controllers

import (
	// "log"
	"net/http"
	"github.com/labstack/echo/v4"
	// "strconv"
	// "strings"
	models "github.com/HendricksK/sacosbego/app/models"
)

func GetPage(c echo.Context) error {
	id := c.Param("id")

	response := models.GetPage(id)
	httpStatusCode := http.StatusOK

	// if len(response) == 0 {
	// 	httpStatusCode = http.StatusNotFound
	// } 

	return c.JSON(httpStatusCode, response)
}
