package main

import (
	"database/sql"
	"frequentei-api/internal/repository"
	"frequentei-api/internal/usecase"
	"frequentei-api/internal/web"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=root password=password dbname=frequentei sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewLocationRepositoryPostgres(db)
	createLocationUseCase := usecase.NewCreateLocationUseCase(repository)
	listLocationUseCase := usecase.NewListLocationUseCase(repository)

	locationHandler := web.NewLocationHandler(createLocationUseCase, listLocationUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/locations", locationHandler.ListLocationsHandler)
	r.Post("/locations", locationHandler.CreateLocationHandler)

	http.ListenAndServe(":3000", r)
}
