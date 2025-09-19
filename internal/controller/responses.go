package controller

import ()

type CommonError struct {
	Error string `json:"error"`
}

type CommonReponse struct {
	Message string `json:"message"`
}
