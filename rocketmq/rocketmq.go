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

// NewProducer returns rocketmq.Producer and error
func NewProducer(c *conf.Configuration) (rocketmq.Producer, error) {
	accessKey, secretKey, addr, retry, err := decr(c)
	if err != nil {
		return nil, err
	}

	return rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		producer.WithRetry(retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
}

// NewConsumer creates a rocketmq.PushConsumer and error
func NewConsumer(c *conf.Configuration, group string) (rocketmq.PushConsumer, error) {
	accessKey, secretKey, addr, retry, err := decr(c)
	if err != nil {
		return nil, err
	}

	return rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		consumer.WithRetry(retry),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}),
	)
}

func decr(c *conf.Configuration) (string, string, []string, int, error) {
	if c == nil {
		return "", "", nil, 0, rocketmqErr
	}
	mq := c.RocketMq
	if mq == nil {
		return "", "", nil, 0, rocketmqErr
	}

	var accessKey string
	var secretKey string
	var err error
	if c.Secure == nil || c.Secure.Key == "" {
		accessKey = mq.AccessKey
		secretKey = mq.SecretKey
	} else {
		accessKey, err = crypto.Decrypt(c.Secure.Key, mq.AccessKey)
		if err != nil {
			return "", "", nil, 0, err
		}
		secretKey, err = crypto.Decrypt(c.Secure.Key, mq.SecretKey)
		if err != nil {
			return "", "", nil, 0, err
		}
	}
	return accessKey, secretKey, mq.Addr, mq.Retry, nil
}
