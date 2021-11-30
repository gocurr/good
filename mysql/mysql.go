package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

const mysql = "mysql"

// DB the global database object
var DB *sql.DB

var mysqlErr = errors.New("bad mysql configuration")

// Init opens Oracle
func Init(c *conf.Configuration) error {
	if c == nil {
		return mysqlErr
	}
	db := c.Mysql
	if db == nil {
		return mysqlErr
	}

	var err error
	var pw string
	if c.Secure == nil || c.Secure.Key == "" {
		pw = db.Password
	} else {
		pw, err = crypto.Decrypt(c.Secure.Key, db.Password)
		if err != nil {
			return err
		}
	}

	ds := fmt.Sprintf("%s:%s@%s", db.User, pw, db.Datasource)
	DB, err = sql.Open(mysql, ds)
	return err

}
