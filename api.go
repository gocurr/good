package good

import (
	"context"
	"database/sql"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/db"
	redisdb "github.com/gocurr/good/redis"
	mq "github.com/gocurr/good/rocketmq"
	ts "github.com/gocurr/good/tablestore"
)

func DB() *sql.DB {
	return db.Db
}

// RocketMQProducer returns rocketMQProducer
func RocketMQProducer() rocketmq.Producer {
	return mq.Producer
}

// CreateRocketMQConsumer creates a rocketmq.PushConsumer via groupname
func CreateRocketMQConsumer(c *conf.Configuration, group string) (rocketmq.PushConsumer, error) {
	rmq := c.RocketMq
	return rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(rmq.Addr)),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: rmq.AccessKey,
			SecretKey: rmq.SecretKey,
		}),
	)
}

// TableStoreClient returns tsc
func TableStoreClient() *tablestore.TableStoreClient {
	return ts.TSC
}

// Redis returns rdb
func Redis() (*redis.Client, context.Context) {
	return redisdb.Rdb, redisdb.Ctx
}

// NameFns name-function pairs
type NameFns []struct {
	Name string
	Fn   func()
}

func StartCrontab(nameFns NameFns) error {
	var nf crontab.NameFns
	for _, v := range nameFns {
		nf = append(nf, v)
	}
	return crontab.StartCrontab(nf)
}
