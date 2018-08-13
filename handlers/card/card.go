package source

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UseCase interface {
	AddCardToUser(cardID, userID int, correct bool) error
}

type Render interface {
	JSON(w io.Writer, status int, v interface{}) error
}

type Handler struct {
	UseCase UseCase
	Render  Render
}

type AddCardToUserParams struct {
	CardID string `valid:"required,email"`
	UserID string `valid:"required"`
}

func (h *Handler) AddCardToUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cardID, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		http.Error(w, "`id` should be a number", http.StatusBadRequest)
		return
	}

	correct := params.ByName("correct") == "correct"

	userID := r.Context().Value("userID").(int)

	err = h.UseCase.AddCardToUser(cardID, userID, correct)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
