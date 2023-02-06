package app

import (
	"database/sql"
	"time"

	"github.com/Billy278/bitly-sederhana/helper"
)

func NewDB() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bitly_sederhana")
	helper.PanicIfError(err)
	//minimal coneksi 10
	DB.SetMaxIdleConns(10)
	//maksiml koneksi 20
	DB.SetConnMaxLifetime(20)
	//setiap 10 menit sekali apabila koneksi tidak digunakan akan
	// ditutup dan diset ke jumlah koneksi minimal yg diijinkan
	DB.SetConnMaxIdleTime(10 * time.Minute)

	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
