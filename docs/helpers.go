package authgateway

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(username string, password string, duration int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Second)),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}

func GeneratePasswordHash(password string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func ComparePasswordHash(password string, hash string) bool {
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}
	sha256hash := sha256.New()
	sha256hash.Write([]byte(password))
	return strings.Compare(hex.EncodeToString(sha256hash.Sum(nil)), hash) == 0
}

func IsEmailValid(email string) bool {
	return strings.Contains(email, "@") && strings.Count(email, "@") == 1
}

func IsPasswordValid(password string) bool {
	return len(password) >= 8 && !strings.ContainsAny(password, " ")
}

func ValidateUserInput(user map[string]string) error {
	if !IsEmailValid(user["email"]) {
		return errors.New("invalid email")
	}
	if !IsPasswordValid(user["password"]) {
		return errors.New("invalid password")
	}
	if user["username"] == "" {
		return errors.New("username is required")
	}
	if user["password"] == "" {
		return errors.New("password is required")
	}
	return nil
}

func HandleError(err error) {
	if err != nil {
		log.Println(err)
	}
}