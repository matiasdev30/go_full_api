package service

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matiasdev30/go_api/models"
)

func GenereteToken(user * models.User) (string, error){

	claims := jwt.MapClaims{
		"id" : user.ID,
		"email" : user.Email,
		"name" : user.Name,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (bool, error) {
	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte("secret-key"), nil
		}

		return false, errors.New("invalid token")
	})

	if err != nil {
		return false, errors.New("invalid token")
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	return ok, nil

}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer "){
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}