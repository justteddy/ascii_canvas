package handler

import (
	"net/http"

	store "canvas/storage"

	"github.com/go-chi/chi"
)

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	canvasID := chi.URLParam(r, "canvasID")
	data, err := h.storage.Get(canvasID)
	if err != nil {
		if err == store.ErrNotFound {
			writeErrorResponse(ctx, w, http.StatusNotFound, err)
			return
		}
		writeErrorResponse(ctx, w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
