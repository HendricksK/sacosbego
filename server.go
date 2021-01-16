package main

import (
	"net/http"
	"strconv"
	"log"
	"fmt"
	"os"

	"io/ioutil"
	"github.com/HendricksK/sacosbego/article"
	"github.com/HendricksK/sacosbego/page"
	"github.com/HendricksK/sacosbego/post"
	"github.com/HendricksK/sacosbego/track"
	"github.com/HendricksK/sacosbego/rider"

	"github.com/labstack/echo/v4"
)

/**
* Base calls, returns data from a file.
*/


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func GetApiCalls(c echo.Context) error {
	content, err := ioutil.ReadFile("api.html")
    if err != nil {
        log.Fatal(err)
    }
    text := string(content)

    return c.HTML(http.StatusOK, text)
} 

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

// /**
//  * Track API calls
//  */

func GetTracks(c echo.Context) error {
	data := track.GetTracks()
	return c.JSON(http.StatusOK, data)
}

func GetTrack(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    log.Println(err)
	}
	
	data := track.GetTrack(id)

	return c.JSON(http.StatusOK, data)
}

func GetTrackIds(c echo.Context) error {
	data := track.GetTrackIds()
	return c.JSON(http.StatusOK, data)
}


/**
 * Rider API calls
 */

func GetRiders(c echo.Context) error {
	data := rider.GetRiders()
	return c.JSON(http.StatusOK, data)
}

func GetRider(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	    log.Println(err)
	}
	
	data := rider.GetRider(id)

	return c.JSON(http.StatusOK, data)
}

func GetRiderIds(c echo.Context) error {
	data := rider.GetRiderIds()
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
	e.GET("/", GetApiCalls)
	e.GET("/articles", GetArticles)
	e.GET("/article/:id", GetArticle)
	e.GET("/articleids", GetArticleIds)

	e.GET("/tracks", GetTracks)
	e.GET("/track/:id", GetTrack)
	e.GET("/trackids", GetTrackIds)

	e.GET("/riders", GetRiders)
	e.GET("/rider/:id", GetRider)
	e.GET("/riderids", GetRiderIds)

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