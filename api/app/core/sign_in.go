package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

func (c *Core) SignIn(ctx context.Context, msg *domain.SignInRequest) (*domain.SignInResponse, error) {
	auth, err := c.db.SignIn(ctx, msg.ToDbSignIn())
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrCredentials):
			return nil, ErrCredentials
		}
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	err = c.checkPassword(*auth.Password, msg.Password)
	if err != nil {
		return nil, ErrCredentials
	}

	mySigningKey := []byte("SecretKey") // todo add secret from config

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(740 * time.Hour).Unix(),
		Issuer:    *auth.UserId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(mySigningKey)

	return &domain.SignInResponse{
		Token: &signedToken,
		// todo ... add ExpiresIn
	}, nil
}
