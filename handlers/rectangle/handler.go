package rectangle

import "net/http"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
