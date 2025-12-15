package auth_gateway

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthToken{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
			Issuer:    "auth-gateway",
		},
	})
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string, secretKey []byte) (*AuthToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthToken{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err!= nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AuthToken)
	if!ok ||!token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func GenerateRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err!= nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

func ValidateRefreshToken(tokenString string) (string, error) {
	token, err := base64.StdEncoding.DecodeString(tokenString)
	if err!= nil {
		return "", err
	}
	if len(token)!= 32 {
		return "", http.StatusBadRequest
	}
	return string(token), nil
}

func LogRequest(r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.UserAgent())
}