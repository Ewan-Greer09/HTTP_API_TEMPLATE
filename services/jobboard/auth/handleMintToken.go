package auth

import (
	"log"
	"net/http"
)

// TODO: Clean up error handling
// HandleMintToken handles the minting of a JWT token
func (h *AuthHandler) HandleMintToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.generateJWT(r.Header.Get("APIKey"))
	if err != nil {
		if err.Error() == "Invalid API Key" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid API Key"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return

	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
