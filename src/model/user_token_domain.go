package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
)

var JWT_SECRET_KEY = "JWT_SECRET_KEY"

func (ud *userDomain) GenerateToken() (string, *errs.Errs) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id": ud.id,
		"email": ud.email,
		"name": ud.name,
		"age": ud.age,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errs.NewInternalServerErrs(fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}


func VerifyToken(tokenValue string) (UserDomainInterface, *errs.Errs) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errs.NewBadRequestErrs("invalid token")
	})

	if err != nil {
		return nil, errs.NewUnauthorizedRequestErrs("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errs.NewUnauthorizedRequestErrs("invalid token")
	}

	return &userDomain{
		id: claims["id"].(string),
		email: claims["email"].(string),
		name: claims["name"].(string),
		age: int8(claims["age"].(float64)),
	}, nil
}


func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix("Bearer", token)
	}

	return token
}