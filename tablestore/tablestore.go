package tablestore

import (
	"errors"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

var tablestoreErr = errors.New("bad tablestore configuration")

// TSC the global tablestore client
var TSC *tablestore.TableStoreClient

// Init inits tsc
func Init(c *conf.Configuration) error {
	if c == nil {
		return tablestoreErr
	}
	ts := c.TableStore
	if ts == nil {
		return tablestoreErr
	}

	var err error
	var id, secret string
	if c.Secure == nil || c.Secure.Key == "" {
		id = ts.AccessKeyId
		secret = ts.AccessKeySecret
	} else {
		id, err = crypto.Decrypt(c.Secure.Key, ts.AccessKeyId)
		if err != nil {
			return err
		}
		secret, err = crypto.Decrypt(c.Secure.Key, ts.AccessKeySecret)
		if err != nil {
			return err
		}
	}

	TSC = tablestore.NewClient(ts.EndPoint, ts.InstanceName, id, secret)
	return nil
}
