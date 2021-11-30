package good

import (
	"database/sql"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/mysql"
	"github.com/gocurr/good/oracle"
	redisdb "github.com/gocurr/good/redis"
	mq "github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/sugar"
	ts "github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

// Oracle returns oracle.DB
func Oracle() *sql.DB {
	return oracle.DB
}

// Mysql returns mysql.DB
func Mysql() *sql.DB {
	return mysql.DB
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

// RegisterCron registers a new cron
func RegisterCron(name, spec string, fn func()) {
	crontab.Register(name, spec, fn)
}

// NameFn name-function pair
type NameFn struct {
	Name string
	Fn   func()
}

// nameFns name-function pairs
var nameFns []*NameFn

// BindCron binds name-function to crontab
func BindCron(name string, fn func()) {
	if startCronDone {
		return
	}
	nameFns = append(nameFns, &NameFn{
		Name: name,
		Fn:   fn,
	})
}

// startCronOnce for StartCrontab
var startCronOnce sync.Once

// startCronDone reports StartCrontab invoked
var startCronDone bool

// StartCrontab calls crontab.Start
func StartCrontab() {
	startCronOnce.Do(func() {
		startCronDone = true // set done
		if !configured {
			tryConfig()
		}
		for _, nf := range nameFns {
			if err := crontab.Bind(nf.Name, nf.Fn); err != nil {
				log.Errorf("%v", err)
			}
		}
		if err := crontab.Start(); err != nil {
			log.Errorf("%v", err)
		}
	})
}

// ServerMux set serverMux
func ServerMux(mux *http.ServeMux) {
	sugar.ServerMux(mux)
}

// Route binds route path to fn
func Route(route string, fn func(http.ResponseWriter, *http.Request)) {
	sugar.Route(route, fn)
}

// Fire http server entry
func Fire(callbacks ...func()) {
	if !configured {
		tryConfig()
	}
	sugar.Fire(configuration, callbacks...)
}
