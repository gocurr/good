package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/env"
	"github.com/gocurr/good/pre"
)

const mysql = "mysql"

var errMysql = errors.New("mysql: bad mysql configuration")

// Open returns a mysql database and reports error encountered.
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		return nil, errMysql
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	var key string
	secureField := c.FieldByName(consts.Secure)
	if secureField.IsValid() {
		keyField := secureField.FieldByName(consts.Key)
		if keyField.IsValid() {
			key = keyField.String()
		}
	}
	if key == "" {
		key = env.GoodSecureKey()
	}

	mysqlField := c.FieldByName(pre.Mysql)
	if !mysqlField.IsValid() {
		return nil, errMysql
	}

	userField := mysqlField.FieldByName(consts.User)
	if !userField.IsValid() {
		return nil, errMysql
	}
	user := userField.String()

	passwordField := mysqlField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		return nil, errMysql
	}
	password := passwordField.String()

	datasourceField := mysqlField.FieldByName(consts.Datasource)
	if !datasourceField.IsValid() {
		return nil, errMysql
	}
	datasource := datasourceField.String()

	var err error
	if key != "" {
		password, err = crypto.Decrypt(key, password)
		if err != nil {
			return nil, err
		}
	}

	dsn := fmt.Sprintf("%s:%s@%s", user, password, datasource)
	db, err := sql.Open(mysql, dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
