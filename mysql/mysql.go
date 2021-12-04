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

var mysqlErr = errors.New("bad mysql configuration")

// Open returns *sql.DB and error
func Open(c *conf.Configuration) (*sql.DB, error) {
	if c == nil {
		return nil, mysqlErr
	}
	db := c.Mysql
	if db == nil {
		return nil, mysqlErr
	}

	var err error
	var pw string
	if c.Secure == nil || c.Secure.Key == "" {
		pw = db.Password
	} else {
		pw, err = crypto.Decrypt(c.Secure.Key, db.Password)
		if err != nil {
			return nil, err
		}
	}

	dsn := fmt.Sprintf("%s:%s@%s", db.User, pw, db.Datasource)
	return sql.Open(mysql, dsn)
}
