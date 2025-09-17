package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	// "myasd/internal/repository"
	// "myasd/internal/controller"
	"myasd/internal/models"
)

func (s *ServiceStruct) CreateUser(user models.User) error {
	// хешируем пароль
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.CreateUser(user)
}

func (s *ServiceStruct) GetUser(login, password string) (models.TokenResponse, error) {
	user, err := s.repo.GetUserByLogin(login)
	if err != nil {
		return models.TokenResponse{}, errors.New("invalid login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.TokenResponse{}, errors.New("invalid password")
	}

	access, refresh, err := s.GenerateTokens(user.ID)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return models.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *ServiceStruct) RefreshToken(refreshToken string) (string, error) {
	userID, err := s.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	access, _, err := s.GenerateTokens(userID)
	if err != nil {
		return "", err
	}
	return access, err
}
