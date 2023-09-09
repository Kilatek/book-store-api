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

	authorRepo, err := mongorepo.NewAuthorRepository(conf.DB.URL, conf.DB.Name, conf.DB.Timeout)
	if err != nil {
		panic(err)
	}

	bookRepo, err := mongorepo.NewBookRepository(conf.DB.URL, conf.DB.Name, conf.DB.Timeout)
	if err != nil {
		panic(err)
	}

	authorSvc := service.NewAuthorService(authorRepo)
	bookSvc := service.NewBookService(bookRepo, authorRepo)

	authorHandler := api.NewAuthorHandler(authorSvc)
	bookHandler := api.NewBookHandler(bookSvc)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/authors", func(r chi.Router) {
			r.Get("/{id}", authorHandler.Get)
			r.Post("/", authorHandler.Post)
			r.Put("/{id}", authorHandler.Put)
			r.Delete("/{id}", authorHandler.Delete)
			r.Get("/", authorHandler.GetAll)
		})
		r.Route("/books", func(r chi.Router) {
			r.Get("/{id}", bookHandler.Get)
			r.Post("/", bookHandler.Post)
			r.Put("/{id}", bookHandler.Put)
			r.Delete("/{id}", bookHandler.Delete)
			r.Get("/", bookHandler.GetAll)
		})
	})

	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
