package main

import (
	_ "github.com/gocurr/good/conf"
	_ "github.com/gocurr/good/crontab"
	_ "github.com/gocurr/good/crypto"
	_ "github.com/gocurr/good/httpclient"
	_ "github.com/gocurr/good/logger"
	_ "github.com/gocurr/good/mysql"
	_ "github.com/gocurr/good/oracle"
	_ "github.com/gocurr/good/postgres"
	_ "github.com/gocurr/good/redis"
	_ "github.com/gocurr/good/rocketmq"
	_ "github.com/gocurr/good/server"
	_ "github.com/gocurr/good/sugar"
	_ "github.com/gocurr/good/tablestore"
	_ "testing"
	_ "time"
)

// for pkg-updates only
func main() {

}
