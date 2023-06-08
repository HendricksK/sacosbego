package controllers

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	log.Println("Pong")
	return c.JSON(http.StatusOK, "oh hi mark")
}

func Healthz(c echo.Context) error {
	log.Println("Healthz")
	return c.JSON(http.StatusOK, "{'name':'sacos-api'}")
}