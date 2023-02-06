package service

import (
	"context"
	"database/sql"

	"github.com/Billy278/bitly-sederhana/exception"
	"github.com/Billy278/bitly-sederhana/helper"
	"github.com/Billy278/bitly-sederhana/model/domain"
	"github.com/Billy278/bitly-sederhana/model/web"
	"github.com/Billy278/bitly-sederhana/repository"
	"github.com/go-playground/validator/v10"
)

type ServiceBitlyImpl struct {
	DB              *sql.DB
	BitlyRepository repository.BitlyRepository
	Validate        *validator.Validate
}

func NewServiceBitlyImpl(db *sql.DB, bitlyrepository repository.BitlyRepository, validate *validator.Validate) ServiceBitly {
	return &ServiceBitlyImpl{
		DB:              db,
		BitlyRepository: bitlyrepository,
		Validate:        validate,
	}
}

func (service_impl *ServiceBitlyImpl) Create(ctx context.Context, request web.CreateRequest) web.ResponseBitly {
	err := service_impl.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service_impl.DB
	bitly := domain.Bitly{
		LongLink: request.LongLink,
	}
	bitly = service_impl.BitlyRepository.Create(ctx, db, bitly)
	return helper.ToResponseBitly(bitly)
}
func (service_impl *ServiceBitlyImpl) FindById(ctx context.Context, link string) web.ResponseBitly {
	db := service_impl.DB
	bitly, err := service_impl.BitlyRepository.FindById(ctx, db, link)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	return helper.ToResponseBitly(bitly)
}
