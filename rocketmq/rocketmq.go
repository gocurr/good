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

// Init inits rocketMQProducer
func Init(c *conf.Configuration) error {
	mqConf := c.RocketMq

	accessKey, err := crypto.Decrypt(c.Secure.Key, mqConf.AccessKey)
	if err != nil {
		return err
	}
	secretKey, err := crypto.Decrypt(c.Secure.Key, mqConf.SecretKey)
	if err != nil {
		return err
	}

	Producer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(mqConf.Addr)),
		producer.WithRetry(mqConf.Retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
	return err
}
