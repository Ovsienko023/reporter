package http

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
)

var ErrUnauthorized = errors.New("unauthorized error")

func authorize(token string) (string, error) {
	if token == "" {
		return "", ErrUnauthorized
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil // todo secret into config
	})
	if err != nil {
		return "", ErrUnauthorized
	}

	if _, ok := data.Method.(*jwt.SigningMethodHMAC); !ok {
		//return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		fmt.Println(ok)
	}

	if claims, ok := data.Claims.(jwt.MapClaims); ok && data.Valid {
		fmt.Println(claims["exp"], claims["grant_type"], claims["user_id"])

		return claims["iss"].(string), nil
	} else {
		//errorContainer.Done(w, http.StatusUnauthorized, err.Error()) // todo check error
		return "", ErrUnauthorized
	}
}
