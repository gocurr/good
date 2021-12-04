package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/mysql"
)

func mysqlOp(c *conf.Configuration) {
	db, err := mysql.Get(c)
	Panic(err)
	insert(db, "joy")
	query(db)
	fmt.Println("-----")
	del(db, "joy")
	query(db)
}

func insert(db *sql.DB, name string) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	Panic(err)
	defer func() { _ = tx.Rollback() }()

	_, err = db.ExecContext(ctx, "insert into names (name) values(?)", name)
	Panic(err)
	err = tx.Commit()
	Panic(err)
}

func query(db *sql.DB) {
	rows, err := db.Query("select name from names")
	Panic(err)
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		Panic(err)
		fmt.Println(name)
	}
}

func del(db *sql.DB, name string) {
	tx, err := db.BeginTx(context.Background(), nil)
	Panic(err)
	defer func() { _ = tx.Rollback() }()

	_, err = db.Exec("delete from names where name = ?", name)
	err = tx.Commit()
	Panic(err)
}
