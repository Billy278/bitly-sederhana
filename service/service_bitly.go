package service

import (
	"context"

	"github.com/Billy278/bitly-sederhana/model/web"
)

type ServiceBitly interface {
	Create(ctx context.Context, request web.CreateRequest) web.ResponseBitly
	FindById(ctx context.Context, link string) web.ResponseBitly
}
