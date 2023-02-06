package repository

import (
	"context"
	"database/sql"

	"errors"

	"github.com/Billy278/bitly-sederhana/helper"
	"github.com/Billy278/bitly-sederhana/model/domain"
)

type BitlyRepositoryImpl struct {
}

func NewBitlyRepositoryImpl() BitlyRepository {
	return &BitlyRepositoryImpl{}
}

func (bitly_repository *BitlyRepositoryImpl) Create(ctx context.Context, DB *sql.DB, bitly domain.Bitly) domain.Bitly {
	sql := " SELECT COUNT(*) as jumlah_baris FROM bitly"
	sql_create := "Insert Into bitly(long_link,short_link) Values(?,?)"
	rows, err := DB.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&bitly.ShortLink)
		data := "bitly" + bitly.ShortLink
		//fmt.Println("cek lokasi")
		bitly.ShortLink = data
		result, err := DB.ExecContext(ctx, sql_create, bitly.LongLink, bitly.ShortLink)
		helper.PanicIfError(err)
		id, err := result.LastInsertId()
		helper.PanicIfError(err)
		bitly.Id = int(id)
		return bitly
		// } else {
		// 	return bitly
		// }
	}
	return domain.Bitly{}
}

func (bitly_repository *BitlyRepositoryImpl) FindById(ctx context.Context, DB *sql.DB, link string) (domain.Bitly, error) {
	sql := "Select id,long_link,short_link from bitly where short_link=?"
	rows, err := DB.QueryContext(ctx, sql, link)
	helper.PanicIfError(err)
	bitly := domain.Bitly{}
	if rows.Next() {
		err := rows.Scan(&bitly.Id, &bitly.LongLink, &bitly.ShortLink)
		helper.PanicIfError(err)
		return bitly, nil
	} else {
		return bitly, errors.New("SHORT LINK NOT REGISTRATION")
	}
}
