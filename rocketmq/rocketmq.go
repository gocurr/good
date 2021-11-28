package rocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

var Producer rocketmq.Producer

var accessKey string
var secretKey string
var addr []string

// Init inits rocketMQProducer
func Init(c *conf.Configuration) error {
	mqConf := c.RocketMq

	var err error
	accessKey, err = crypto.Decrypt(c.Secure.Key, mqConf.AccessKey)
	if err != nil {
		return err
	}
	secretKey, err = crypto.Decrypt(c.Secure.Key, mqConf.SecretKey)
	if err != nil {
		return err
	}
	addr = mqConf.Addr

	Producer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		producer.WithRetry(mqConf.Retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))

	return err
}
