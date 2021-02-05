package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) HandleFloodFill(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	canvasID := chi.URLParam(r, "canvasID")

	var p FloodFillParams
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeErrorResponse(ctx, w, http.StatusBadRequest, err)
		return
	}

	canvas, err := h.restoreCanvas(ctx, canvasID)
	if err != nil {
		writeErrorResponse(ctx, w, http.StatusInternalServerError, err)
		return
	}

	if err := canvas.FloodFill(p.X, p.Y, p.Fill); err != nil {
		writeErrorResponse(ctx, w, http.StatusBadRequest, err)
		return
	}

	data, err := h.storeCanvas(canvasID, canvas)
	if err != nil {
		writeErrorResponse(ctx, w, http.StatusInternalServerError, err)
		return
	}

	writeSuccessResponse(ctx, w, data)
}
