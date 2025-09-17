package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"myasd/internal/repository"
	// "myasd/internal/controller"
	"myasd/internal/models"
)

func CreateUser(user models.User) error {
	// хешируем пароль
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return repository.CreateUser(user)
}

func GetUser(login, password string) (models.TokenResponse, error) {
	user, err := repository.GetUserByLogin(login)
	if err != nil {
		return models.TokenResponse{}, errors.New("invalid login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.TokenResponse{}, errors.New("invalid password")
	}

	access, refresh, err := GenerateTokens(user.ID)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return models.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func RefreshToken(refreshToken string) (string, error) {
	userID, err := ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	access, _, err := GenerateTokens(userID)
	if err != nil {
		return "", err
	}
	return access, err
}
