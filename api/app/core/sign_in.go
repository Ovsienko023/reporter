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

func (c *Core) SignIn(ctx context.Context, msg *domain.SignInRequest) (*domain.SignInResponse, error) {
	auth, err := c.db.SignIn(ctx, msg.ToDbSignIn())
	if err != nil {
		switch {
		case errors.Is(err, database.ErrCredentials):
			return nil, ErrCredentials
		}
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if msg.Password != *auth.Password {
		return nil, ErrCredentials
	}
	mySigningKey := []byte("SecretKey") // todo add secret from config

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(100 * time.Minute).Unix(),
		Issuer:    *auth.UserId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(mySigningKey)

	return &domain.SignInResponse{
		Token: &signedToken,
		// todo ... add ExpiresIn
	}, nil
}
