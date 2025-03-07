package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateToken accepts userId and [TokenType]. Optionally you can give it expiration time in `exp` parameter.
func GenerateToken(userId uint, tokenType TokenType, exp ...time.Duration) (string, error) {
	if len(exp) == 0 {
		exp = append(exp, day)
	}

	switch tokenType {
	case RefreshToken, AccessToken:
	default:
		return "", fmt.Errorf("invalid token type: %s, choices are `%s` and `%s`", tokenType, AccessToken, RefreshToken)
	}

	claims := jwtClaim{
		TokenType: tokenType,
		Exp:       uint64(time.Now().UTC().Add(exp[0]).Unix()),
		Iat:       uint64(time.Now().UTC().Unix()),
		Jti:       strings.Replace(uuid.New().String(), "-", "", -1),
		UserId:    uint64(userId),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, claims,
	)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(t string) (bool, error) {
	token, err := jwt.Parse(
		t, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
