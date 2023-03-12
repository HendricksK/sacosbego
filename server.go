package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/HendricksK/sacosbego/article"
	"github.com/HendricksK/sacosbego/auth"
	"github.com/HendricksK/sacosbego/page"
	"github.com/HendricksK/sacosbego/post"
	"github.com/HendricksK/sacosbego/rider"
	"github.com/HendricksK/sacosbego/track"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	content := []string{
		"API calls<br><br>",
		"/articles<br>",
		"/article/:id<br>",
		"/articleids<br>",
		"/tracks<br>",
		"/track/:id<br>",
		"/trackids<br>",
		"/riders<br>",
		"/rider/:id<br>",
		"/riderids<br>",
		"/page/:id<br>",
		"/posts/:id<br>",
		"/posts/:id/:section<br>",
		"/auth/:name<br>",
	}

	return c.HTML(http.StatusOK, strings.Join(content, " "))
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

func GetSacosUploadToken(c echo.Context) error {
	name := c.Param("name")

	data := auth.GetSacosUploadToken(name)

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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// https://echo.labstack.com/cookbook/cors/#server-using-a-custom-function-to-allow-origins
	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// wth GET, PUT, POST or DELETE method.
	//
	// TODO: work on auth using https://echo.labstack.com/middleware/key-auth/
	// https://webdevstation.com/posts/user-authentication-with-go-using-jwt-token/

	var environment = os.Getenv("ENVIRONMENT")

	log.Println(environment)

	if environment == "local" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://127.0.0.1:8080", "http://localhost:8080"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		}))
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://hendricksk.github.io", "https://cycling.sacoshistory.org", "https://hendricksk.github.io/sacos-dataform", "*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		}))
	}

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

	e.GET("/auth/:name", GetSacosUploadToken)
	// Here ends API calls

	// Port setup for echo webserver
	port, err := GetPort()
	log.Println("starting on port: ", port)

	if err != nil {
		log.Println(err)
	}

	e.Logger.Fatal(e.Start(port))
}

func GetPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return ":80", nil
	}
	return ":" + port, nil
}
