package controller

import (
	"myasd/internal/service"
)

type ControllerStruct struct {
	serv *service.ServiceStruct
}

func NewController(s *service.ServiceStruct) *ControllerStruct {
	return &ControllerStruct{
		serv: s,
	}
}
