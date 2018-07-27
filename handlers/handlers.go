package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Repo interface {
	Ping() string
}

type Handler struct {
	Repo Repo
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := h.Repo.Ping()

	fmt.Fprint(w, p)
}
