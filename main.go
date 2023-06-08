package main

import (
	"fmt"
	"log"
	"net/http"

	routes "routes/routes"
)

const webPort = "81"

type Config struct{}

func main() {
	log.Println("you're know that you are toxic")

	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}