package repository

import (
	"context"
	"database/sql"

	"github.com/Billy278/bitly-sederhana/model/domain"
)

type BitlyRepository interface {
	Create(ctx context.Context, DB *sql.DB, bitly domain.Bitly) domain.Bitly
	FindById(ctx context.Context, DB *sql.DB, link string) (domain.Bitly, error)
}
