package rocketmq

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/pre"
	"reflect"
)

var err = errors.New("bad rocketmq configuration")

// NewProducer returns rocketmq.Producer and error
func NewProducer(i interface{}) (rocketmq.Producer, error) {
	accessKey, secretKey, addrs, retry, err := decrypt(i)
	if err != nil {
		return nil, err
	}

	return rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(addrs)),
		producer.WithRetry(retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
}

// NewConsumer returns rocketmq.PushConsumer and error
func NewConsumer(i interface{}, group string) (rocketmq.PushConsumer, error) {
	accessKey, secretKey, addrs, retry, err := decrypt(i)
	if err != nil {
		return nil, err
	}

	return rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(addrs)),
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
		return "", "", nil, 0, err
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

	rocketmqField := c.FieldByName(pre.RocketMq)
	if !rocketmqField.IsValid() {
		return "", "", nil, 0, err
	}

	addrsField := rocketmqField.FieldByName(consts.Addrs)
	if !addrsField.IsValid() {
		return "", "", nil, 0, err
	}
	var addrs []string
	for i := 0; i < addrsField.Len(); i++ {
		element := addrsField.Index(i)
		addrs = append(addrs, element.String())
	}

	var accessKey string
	accessKeyField := rocketmqField.FieldByName(consts.AccessKey)
	if accessKeyField.IsValid() {
		accessKey = accessKeyField.String()
	}

	var secretKey string
	secretKeyField := rocketmqField.FieldByName(consts.SecretKey)
	if secretKeyField.IsValid() {
		secretKey = secretKeyField.String()
	}

	var retry int64
	retryField := rocketmqField.FieldByName(consts.Retry)
	if retryField.IsValid() {
		retry = retryField.Int()
	}

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
	return accessKey, secretKey, addrs, int(retry), nil
}
