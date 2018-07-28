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

	r := flashcards.Repo{
		DB: db,
	}

	h := handlers.Handler{
		Repo: &r,
	}

	router := httprouter.New()
	router.GET("/", h.Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}