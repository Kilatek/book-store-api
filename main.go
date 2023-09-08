package main

import (
	"flag"
	"log"
	"net/http"

	"bookstore.com/api"
	"bookstore.com/config"
	"bookstore.com/domain/service"
	mongorepo "bookstore.com/repository/mongo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "f", "./config/config.yaml", "the path to the config file")
	flag.Parse()

	var err error
	conf, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	repo, err := mongorepo.NewAuthorRepository(conf.DB.URL, conf.DB.Name, conf.DB.Timeout)
	if err != nil {
		panic(err)
	}

	authorSvc := service.NewAuthorService(repo)

	handler := api.NewAuthorHandler(authorSvc)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/authors", func(r chi.Router) {
			r.Get("/{id}", handler.Get)
			r.Post("/", handler.Post)
			r.Put("/{id}", handler.Delete)
			r.Delete("/{id}", handler.Delete)
			r.Get("/", handler.GetAll)
		})
	})

	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
