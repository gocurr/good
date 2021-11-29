package db

import (
	"database/sql"
	"errors"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
	"strings"
)

const (
	mysql  = "mysql"
	godror = "godror"
)

var dbErr = errors.New("bad db configuration")

// DB the global database object
var DB *sql.DB

// Init inits DB
func Init(c *conf.Configuration) error {
	if c == nil {
		return dbErr
	}
	db := c.DB
	if db == nil {
		return dbErr
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

	switch strings.ToLower(db.Driver) {
	default:
		return errors.New("unsupported database")
	case mysql:
		DB, err = openMysql(c, pw)
	case godror:
		DB, err = openOracle(c, pw)
	}

	if err != nil {
		return err
	}
	return DB.Ping()
}
