package main

import (
	"log"

	routes "github.com/HendricksK/sacosbego/app/routes"
)

func main() {
	log.Println("you're know that you are toxic")

	// app := Config{}

	// log.Printf("Starting broker service on port %s\n", webPort)

	// // define http server
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%s", webPort),
	// 	Handler: routes.routes(),
	// }

	// // start the server
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Panic(err)
	// }

	routes.Routes()
}