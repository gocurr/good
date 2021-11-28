package db

import (
	"database/sql"
	_ "github.com/godror/godror"
)

// openOracle open oracle
func openOracle(ds string) (*sql.DB, error) {
	return sql.Open("godror", ds)
}
