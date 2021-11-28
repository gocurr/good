package good

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
)

var Db *sql.DB

func initDb() error {
	dbConf := conf.DB
	if dbConf == nil {
		return nil
	}

	var err error
	ds := `user="` + dbConf.User + `" password="` + dbConf.Password + `" connectString="` + dbConf.Datasource + `"`
	Db, err = sql.Open(dbConf.Driver, ds)
	if err != nil {
		return err
	}
	return Db.Ping()
}
