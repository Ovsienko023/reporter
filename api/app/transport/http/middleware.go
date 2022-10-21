package http

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func Authorization(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			errorContainer := ErrorResponse{}

			token := r.Header.Get("Authorization")
			if token == "" {
				return
			}

			token = strings.Replace(token, "Bearer ", "", 1)

			data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				return []byte("SecretKey"), nil // todo secret into config
			})
			if _, ok := data.Method.(*jwt.SigningMethodHMAC); !ok {
				//return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				fmt.Println(ok)
			}

			if claims, ok := data.Claims.(jwt.MapClaims); ok && data.Valid {
				fmt.Println(claims["exp"], claims["grant_type"], claims["user_id"])
			} else {
				errorContainer.Done(w, http.StatusUnauthorized, err.Error()) // todo check error
				return
			}

		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
