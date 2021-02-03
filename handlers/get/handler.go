package get

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	canvasID := chi.URLParam(r, "canvasID")

	panic("asd")
	w.Write([]byte("welcome"))
}
