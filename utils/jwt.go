package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

type AuthClaims struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

func SignToken(claims AuthClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := jwtToken.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func VerifyToken(token string, checkExpiration ...bool) (*AuthClaims, error) {
	check, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})
	claims, ok := check.Claims.(*AuthClaims)
	if !ok || !check.Valid {
		return &AuthClaims{}, err
	}

	checkTime := len(checkExpiration) > 0 && checkExpiration[0]
	if checkTime && claims.ExpiresAt < time.Now().Unix() {
		return &AuthClaims{}, errors.New("token expired")
	}

	return claims, nil
}
