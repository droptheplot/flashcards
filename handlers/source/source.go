package source

import (
	"io"
	"net/http"
	"strconv"

	"github.com/droptheplot/flashcards/entities"
	"github.com/julienschmidt/httprouter"
)

type UseCase interface {
	GetSources() ([]entities.Source, error)
	GetSourceByID(ID int) (entities.Source, error)
}

type Render interface {
	JSON(w io.Writer, status int, v interface{}) error
}

type Handler struct {
	UseCase UseCase
	Render  Render
}

func (h *Handler) GetSources(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sources, _ := h.UseCase.GetSources()

	h.Render.JSON(w, http.StatusOK, sources)
}

func (h *Handler) GetSourceByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ID, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		http.Error(w, "`id` should be a number", http.StatusBadRequest)
		return
	}

	source, err := h.UseCase.GetSourceByID(ID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Render.JSON(w, http.StatusOK, source)
}
