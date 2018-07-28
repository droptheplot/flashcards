package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetSources(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sources, _ := h.Repository.GetSources()

	js, err := json.Marshal(sources)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h *Handler) GetSourceByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ID, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		http.Error(w, "`id` should be a number", http.StatusBadRequest)
		return
	}

	source, err := h.Repository.GetSourceByID(ID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cards, _ := h.Repository.GetCardsBySourceID(ID)

	source.Cards = cards

	js, err := json.Marshal(source)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
