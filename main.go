package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", storeIndex)
	http.ListenAndServe(":8081", nil)
}

func storeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here lies an API call")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
}
