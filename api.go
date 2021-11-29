package good

import (
	"database/sql"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/db"
	redisdb "github.com/gocurr/good/redis"
	mq "github.com/gocurr/good/rocketmq"
	ts "github.com/gocurr/good/tablestore"
)

// DB returns db.Db
func DB() *sql.DB {
	return db.Db
}

// RocketMQProducer returns rocketMQProducer
func RocketMQProducer() rocketmq.Producer {
	return mq.Producer
}

// CreateRocketMQConsumer creates a rocketmq.PushConsumer via group
func CreateRocketMQConsumer(group string) (rocketmq.PushConsumer, error) {
	return rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(mq.Addr)),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: mq.AccessKey,
			SecretKey: mq.SecretKey,
		}),
	)
}

// TableStoreClient returns tsc
func TableStoreClient() *tablestore.TableStoreClient {
	return ts.TSC
}

// Redis returns rdb
func Redis() *redis.Client {
	return redisdb.Rdb
}

// NameFns name-function pairs
type NameFns struct {
	Name string
	Fn   func()
}

// nameFns name function pairs
var nameFns []*NameFns

// RegisterCron registers name-function to crontab
func RegisterCron(name string, fn func()) {
	nameFns = append(nameFns, &NameFns{
		Name: name,
		Fn:   fn,
	})
}

// StartCrontab calls crontab.StartCrontab
func StartCrontab() error {
	if !configured {
		tryConfig()
	}
	for _, nf := range nameFns {
		if err := crontab.Register(nf.Name, nf.Fn); err != nil {
			return err
		}
	}
	return crontab.StartCrontab()
}

// Custom returns custom field
func Custom(name string) interface{} {
	if !configured {
		tryConfig()
	}
	field, ok := custom[name]
	if ok {
		return field
	}
	return nil
}
