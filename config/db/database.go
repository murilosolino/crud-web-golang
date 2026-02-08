package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@/loja_go")
	if err != nil {
		panic(err)
	}
	return db
}
