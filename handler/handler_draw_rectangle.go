package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) HandleDrawRectangle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	canvasID := chi.URLParam(r, "canvasID")

	var p RectangleParams
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeErrorResponse(ctx, w, http.StatusBadRequest, err)
		return
	}

	canvas, err := h.restoreCanvas(ctx, canvasID)
	if err != nil {
		writeErrorResponse(ctx, w, http.StatusInternalServerError, err)
		return
	}

	if err := canvas.DrawRectangle(p.X, p.Y, p.Width, p.Height, p.Fill, p.Outline); err != nil {
		writeErrorResponse(ctx, w, http.StatusBadRequest, err)
		return
	}

	data, err := h.storeCanvas(canvasID, canvas)
	if err != nil {
		writeErrorResponse(ctx, w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data) // nolint
}
