package good

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
)

var db *sql.DB

// initDb inits db
func initDb() error {
	dbConf := conf.DB
	if dbConf == nil {
		return nil
	}

	var err error
	ds := `user="` + dbConf.User + `" password="` + dbConf.Password + `" connectString="` + dbConf.Datasource + `"`
	db, err = sql.Open(dbConf.Driver, ds)
	if err != nil {
		return err
	}
	return db.Ping()
}

// DB return db
func DB() *sql.DB {
	return db
}
