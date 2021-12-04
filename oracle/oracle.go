package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
	_ "github.com/godror/godror"
)

const godror = "godror"

var oracleErr = errors.New("bad oracle configuration")

// Open returns *sql.DB and error
func Open(c *conf.Configuration) (*sql.DB, error) {
	if c == nil {
		return nil, oracleErr
	}
	db := c.Oracle
	if db == nil {
		return nil, oracleErr
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

	dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`, db.User, pw, db.Datasource)
	return sql.Open(godror, dsn)
}
