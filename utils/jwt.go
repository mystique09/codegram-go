package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type (
	UserPayload struct {
		Id       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Emai     string    `json:"email"`
	}

	JwtClaims struct {
		Payload UserPayload
		jwt.StandardClaims
	}
)

func CreateJwt(payload UserPayload) (string, error) {
	claims := JwtClaims{
		payload,
		jwt.StandardClaims{
			Id:        payload.Id.String(),
			ExpiresAt: time.Now().Add(24 * time.Minute).Unix(),
		},
	}

	raw_token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := raw_token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
