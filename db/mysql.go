package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocurr/good/conf"
)

// openMysql open mysql
func openMysql(c *conf.Configuration, pw string) (*sql.DB, error) {
	dbConf := c.DB
	ds := fmt.Sprintf("%s:%s@%s", dbConf.User, pw, dbConf.Datasource)
	return sql.Open("mysql", ds)
}
