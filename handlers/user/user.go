package user

import (
	"io"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
)

type UseCase interface {
	CreateUser(email string, password string) error
	CreateToken(email string, password string) (string, error)
}

type Render interface {
	JSON(w io.Writer, status int, v interface{}) error
}

type Handler struct {
	UseCase UseCase
	Render  Render
}

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

	err = h.UseCase.CreateUser(signUpUserParams.Email, signUpUserParams.Password)

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

	token, err := h.UseCase.CreateToken(createTokenParams.Email, createTokenParams.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.Render.JSON(w, http.StatusOK, CreateTokenResponse{Token: token})
}
