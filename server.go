package main

import (
	"net/http"
	"strconv"
	"log"
	"fmt"
	"os"
	"github.com/HendricksK/sacosbego/article"

	"github.com/labstack/echo/v4"
)

func getArticles(c echo.Context) error {
	data := article.GetArticles()
	return c.JSON(http.StatusOK, data)
}

func getArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    log.Println(err)
	}
	
	data := article.GetArticle(id)

	return c.String(http.StatusOK, data.Article_data)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "...Sharon, I'm good at stuff\nAnd you're into stuff (wooh)\nLet's make products - Vulfpeck wait for the moment,\nhttps://genius.com/Vulfpeck-wait-for-the-moment-lyrics")
	})

	e.GET("/articles", getArticles)
	e.GET("/article/:id", getArticle)

	port, err := getPort()
	if err != nil {
	    log.Println(err)
	}
	
	e.Logger.Fatal(e.Start(port))
}

func getPort() (string, error) {
  // the PORT is supplied by Heroku
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}