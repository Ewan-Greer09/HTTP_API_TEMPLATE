package auth

import (
	"log"
	"net/http"
)

type JWTTokens struct {
	token  string
	scopes []string
}

func (h *AuthHandler) HandleMintToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.MintToken()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Minted token: " + token)

	w.Write([]byte(token))
}
