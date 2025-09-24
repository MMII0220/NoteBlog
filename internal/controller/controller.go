package controller

import (
	"errors"
	"myasd/internal/contracts"
	"myasd/internal/errs"
	"myasd/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type ControllerStruct struct {
	serv   contracts.ServiceI
	logger zerolog.Logger
}

func NewController(s *service.ServiceStruct, logger zerolog.Logger) *ControllerStruct {
	return &ControllerStruct{
		serv:   s,
		logger: logger,
	}
}

func (ctrl *ControllerStruct) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrNotFound) || errors.Is(err, errs.ErrArticleNotFound) || errors.Is(err, errs.ErrUserNotExists):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidRequestBody) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists) ||
		errors.Is(err, errs.ErrInvalidPathParam):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrFillRequiredFields):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectLoginOrPassword) || errors.Is(err, errs.ErrUserIDNotFoundInContext):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
