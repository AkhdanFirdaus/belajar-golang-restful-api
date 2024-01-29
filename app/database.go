package app

import (
	"database/sql"
	"time"

	"github.com/AkhdanFirdaus/belajar-golang-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Passw0rd@!@tcp(localhost:3306)/belajar_golang_restfulapi")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
