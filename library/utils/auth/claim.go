package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	day = time.Hour * 24
)

type TokenType string
type CustomClaim interface {
	jwt.Claims
	GetUserId() uint64
}

const (
	RefreshToken TokenType = "refresh"
	AccessToken  TokenType = "access"
)

type JwtClaim struct {
	TokenType TokenType `json:"token_type"`
	Exp       uint64    `json:"exp"`
	Iat       uint64    `json:"iat"`
	Jti       string    `json:"jti"`
	UserId    uint64    `json:"user_id"`
}

func (j JwtClaim) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Exp), 0)}, nil
}

func (j JwtClaim) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Iat), 0)}, nil
}

func (j JwtClaim) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(int64(j.Iat), 0)}, nil
}

func (j JwtClaim) GetIssuer() (string, error) {
	return "", nil
}

func (j JwtClaim) GetSubject() (string, error) {
	return "", nil
}

func (j JwtClaim) GetAudience() (jwt.ClaimStrings, error) {
	return []string{}, nil
}

func (j JwtClaim) GetUserId() uint64 {
	return j.UserId
}
