package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
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
	secureField := c.FieldByName(consts.Secure)
	if secureField.IsValid() {
		keyField := secureField.FieldByName(consts.Key)
		if keyField.IsValid() {
			key = keyField.String()
		}
	}

	oracleField := c.FieldByName(consts.Oracle)
	if !oracleField.IsValid() {
		panic(oracleErr)
	}

	userField := oracleField.FieldByName(consts.User)
	if !userField.IsValid() {
		panic(oracleErr)
	}
	user := userField.String()

	passwordField := oracleField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		panic(oracleErr)
	}
	password := passwordField.String()

	datasourceField := oracleField.FieldByName(consts.Datasource)
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
