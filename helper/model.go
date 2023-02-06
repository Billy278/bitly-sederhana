package helper

import (
	"github.com/Billy278/bitly-sederhana/model/domain"
	"github.com/Billy278/bitly-sederhana/model/web"
)

func ToResponseBitly(bitly domain.Bitly) web.ResponseBitly {
	return web.ResponseBitly{
		Id:        bitly.Id,
		LongLink:  bitly.LongLink,
		ShortLink: bitly.ShortLink,
	}
}
