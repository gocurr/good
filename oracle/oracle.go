package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/pre"
	_ "github.com/godror/godror"
	"reflect"
)

const godror = "godror"

var errOracle = errors.New("oracle: bad oracle configuration")

// Open returns an oracle database and reports error encountered.
func Open(i interface{}) (*sql.DB, error) {
	if i == nil {
		return nil, errOracle
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

	oracleField := c.FieldByName(pre.Oracle)
	if !oracleField.IsValid() {
		return nil, errOracle
	}

	userField := oracleField.FieldByName(consts.User)
	if !userField.IsValid() {
		return nil, errOracle
	}
	user := userField.String()

	passwordField := oracleField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		return nil, errOracle
	}
	password := passwordField.String()

	datasourceField := oracleField.FieldByName(consts.Datasource)
	if !datasourceField.IsValid() {
		return nil, errOracle
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
	db, err := sql.Open(godror, dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
