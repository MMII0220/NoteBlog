package service

import (
	"github.com/rs/zerolog"
	"myasd/internal/contracts"
	"myasd/internal/repository"
)

type ServiceStruct struct {
	repo   contracts.RepositoryI
	logger zerolog.Logger
}

func NewService(r *repository.RepositoryStruct, logger zerolog.Logger) *ServiceStruct {
	return &ServiceStruct{
		repo:   r,
		logger: logger,
	}
}
