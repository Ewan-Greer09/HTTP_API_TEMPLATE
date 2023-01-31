package handlers

import (
	"net/http"
)

func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
