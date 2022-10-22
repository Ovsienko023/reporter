package core

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/golang-jwt/jwt"
	"strings"
)

type Core struct {
	db database.InterfaceDatabase
}

func NewCore(db database.InterfaceDatabase) *Core {
	return &Core{
		db: db,
	}
}

// authorize возвращает InvokerId или ошибку:
// ErrUnauthorized
func (c *Core) authorize(token string) (string, error) {
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
		return "", ErrUnauthorized // signing method: token.Header["alg"]
	}

	if claims, ok := data.Claims.(jwt.MapClaims); ok && data.Valid {
		return claims["iss"].(string), nil
	} else {
		return "", ErrUnauthorized
	}
}
