package tablestore

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

// TSC the global tablestore client
var TSC *tablestore.TableStoreClient

// Init inits tsc
func Init(c *conf.Configuration) error {
	ts := c.TableStore
	id, err := crypto.Decrypt(c.Secure.Key, ts.AccessKeyId)
	if err != nil {
		return err
	}
	secret, err := crypto.Decrypt(c.Secure.Key, ts.AccessKeySecret)
	if err != nil {
		return err
	}

	TSC = tablestore.NewClient(ts.EndPoint, ts.InstanceName, id, secret)
	return nil
}
