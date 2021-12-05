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

const postgres = "postgres"

var postgresErr = errors.New("bad postgres configuration")

// Open returns *sql.DB and error
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		panic(postgresErr)
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
		panic(postgresErr)
	}

	hostField := postgresField.FieldByName(vars.Host)
	if !hostField.IsValid() {
		panic(postgresErr)
	}
	host := hostField.String()

	portField := postgresField.FieldByName(vars.Port)
	if !portField.IsValid() {
		panic(postgresErr)
	}
	port := portField.Int()

	userField := postgresField.FieldByName(vars.User)
	if !userField.IsValid() {
		panic(postgresErr)
	}
	user := userField.String()

	passwordField := postgresField.FieldByName(vars.Password)
	if !passwordField.IsValid() {
		panic(postgresErr)
	}
	password := passwordField.String()

	dbField := postgresField.FieldByName(vars.DB)
	if !dbField.IsValid() {
		panic(postgresErr)
	}
	db := dbField.String()

	sslMode := "disable"
	sslField := postgresField.FieldByName(vars.SSLMode)
	if sslField.IsValid() {
		sslMode = sslField.String()
	}

	var err error
	if key != "" {
		password, err = crypto.Decrypt(key, password)
		if err != nil {
			return nil, err
		}
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%v",
		host, port, user, password, db, sslMode)
	rdb, err := sql.Open(postgres, dsn)
	if err != nil {
		return nil, err
	}
	return rdb, rdb.Ping()
}
