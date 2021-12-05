package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/postgres"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestPostgres(t *testing.T) {
	c, _ := conf.New("../app.yaml")
	_ = logger.Set(c)

	db, err := postgres.Open(c)
	if err != nil {
		Panic(err)
	}

	rows, err := db.Query("select name from names")
	if err != nil {
		Panic(err)
	}
	defer func() { _ = rows.Close() }()
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			Panic(err)
		}
		names = append(names, name)
	}
	log.Info(names)
}
