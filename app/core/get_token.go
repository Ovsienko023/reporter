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

	response := &domain.GetTokenResponse{
		Token: &ss,
	}

	return response, nil
}
