package exception

import (
	"encoding/json"
	"net/http"

	"github.com/Billy278/bitly-sederhana/helper"
	"github.com/Billy278/bitly-sederhana/model/web"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	if validateError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}
func validateError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exeption, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Data Bitly Not registration",
			Data:   exeption.Error(),
		}
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		er := encoder.Encode(&webResponse)
		helper.PanicIfError(er)
		return true
	} else {
		return false
	}
}
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFound)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Data Bitly Not registration",
			Data:   exception.Error,
		}
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		er := encoder.Encode(&webResponse)
		helper.PanicIfError(er)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	er := encoder.Encode(&webResponse)
	helper.PanicIfError(er)

}
