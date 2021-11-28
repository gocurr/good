package rocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

// Producer the global rocketmq producer
var Producer rocketmq.Producer

var AccessKey string
var SecretKey string
var Addr []string

// Init inits rocketMQProducer
func Init(c *conf.Configuration) error {
	mqConf := c.RocketMq

	var err error
	AccessKey, err = crypto.Decrypt(c.Secure.Key, mqConf.AccessKey)
	if err != nil {
		return err
	}
	SecretKey, err = crypto.Decrypt(c.Secure.Key, mqConf.SecretKey)
	if err != nil {
		return err
	}
	Addr = mqConf.Addr

	Producer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(Addr)),
		producer.WithRetry(mqConf.Retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: AccessKey,
			SecretKey: SecretKey,
		}))
	return err
}
