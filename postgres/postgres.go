package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/vars"
	_ "github.com/lib/pq"
	"reflect"
)

const (
	postgres = "postgres"
	disable  = "disable"
)

var postgresErr = errors.New("bad postgres configuration")

// Open returns *sql.DB and error
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		return nil, postgresErr
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	var key string
	secureField := c.FieldByName(vars.Secure)
	if secureField.IsValid() {
		keyField := secureField.FieldByName(vars.Key)
		if keyField.IsValid() {
			key = keyField.String()
		}
	}

	postgresField := c.FieldByName(vars.Postgres)
	if !postgresField.IsValid() {
		return nil, postgresErr
	}

	hostField := postgresField.FieldByName(vars.Host)
	if !hostField.IsValid() {
		return nil, postgresErr
	}
	host := hostField.String()

	portField := postgresField.FieldByName(vars.Port)
	if !portField.IsValid() {
		return nil, postgresErr
	}
	port := portField.Int()

	userField := postgresField.FieldByName(vars.User)
	if !userField.IsValid() {
		return nil, postgresErr
	}
	user := userField.String()

	passwordField := postgresField.FieldByName(vars.Password)
	if !passwordField.IsValid() {
		return nil, postgresErr
	}
	password := passwordField.String()

	dbField := postgresField.FieldByName(vars.DB)
	if !dbField.IsValid() {
		return nil, postgresErr
	}
	db := dbField.String()

	sslMode := disable
	sslModeField := postgresField.FieldByName(vars.SSLMode)
	if sslModeField.IsValid() {
		mode := sslModeField.String()
		if mode != "" {
			sslMode = mode
		}
	}

	var err error
	if key != "" {
		password, err = crypto.Decrypt(key, password)
		if err != nil {
			return nil, err
		}
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, db, sslMode)
	rdb, err := sql.Open(postgres, dsn)
	if err != nil {
		return nil, err
	}
	return rdb, rdb.Ping()
}
