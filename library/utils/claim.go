package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	RefreshToken TokenType = "refresh"
	AccessToken  TokenType = "access"
)

type jwtClaim struct {
	TokenType TokenType `json:"token_type"`
	Exp       uint64 `json:"exp"`
	Iat       uint64 `json:"iat"`
	Jti       string `json:"jti"`
	UserId    uint64 `json:"user_id"`
	Type      string `json:"type"`
}

func (j jwtClaim) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Exp), 0)}, nil
}

func (j jwtClaim) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Iat), 0)}, nil
}

func (j jwtClaim) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Iat), 0)}, nil
}

func (j jwtClaim) GetIssuer() (string, error) {
	return "", nil
}

func (j jwtClaim) GetSubject() (string, error) {
	return "", nil
}

func (j jwtClaim) GetAudience() (jwt.ClaimStrings, error) {
	return []string{}, nil
}

const (
	day = time.Hour * 24
)
