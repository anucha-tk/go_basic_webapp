package main

import (
	"log"
	"net/http"
	"time"

	"github.com/anucha-tk/go_basic_webapp/pkg/config"
	"github.com/anucha-tk/go_basic_webapp/pkg/handlers"
	"github.com/anucha-tk/go_basic_webapp/pkg/render"
)

func main() {
	const port = ":8080"
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	log.Println("run on port", port)
	srv := http.Server{
		Addr:              "localhost" + port,
		ReadHeaderTimeout: time.Second * 5,
	}
	_ = srv.ListenAndServe()
}
