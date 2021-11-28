package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func openMysql(ds string) (*sql.DB, error) {
	return sql.Open("mysql", ds)
}
