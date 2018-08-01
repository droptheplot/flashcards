package main

import (
	"log"
	"net/http"
	"os"

	"github.com/droptheplot/flashcards/handlers"
	"github.com/droptheplot/flashcards/repositories/db"
	"github.com/droptheplot/flashcards/repositories/kv"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	pg, err := sqlx.Open("postgres", os.Getenv("DB"))

	if err != nil {
		log.Fatal(err)
	}

	defer pg.Close()

	rd := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	defer rd.Close()

	h := handlers.Handler{
		DBRepository: &db.Repository{pg},
		KVRepository: &kv.Repository{rd},
	}

	router := httprouter.New()
	router.GET("/api/v1/sources", h.GetSources)
	router.GET("/api/v1/sources/:id", h.GetSourceByID)
	router.POST("/api/v1/users", h.CreateUser)
	router.POST("/api/v1/tokens", h.CreateToken)

	log.Fatal(http.ListenAndServe(":8080", router))
}
