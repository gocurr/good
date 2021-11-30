package main

import (
	"context"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/mysql"
)

func mysqlOp() {
	c, err := conf.ReadDefault()
	Panic(err)
	err = mysql.Init(c)
	Panic(err)
	insert("joy")
	query()
	fmt.Println("-----")
	del("joy")
	query()
}

func insert(name string) {
	db := mysql.DB
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
	db := mysql.DB
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

func del(name string) {
	db := mysql.DB
	tx, err := db.BeginTx(context.Background(), nil)
	Panic(err)
	defer func() { _ = tx.Rollback() }()

	_, err = db.Exec("delete from names where name = ?", name)
	err = tx.Commit()
	Panic(err)
}
