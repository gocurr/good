package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
	_ "github.com/godror/godror"
)

// DB the global database object
var DB *sql.DB

const godror = "godror"

var oracleErr = errors.New("bad mysql configuration")

// Init opens Oracle
func Init(c *conf.Configuration) error {
	if c == nil {
		return oracleErr
	}
	db := c.Oracle
	if db == nil {
		return oracleErr
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

	ds := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`, db.User, pw, db.Datasource)
	DB, err = sql.Open(godror, ds)
	return err
}
