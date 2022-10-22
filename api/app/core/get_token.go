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
	auth, err := c.db.GetAuthUser(ctx, msg.ToDbGetToken())
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

	return &domain.GetTokenResponse{
		Token: &signedToken,
	}, nil
}
