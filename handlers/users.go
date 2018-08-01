package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/droptheplot/flashcards/auth"
	"github.com/julienschmidt/httprouter"
)

type CreateUserParams struct {
	Email    string `valid:"required,email"`
	Password string `valid:"required"`
}

type CreateTokenParams struct {
	Email    string `valid:"required,email"`
	Password string `valid:"required"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signUpUserParams := CreateUserParams{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	_, err = govalidator.ValidateStruct(signUpUserParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signUpUserParams.Password = auth.HashPassword(signUpUserParams.Password)

	err = h.DBRepository.CreateUser(signUpUserParams.Email, signUpUserParams.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) CreateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createTokenParams := CreateTokenParams{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	_, err = govalidator.ValidateStruct(createTokenParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.DBRepository.GetUserByEmail(createTokenParams.Email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if user.Password != auth.HashPassword(createTokenParams.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateToken()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.KVRepository.CreateToken(token, user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(CreateTokenResponse{Token: token})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
