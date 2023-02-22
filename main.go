package main

import (
	"net/http"

	"github.com/Billy278/bitly-sederhana/app"
	"github.com/Billy278/bitly-sederhana/controller"
	"github.com/Billy278/bitly-sederhana/exception"
	"github.com/Billy278/bitly-sederhana/helper"
	"github.com/Billy278/bitly-sederhana/repository"
	"github.com/Billy278/bitly-sederhana/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	repositoryBitly := repository.NewBitlyRepositoryImpl()
	serviceBitly := service.NewServiceBitlyImpl(db, repositoryBitly, validate)
	controllerBitly := controller.NewBitlyControllerImpl(serviceBitly)

	router := httprouter.New()
	router.GET("/bitly/:shortLink", controllerBitly.FindById)
	router.POST("/bitly", controllerBitly.Create)
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:7000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
