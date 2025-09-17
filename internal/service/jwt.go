package service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var accessSecret = []byte("access_secret")
var refreshSecret = []byte("refresh_secret")

// GenerateTokens генерирует Access и Refresh токены
func GenerateTokens(userID int) (accessToken, refreshToken string, err error) {
	// Access Token (живет 1 минуту)
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(accessSecret)
	if err != nil {
		return
	}

	// Refresh Token (живет 7 дней)
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(refreshSecret)
	return
}

// ValidateAccessToken проверяет access token
func ValidateAccessToken(tokenStr string) (userID int, err error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID = int(claims["user_id"].(float64))
	return
}

// ValidateRefreshToken проверяет refresh token
func ValidateRefreshToken(tokenStr string) (userID int, err error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID = int(claims["user_id"].(float64))
	return
}
