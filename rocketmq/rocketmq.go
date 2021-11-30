package rocketmq

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

var rocketmqErr = errors.New("bad rocketmq configuration")

// Producer the global rocketmq producer
var Producer rocketmq.Producer

var accessKey string
var secretKey string
var addr []string

// Init inits rocketMQProducer
func Init(c *conf.Configuration) error {
	if c == nil {
		return rocketmqErr
	}
	mq := c.RocketMq
	if mq == nil {
		return rocketmqErr
	}

	var err error
	if c.Secure == nil || c.Secure.Key == "" {
		accessKey = mq.AccessKey
		secretKey = mq.SecretKey
	} else {
		accessKey, err = crypto.Decrypt(c.Secure.Key, mq.AccessKey)
		if err != nil {
			return err
		}
		secretKey, err = crypto.Decrypt(c.Secure.Key, mq.SecretKey)
		if err != nil {
			return err
		}
	}

	addr = mq.Addr

	Producer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		producer.WithRetry(mq.Retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
	return err
}

// CreateConsumer creates a rocketmq.PushConsumer via group
func CreateConsumer(group string) (rocketmq.PushConsumer, error) {
	return rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}),
	)
}
