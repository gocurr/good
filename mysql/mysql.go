package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/pre"
	"reflect"
)

const mysql = "mysql"

var err = errors.New("bad mysql configuration")

// Open returns a mysql DB and reports error
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		return nil, err
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

	mysqlField := c.FieldByName(pre.Mysql)
	if !mysqlField.IsValid() {
		return nil, err
	}

	userField := mysqlField.FieldByName(consts.User)
	if !userField.IsValid() {
		return nil, err
	}
	user := userField.String()

	passwordField := mysqlField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		return nil, err
	}
	password := passwordField.String()

	datasourceField := mysqlField.FieldByName(consts.Datasource)
	if !datasourceField.IsValid() {
		return nil, err
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
