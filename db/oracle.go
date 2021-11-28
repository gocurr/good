package db

import (
	"database/sql"
	_ "github.com/godror/godror"
)

func openOracle(ds string) (*sql.DB, error) {
	return sql.Open("godror", ds)
}
