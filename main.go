package main

import (
	"log"
	"net/http"
	"os"

	sourceHandler "github.com/droptheplot/flashcards/handlers/source"
	userHandler "github.com/droptheplot/flashcards/handlers/user"
	"github.com/droptheplot/flashcards/repositories/db"
	"github.com/droptheplot/flashcards/repositories/kv"
	sourceUseCase "github.com/droptheplot/flashcards/usecases/source"
	userUseCase "github.com/droptheplot/flashcards/usecases/user"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	pg, err := sqlx.Open("postgres", os.Getenv("POSTGRES"))

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

	dbr := db.Repository{DB: pg}
	kvr := kv.Repository{DB: rd}

	uh := userHandler.Handler{
		UseCase: &userUseCase.UseCase{
			DBRepository: &dbr,
			KVRepository: &kvr,
		},
	}

	sh := sourceHandler.Handler{
		UseCase: &sourceUseCase.UseCase{
			DBRepository: &dbr,
		},
	}

	router := httprouter.New()

	router.GET("/api/v1/sources", sh.GetSources)
	router.GET("/api/v1/sources/:id", sh.GetSourceByID)

	router.POST("/api/v1/users", uh.CreateUser)
	router.POST("/api/v1/tokens", uh.CreateToken)

	log.Fatal(http.ListenAndServe(":8080", router))
}
