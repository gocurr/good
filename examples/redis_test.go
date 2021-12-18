package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/redis"
	"testing"
)

func Test_Redis(t *testing.T) {
	c, err := conf.New("../app.yaml")
	Panic(err)
	_, err = redis.New(c, 1)
	//Panic(err)

	//var ctx = context.Background()
	//result, _ := rdb.Get(ctx, "abc").Result()
	//log.Info(result)
	//
	//rdb.Set(ctx, "abc", "better", 0)
	//result, _ = rdb.Get(ctx, "good").Result()
	//log.Info(result)
}
