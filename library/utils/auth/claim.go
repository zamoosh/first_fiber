package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	day = time.Hour * 24
)

type TokenType string

const (
	RefreshToken TokenType = "refresh"
	AccessToken  TokenType = "access"
)

// JwtClaim, extends the [jwt.Claims] and adds UserId and TokenType to it.
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

func ToJwtClaim(claim jwt.Claims) (*JwtClaim, error) {
	data, ok := claim.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("توکن معتبر نیست")
	}

	if _, ok = data["user_id"]; !ok {
		return nil, errors.New("توکن معتبر نیست")
	}
	userId := uint64(data["user_id"].(float64))

	if _, ok = data["token_type"]; !ok {
		return nil, errors.New("توکن معتبر نیست")
	}
	tt := TokenType(data["token_type"].(string))

	var jti string
	if jti, ok = data["jti"].(string); !ok || len(jti) == 0 {
		return nil, errors.New("توکن معتبر نیست")
	}

	exp, _ := claim.GetExpirationTime()
	iat, _ := claim.GetIssuedAt()

	return &JwtClaim{
		TokenType: tt,
		Exp:       uint64(exp.Time.Unix()),
		Iat:       uint64(iat.Time.Unix()),
		Jti:       jti,
		UserId:    userId,
	}, nil
}
