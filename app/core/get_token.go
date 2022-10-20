package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/golang-jwt/jwt"
	"time"
)

func (c *Core) GetToken(ctx context.Context, msg *domain.GetTokenRequest) (*domain.GetTokenResponse, error) {
	message := database.GetAuthUser{
		Login: msg.Login,
	}

	auth, err := c.db.GetAuthUser(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrCredentials):
			return nil, ErrCredentials
		}
		fmt.Println("LOG: ", err) // todo add logger
		return nil, ErrInternal
	}

	if msg.Password != *auth.Password {
		return nil, ErrCredentials
	}
	mySigningKey := []byte("SecretKey")

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(100 * time.Minute).Unix(),
		Issuer:    *auth.UserId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256,
	//	jwt.MapClaims{
	//		"exp":        time.Now().Add(43_800 * time.Minute),
	//		"grant_type": "password",
	//		"user_id":    auth.UserId,
	//	},
	//)
	//
	//tokenString, err := token.SignedString([]byte("SecretKey"))
	//if err != nil {
	//	return nil, err // todo
	//}

	response := &domain.GetTokenResponse{
		Token: &ss,
	}

	return response, nil
}

//func generateJWT() (string, error) {
//	token := jwt.New(jwt.SigningMethodRS256)
//
//	claims := token.Claims.(jwt.MapClaims)
//	claims["exp"] = time.Now().Add(43_800 * time.Minute)
//	claims["grant_type"] = "password"
//	claims["user_id"] =
//}
