package tablestore

import (
	"errors"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
)

var tablestoreErr = errors.New("bad tablestore configuration")

// Get returns *tablestore.TableStoreClient
func Get(c *conf.Configuration) (*tablestore.TableStoreClient, error) {
	if c == nil {
		return nil, tablestoreErr
	}
	ts := c.TableStore
	if ts == nil {
		return nil, tablestoreErr
	}

	var err error
	var id, secret string
	if c.Secure == nil || c.Secure.Key == "" {
		id = ts.AccessKeyId
		secret = ts.AccessKeySecret
	} else {
		id, err = crypto.Decrypt(c.Secure.Key, ts.AccessKeyId)
		if err != nil {
			return nil, err
		}
		secret, err = crypto.Decrypt(c.Secure.Key, ts.AccessKeySecret)
		if err != nil {
			return nil, err
		}
	}

	return tablestore.NewClient(ts.EndPoint, ts.InstanceName, id, secret), nil
}
