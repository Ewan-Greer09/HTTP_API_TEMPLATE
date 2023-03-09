package auth

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
)

var secretkey = []byte("SuperSecrtetKey")

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
	token, err := generateJWT()
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authorized"] = true

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token
func (h *AuthHandler) VerifyJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}

				return secretkey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	})
}
