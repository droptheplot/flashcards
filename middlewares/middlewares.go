package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type authUseCase interface {
	AuthenticateUser(token string) (int, error)
}

func WithLogger(router http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)

		router.ServeHTTP(w, r)
	})
}

func WithAuth(handle httprouter.Handle, u authUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("Authorization")

		userID, err := u.AuthenticateUser(token)

		if err != nil {
			log.Printf("%d: %s", http.StatusUnauthorized, err.Error())

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		handle(w, r.WithContext(ctx), p)
	}
}
