package rocketmq

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"reflect"
)

var rocketmqErr = errors.New("bad rocketmq configuration")

// NewProducer returns rocketmq.Producer and error
func NewProducer(c *conf.Configuration) (rocketmq.Producer, error) {
	accessKey, secretKey, addr, retry, err := decrypt(c)
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

// NewConsumer returns rocketmq.PushConsumer and error
func NewConsumer(c *conf.Configuration, group string) (rocketmq.PushConsumer, error) {
	accessKey, secretKey, addr, retry, err := decrypt(c)
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

// decrypt returns decrypted attributes
func decrypt(i interface{}) (string, string, []string, int, error) {
	if i == nil {
		panic(rocketmqErr)
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	var key string
	secureField := c.FieldByName(consts.Secure)
	if secureField.IsValid() {
		keyField := secureField.FieldByName(consts.Key)
		if keyField.IsValid() {
			key = keyField.String()
		}
	}

	rocketmqField := c.FieldByName(consts.RocketMq)
	if !rocketmqField.IsValid() {
		panic(rocketmqErr)
	}

	addrField := rocketmqField.FieldByName(consts.Addr)
	if !addrField.IsValid() {
		panic(rocketmqErr)
	}
	var addr []string
	for i := 0; i < addrField.Len(); i++ {
		element := addrField.Index(i)
		addr = append(addr, element.String())
	}

	accessKeyField := rocketmqField.FieldByName(consts.AccessKey)
	if !accessKeyField.IsValid() {
		panic(rocketmqErr)
	}
	accessKey := accessKeyField.String()

	secretKeyField := rocketmqField.FieldByName(consts.SecretKey)
	if !secretKeyField.IsValid() {
		panic(rocketmqErr)
	}
	secretKey := secretKeyField.String()

	retryField := rocketmqField.FieldByName(consts.Retry)
	if !retryField.IsValid() {
		panic(rocketmqErr)
	}
	retry := retryField.Int()

	var err error
	if key != "" {
		accessKey, err = crypto.Decrypt(key, accessKey)
		if err != nil {
			return "", "", nil, 0, err
		}
		secretKey, err = crypto.Decrypt(key, secretKey)
		if err != nil {
			return "", "", nil, 0, err
		}
	}
	return accessKey, secretKey, addr, int(retry), nil
}
