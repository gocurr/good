package rocketmq

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

var rocketmqErr = errors.New("bad rocketmq configuration")

// Producer the global rocketmq producer
var Producer rocketmq.Producer

// exported fields

var AccessKey string
var SecretKey string
var Addr []string

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
		AccessKey = mq.AccessKey
		SecretKey = mq.SecretKey
	} else {
		AccessKey, err = crypto.Decrypt(c.Secure.Key, mq.AccessKey)
		if err != nil {
			return err
		}
		SecretKey, err = crypto.Decrypt(c.Secure.Key, mq.SecretKey)
		if err != nil {
			return err
		}
	}

	Addr = mq.Addr

	Producer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(Addr)),
		producer.WithRetry(mq.Retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: AccessKey,
			SecretKey: SecretKey,
		}))
	return err
}
