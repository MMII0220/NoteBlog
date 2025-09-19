package controller

import (
	"myasd/internal/contracts"
	"myasd/internal/service"
)

type ControllerStruct struct {
	serv contracts.ServiceI
}

func NewController(s *service.ServiceStruct) *ControllerStruct {
	return &ControllerStruct{
		serv: s,
	}
}
