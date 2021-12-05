package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/vars"
	_ "github.com/godror/godror"
	"reflect"
)

const godror = "godror"

var oracleErr = errors.New("bad oracle configuration")

// Open returns *sql.DB and error
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		panic(oracleErr)
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

	oracleField := c.FieldByName(vars.Oracle)
	if !oracleField.IsValid() {
		panic(oracleErr)
	}

	userField := oracleField.FieldByName(vars.User)
	if !userField.IsValid() {
		panic(oracleErr)
	}
	user := userField.String()

	passwordField := oracleField.FieldByName(vars.Password)
	if !passwordField.IsValid() {
		panic(oracleErr)
	}
	password := passwordField.String()

	datasourceField := oracleField.FieldByName(vars.Datasource)
	if !datasourceField.IsValid() {
		panic(oracleErr)
	}
	datasource := datasourceField.String()

	var err error
	if key != "" {
		password, err = crypto.Decrypt(key, password)
		if err != nil {
			return nil, err
		}
	}

	dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`, user, password, datasource)
	return sql.Open(godror, dsn)
}
