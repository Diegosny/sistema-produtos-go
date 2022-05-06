package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "diego:root@123@tcp(localhost)/banco_teste")

	if err != nil {
		panic(err.Error())
	}
	return db
}
