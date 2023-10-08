package controllers

import (
	// "log"
	"net/http"
	"github.com/labstack/echo/v4"
	// "strconv"

	models "github.com/HendricksK/sacosbego/app/models"
)

func GetRider(c echo.Context) error {
	id := c.Param("id")
	response := models.GetRider(id)
	httpStatusCode := http.StatusOK

	if response.Id == nil {
		httpStatusCode = http.StatusNotFound
	} 

	return c.JSON(httpStatusCode, response)
}

func GetRiders(c echo.Context) error {
	// For now, not enough data to stress about pagination

	response := models.GetRiders()
	httpStatusCode := http.StatusOK

	if len(response) == 0 {
		httpStatusCode = http.StatusNotFound
	} 

	return c.JSON(httpStatusCode, response)
}

func CreateRider(c echo.Context) error {
	httpStatusCode := models.CreateRider(c)
	return c.JSON(httpStatusCode, 0)
}


func PatchRider(c echo.Context) error {
	httpStatusCode := models.UpdateRider(c)
	return c.JSON(httpStatusCode, 0)
}

func DeleteRider(c echo.Context) error {
	httpStatusCode := models.DeleteRider(c)
	return c.JSON(httpStatusCode, 0)
}