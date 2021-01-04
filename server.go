package main

import (
	"net/http"
	"strconv"
	"log"
	"fmt"
	"os"
	"github.com/HendricksK/sacosbego/article"
	"github.com/HendricksK/sacosbego/page"
	"github.com/HendricksK/sacosbego/post"

	"github.com/labstack/echo/v4"
)

/**
 * Article API calls
 */

func GetArticles(c echo.Context) error {
	data := article.GetArticles()
	return c.JSON(http.StatusOK, data)
}

func GetArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    log.Println(err)
	}
	
	data := article.GetArticle(id)

	return c.JSON(http.StatusOK, data)
}

func GetArticleIds(c echo.Context) error {
	data := article.GetArticlesIds()
	return c.JSON(http.StatusOK, data)
}

/**
 * Page API calls
 */

func GetPage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    log.Println(err)
	}
	
	data := page.GetPage(id)

	return c.JSON(http.StatusOK, data)	
}

/**
 * Post API calls
 */

func GetPosts(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    log.Println(err)
	}
	
	data := post.GetPosts(id)

	return c.JSON(http.StatusOK, data)	
}

func GetPostSection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	section := c.Param("section")
	if err != nil {
	    log.Println(err)
	}
	
	data := post.GetPostSection(id, section)

	return c.JSON(http.StatusOK, data)
}

/**
* Main function call init echo server
* Create our API calls as well
* Setup our PORT	
*/
func main() {

	// Echo init
	e := echo.New()
	// Here lies API calls
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "...Sharon, I'm good at stuff\nAnd you're into stuff (wooh)\nLet's make products - Vulfpeck wait for the moment,\nhttps://genius.com/Vulfpeck-wait-for-the-moment-lyrics")
	}) 
	e.GET("/articles", GetArticles)
	e.GET("/article/:id", GetArticle)
	e.GET("/articleids", GetArticleIds)
	e.GET("/page/:id", GetPage)
	e.GET("/posts/:id", GetPosts)
	e.GET("/posts/:id/:section", GetPostSection)
	// Here ends API calls

	// Port setup for echo webserver
	port, err := GetPort()
	if err != nil {
	    log.Println(err)
	}
	
	e.Logger.Fatal(e.Start(port))
}

func GetPort() (string, error) {
  // the PORT is supplied by Heroku
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("PORT not set")
  }
  return ":" + port, nil
}