package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ubaidillahhf/go-clarch/app/domain"
)

type JwtCustomClaims struct {
	Fullname string `json:"fullname"`
	ID       string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type tokenGenerator struct {
	secret string
}

func NewTokenGenerator(secret string) domain.TokenGenerator {
	return &tokenGenerator{
		secret: secret,
	}
}

func (t *tokenGenerator) GenerateAccessToken(user *domain.User, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	nD := jwt.NewNumericDate(exp)
	claims := &JwtCustomClaims{
		Fullname: user.Fullname,
		ID:       user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: nD,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.secret))
}

func (t *tokenGenerator) GenerateRefreshToken(user *domain.User, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	nD := jwt.NewNumericDate(exp)
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: nD,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	return token.SignedString([]byte(t.secret))
}
