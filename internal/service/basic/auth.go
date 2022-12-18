package basic

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yudgxe/sima-rest-api/internal/service"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type Auth struct {
	store      store.Store
	signingKey string
	tokenTTL   time.Duration
}

var _ service.AuthService = (*Auth)(nil)

func NewAuthService(s store.Store, signingKey string, tokenTTL time.Duration) *Auth {
	return &Auth{
		store:      s,
		signingKey: signingKey,
		tokenTTL:   tokenTTL,
	}
}

func (a *Auth) GenerateToken(login, password string) (string, error) {
	p, err := a.store.Auth().GetUser(login, password)
	if err != nil {
		return "", err
	}

	type tokenClaims struct {
		UserId     int    `json:"user_id"`
		Permission string `json:"permission"`
		jwt.StandardClaims
	}

	claims := &tokenClaims{
		p.UserID,
		p.Permission,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(a.tokenTTL).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(a.signingKey))
}
