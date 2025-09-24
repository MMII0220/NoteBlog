package service

import (
	// "errors"
	"golang.org/x/crypto/bcrypt"
	// "myasd/internal/repository"
	// "myasd/internal/controller"
	"myasd/internal/errs"
	"myasd/internal/models"
)

func (s *ServiceStruct) CreateUser(user models.User) error {
	// Проверяем, существует ли уже пользователь с таким логином
	_, err := s.repo.GetUserByLogin(user.Login)
	if err == nil {
		// Если пользователь найден, возвращаем ошибку
		return errs.ErrUsernameAlreadyExists
	}

	if user.ID != 0 {
		return errs.ErrUsernameAlreadyExists
	}

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
		return models.TokenResponse{}, errs.ErrUserNotExists
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.TokenResponse{}, errs.ErrIncorrectLoginOrPassword
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
		return "", errs.ErrIncorrectRefreshToken
	}

	access, _, err := s.GenerateTokens(userID)
	if err != nil {
		return "", err
	}
	return access, err
}
