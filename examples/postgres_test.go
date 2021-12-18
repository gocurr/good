package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/postgres"
	"testing"
)

func TestPostgres(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		panic(err)
	}
	_ = logger.Set(c)

	_, err = postgres.Open(c)
	if err != nil {
		//Panic(err)
	}

	/*rows, err := db.Query("select name from names")
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
	log.Info(names)*/
}
