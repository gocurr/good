package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/postgres"
	"github.com/gocurr/good/sugar"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestPostgres(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		log.Error(err)
		return
	}

	if err := logger.Set(c); err == logger.LogrusErr {
		panic(err)
	}

	db, err := postgres.Open(c)
	if err != nil {
		log.Error(err)
		return
	}

	rows, err := db.Query("select name from names")
	if err != nil {
		log.Error(err)
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Error(err)
			return
		}
		log.Info(name)
	}

	sugar.Fire(c)
}
