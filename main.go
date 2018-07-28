package main

import (
	"log"
	"net/http"
	"os"

	"github.com/droptheplot/flashcards/handlers"
	"github.com/droptheplot/flashcards/repositories/flashcards"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("DB"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	f := flashcards.Repository{
		DB: db,
	}

	h := handlers.Handler{
		Repository: &f,
	}

	router := httprouter.New()
	router.GET("/api/v1/sources", h.GetSources)
	router.GET("/api/v1/sources/:id", h.GetSourceByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
