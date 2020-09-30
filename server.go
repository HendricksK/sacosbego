package main

import (
	"net/http"
	"strconv"
	"log"
	"fmt"
	"github.com/HendricksK/sacosbego/article"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func getArticles(c echo.Context) error {
	return c.String(http.StatusOK, "thats an opportunity")
}

func getArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
	    log.Println(err)
	}
	
	data := article.GetArticle(id)
	fmt.Print(data)
	return c.String(http.StatusOK, data.Article_data)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "...Sharon, I'm good at stuff\nAnd you're into stuff (wooh)\nLet's make products - Vulfpeck wait for the moment,\nhttps://genius.com/Vulfpeck-wait-for-the-moment-lyrics")
	})

	e.GET("/articles", getArticles)
	e.GET("/article/:id", getArticle)

	e.Logger.Fatal(e.Start(":1323"))
}
