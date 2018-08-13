package main

import (
	"log"
	"net/http"
	"os"

	cardHandler "github.com/droptheplot/flashcards/handlers/card"
	sourceHandler "github.com/droptheplot/flashcards/handlers/source"
	userHandler "github.com/droptheplot/flashcards/handlers/user"
	"github.com/droptheplot/flashcards/middlewares"
	"github.com/droptheplot/flashcards/repositories/db"
	"github.com/droptheplot/flashcards/repositories/kv"
	cardUseCase "github.com/droptheplot/flashcards/usecases/card"
	sourceUseCase "github.com/droptheplot/flashcards/usecases/source"
	userUseCase "github.com/droptheplot/flashcards/usecases/user"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
)

func main() {
	pgc, err := sqlx.Open("postgres", os.Getenv("POSTGRES"))

	if err != nil {
		log.Fatal(err)
	}

	defer pgc.Close()

	rdc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	defer rdc.Close()

	dbr := db.Repository{Client: pgc}
	kvr := kv.Repository{Client: rdc}

	rd := render.New()

	uuc := userUseCase.UseCase{
		DBRepository: &dbr,
		KVRepository: &kvr,
		Secret:       os.Getenv("SECRET"),
	}

	uh := userHandler.Handler{
		UseCase: &uuc,
		Render:  rd,
	}

	sh := sourceHandler.Handler{
		UseCase: &sourceUseCase.UseCase{
			DBRepository: &dbr,
		},
		Render: rd,
	}

	ch := cardHandler.Handler{
		UseCase: &cardUseCase.UseCase{
			DBRepository: &dbr,
		},
		Render: rd,
	}

	router := httprouter.New()

	router.GET("/api/v1/sources", sh.GetSources)
	router.GET("/api/v1/sources/:id", sh.GetSourceByID)

	router.POST("/api/v1/cards/:id/:correct", middlewares.WithAuth(ch.AddCardToUser, &uuc))

	router.POST("/api/v1/users", uh.CreateUser)
	router.POST("/api/v1/tokens", uh.CreateToken)

	log.Fatal(http.ListenAndServe(":8080", middlewares.WithLogger(router)))
}
