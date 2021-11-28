package good

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var rocketMQProducer rocketmq.Producer

var accessKey string
var secretKey string
var addr []string

// initRocketMq inits rocketMQProducer
func initRocketMq() error {
	mqConf := conf.RocketMq
	if mqConf == nil {
		return nil
	}

	var err error
	accessKey, err = decrypt(mqConf.AccessKey)
	if err != nil {
		return err
	}
	secretKey, err = decrypt(mqConf.SecretKey)
	if err != nil {
		return err
	}
	addr = mqConf.Addr

	rocketMQProducer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		producer.WithRetry(mqConf.Retry),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))

	return err
}

func RocketMQProducer() rocketmq.Producer {
	return rocketMQProducer
}

// CreateRocketMQConsumer creates a rocketmq.PushConsumer via groupname
func CreateRocketMQConsumer(group string) (rocketmq.PushConsumer, error) {
	return rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(addr)),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}),
	)
}
