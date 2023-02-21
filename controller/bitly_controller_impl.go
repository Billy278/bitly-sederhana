package controller

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Billy278/bitly-sederhana/helper"
	"github.com/Billy278/bitly-sederhana/model/web"
	"github.com/Billy278/bitly-sederhana/service"
	"github.com/julienschmidt/httprouter"
)

type BitlyControllerImpl struct {
	BitlyService service.ServiceBitly
}

func ValidAndFixed(input string) string {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return "https://" + input
	} else {
		return input
	}
}
func NewBitlyControllerImpl(bitlyservice service.ServiceBitly) BitlyController {
	return &BitlyControllerImpl{
		BitlyService: bitlyservice,
	}
}
func (bitly_controller *BitlyControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	requestCreate := web.CreateRequest{}
	err := decoder.Decode(&requestCreate)
	helper.PanicIfError(err)
	requestCreate.LongLink = ValidAndFixed(requestCreate.LongLink)
	responseBitly := bitly_controller.BitlyService.Create(request.Context(), requestCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseBitly,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (bitly_controller *BitlyControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	shortLink := params.ByName("shortLink")

	responseBitly := bitly_controller.BitlyService.FindById(request.Context(), shortLink)

	// webResponse := web.WebResponse{
	// 	Code:   200,
	// 	Status: "Ok",
	// 	Data:   responseBitly.LongLink,
	// }

	// writer.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(writer)
	// err := encoder.Encode(&webResponse)
	// helper.PanicIfError(err)
	http.Redirect(writer, request, responseBitly.LongLink, http.StatusSeeOther)
}
