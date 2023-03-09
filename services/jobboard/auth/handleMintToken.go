package auth

import (
	"log"
	"net/http"
)

// HandleMintToken handles the minting of a JWT token
func (h *AuthHandler) HandleMintToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.generateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
