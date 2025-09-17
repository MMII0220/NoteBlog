package service

import (
	"myasd/internal/repository"
)

type ServiceStruct struct {
	repo *repository.RepositoryStruct
}

func NewService(r *repository.RepositoryStruct) *ServiceStruct {
	return &ServiceStruct{
		repo: r,
	}
}
