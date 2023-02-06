package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BitlyController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
