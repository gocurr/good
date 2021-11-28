package db

import (
	"database/sql"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
	"strings"
)

var Db *sql.DB

// Init inits Db via driver
func Init(c *conf.Configuration) error {
	dbConf := c.DB
	pw, err := crypto.Decrypt(c.Secure.Key, dbConf.Password)
	if err != nil {
		return err
	}

	ds := `user="` + dbConf.User + `" password="` + pw + `" connectString="` + dbConf.Datasource + `"`
	switch strings.ToLower(dbConf.Driver) {
	case "mysql":
		Db, err = openMysql(ds)
	case "godror":
		Db, err = openOracle(ds)
	}

	if err != nil {
		return err
	}
	return Db.Ping()
}
