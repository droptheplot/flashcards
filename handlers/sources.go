package handlers

import (
	"encoding/json"
	"net/http"

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
