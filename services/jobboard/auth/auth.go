package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
)

// TODO: add func to generate a new API Key when a user registers
// TODO: move secretKey and APIKeys to github secrets for security

var secretkey = []byte("SuperSecrtetKey")
var APIKeys = []string{"1234567890"}

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Routes returns a http.Handler that handles all the routes for the auth router
func (h *AuthHandler) Routes() http.Handler {
	router := chi.NewRouter()
	router.Post("/minttoken", h.HandleMintToken)
	return router
}

// generateJWT generates a JWT token and returns it as a string
func (h *AuthHandler) generateJWT(key string) (string, error) {
	err := checkAPIKey(key)
	if err != nil {
		return "", errors.New("Invalid API Key")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 24 * 7 * 52).Unix()
	claims["authorized"] = true

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token and returns a 401 if it is invalid
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

func checkAPIKey(key string) error {
	for _, k := range APIKeys {
		if k == key {
			return nil
		}
	}
	return errors.New("Invalid API Key")
}
