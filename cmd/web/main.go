package main

import (
	"log"
	"net/http"

	"github.com/anucha-tk/go_basic_webapp/pkg/handlers"
)

func main() {
	const port = ":8080"
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Println("run on port", port)
	_ = http.ListenAndServe(port, nil)
}
