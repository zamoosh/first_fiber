package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	TokenType string `json:"token_type"`
	Exp       uint64 `json:"exp"`
	Iat       uint64 `json:"iat"`
	Jti       string `json:"jti"`
	UserId    uint64 `json:"user_id"`
}

func (j jwtClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Exp), 0)}, nil
}

func (j jwtClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Iat), 0)}, nil
}

func (j jwtClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Iat), 0)}, nil
}

func (j jwtClaims) GetIssuer() (string, error) {
	return "", nil
}

func (j jwtClaims) GetSubject() (string, error) {
	return "", nil
}

func (j jwtClaims) GetAudience() (jwt.ClaimStrings, error) {
	return []string{}, nil
}

const (
	day = time.Hour * 24
)
