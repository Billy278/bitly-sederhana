package service

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

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

var cache Cache

type Cache struct {
	Data sync.Map
}
type CacheItem struct {
	Value    interface{}
	LongLife int64
}

func (c *Cache) Set(key interface{}, value interface{}, ttl time.Duration) {
	item := CacheItem{}
	item.Value = value
	item.LongLife = time.Now().Add(ttl).UnixNano()
	c.Data.Store(key, item)
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	item, ok := c.Data.Load(key)
	if ok {
		if item.(CacheItem).LongLife > time.Now().UnixNano() {
			return item.(CacheItem).Value, true
		}
		c.Data.Delete(key)
	}
	return nil, false
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

	result, ok := cache.Get(link)
	if ok {
		data := domain.Bitly{
			LongLink: result.(string),
		}
		fmt.Println("data ditemukan di cache")
		return helper.ToResponseBitly(data)
	}

	db := service_impl.DB
	bitly, err := service_impl.BitlyRepository.FindById(ctx, db, link)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	cache.Set(bitly.ShortLink, bitly.LongLink, 300*time.Second)
	fmt.Println("data tidak ditemukan di cache")
	return helper.ToResponseBitly(bitly)
}
