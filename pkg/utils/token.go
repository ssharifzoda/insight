package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type CustomClaims struct {
	UserId      int `json:"username"`
	Permissions []int
	jwt.RegisteredClaims
}

func GenerateTokens(userId int, permissions []int) (string, string, error) {

	accessTokenSecret := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	refreshTokenSecret := []byte(os.Getenv("TOKEN_SECRET_KEY"))

	accessClaims := CustomClaims{
		UserId:      userId,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "insight",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(accessTokenSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "insight",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ParseRefreshToken(refreshToken string) (int, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}
