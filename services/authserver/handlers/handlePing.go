package handlers

import "net/http"

func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

//call with curl http://localhost:3000/api/ping
