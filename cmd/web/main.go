package main

import (
	"log"
	"net/http"
	"time"

	"github.com/anucha-tk/go_basic_webapp/pkg/handlers"
)

func main() {
	const port = ":8080"
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Println("run on port", port)
	srv := http.Server{
		Addr:              "localhost:" + port,
		ReadHeaderTimeout: time.Second * 5,
	}
	_ = srv.ListenAndServe()
}
