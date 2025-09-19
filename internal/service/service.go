package service

import (
	"myasd/internal/contracts"
	"myasd/internal/repository"
)

type ServiceStruct struct {
	repo contracts.RepositoryI
}

func NewService(r *repository.RepositoryStruct) *ServiceStruct {
	return &ServiceStruct{repo: r}
}
