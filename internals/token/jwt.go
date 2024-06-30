package jwt

import (
	"time"

	"errors"

	"github.com/golang-jwt/jwt/v4"
	config "github.com/masterghost2002/videotube/configs"
)

var jwtKey = []byte(config.ENVS.JWTSecret)

type UserPayload struct {
	jwt.StandardClaims
	FullName string `json:"firstName"`
	Email    string `json:"email"`
}

func GenerateJWT(payload UserPayload) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	payload.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*UserPayload, error) {
	claims := &UserPayload{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// custom validation logic
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid signature")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil

}
