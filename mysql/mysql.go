package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"reflect"
)

const mysql = "mysql"

var mysqlErr = errors.New("bad mysql configuration")

// Open returns *sql.DB and error
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		panic(mysqlErr)
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

	mysqlField := c.FieldByName(consts.Mysql)
	if !mysqlField.IsValid() {
		panic(mysqlErr)
	}

	userField := mysqlField.FieldByName(consts.User)
	if !userField.IsValid() {
		panic(mysqlErr)
	}
	user := userField.String()

	passwordField := mysqlField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		panic(mysqlErr)
	}
	password := passwordField.String()

	datasourceField := mysqlField.FieldByName(consts.Datasource)
	if !datasourceField.IsValid() {
		panic(mysqlErr)
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
	return sql.Open(mysql, dsn)
}
