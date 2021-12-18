package examples

import (
	"context"
	"database/sql"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/mysql"
	log "github.com/sirupsen/logrus"
	"testing"
)

var err error
var db *sql.DB

func Test_Mysql(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		panic(err)
	}

	db, err = mysql.Open(c)
	if err != nil {
		//panic(err)
	}
	/*Panic(err)
	insert("joy")
	query()
	del("joy")
	query()*/
}

func insert(name string) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	Panic(err)
	defer func() { _ = tx.Rollback() }()

	_, err = db.ExecContext(ctx, "insert into names (name) values(?)", name)
	Panic(err)
	err = tx.Commit()
	Panic(err)
}

func query() {
	rows, err := db.Query("select name from names")
	Panic(err)
	defer func() { _ = rows.Close() }()
	var names []string

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		Panic(err)
		names = append(names, name)
	}
	log.Info(names)
}

func del(name string) {
	tx, err := db.BeginTx(context.Background(), nil)
	Panic(err)
	defer func() { _ = tx.Rollback() }()

	_, err = db.Exec("delete from names where name = ?", name)
	err = tx.Commit()
	Panic(err)
}
