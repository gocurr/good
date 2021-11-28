package db

import (
	"database/sql"
	"fmt"
	"github.com/gocurr/good/conf"
	_ "github.com/godror/godror"
)

// openOracle open oracle
func openOracle(c *conf.Configuration, pw string) (*sql.DB, error) {
	dbConf := c.DB
	ds := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`, dbConf.User, pw, dbConf.Datasource)
	return sql.Open("godror", ds)
}
