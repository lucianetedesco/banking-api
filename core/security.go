package core

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/lucianetedesco/banking-api/settings"
	"strconv"
	"time"
)

type Claims struct {
	AccountID uint `json:"account_id"`
	jwt.StandardClaims
}

func GenerateToken(accountID uint) (string, error) {
	claims := &Claims{
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(settings.Environment.Auth.ExpiresAt * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(settings.Environment.Auth.SecretKey))
	return tokenString, err
}

func GetAccountID(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(settings.Environment.Auth.SecretKey), nil
		},
	)
	if err != nil {
		return " ", errors.New("error validating token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return " ", errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return " ", errors.New("token expired")
	}

	return strconv.Itoa(int(claims.AccountID)), nil
}
