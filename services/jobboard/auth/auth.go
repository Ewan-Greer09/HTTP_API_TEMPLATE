package auth

import (
	"net/http"

	"github.com/go-chi/chi"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// route: /api/auth
func (h *AuthHandler) Routes() http.Handler {
	router := chi.NewRouter()
	router.Post("/minttoken", h.HandleMintToken)
	//router.Post("/verifytoken", h.HandleVerifyToken)
	return router
}

func (h *AuthHandler) HandleAuth() error {
	return nil
}

// MintToken creates a new JWT token
func (h *AuthHandler) MintToken() (string, error) {
	// TODO: Implement JWT token minting
	token := "testtoken"
	return token, nil
}

// VerifyToken verifies the JWT token
func (h *AuthHandler) VerifyToken() error {
	return nil
}
