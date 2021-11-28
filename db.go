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

	pw, err := decrypt(dbConf.Password)
	if err != nil {
		return err
	}
	ds := `user="` + dbConf.User + `" password="` + pw + `" connectString="` + dbConf.Datasource + `"`
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
