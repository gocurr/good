package good

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

var tsc *tablestore.TableStoreClient

// initTableStore inits tsc
func initTableStore() error {
	ts := conf.TableStore
	if ts == nil {
		return nil
	}

	id, err := decrypt(ts.AccessKeyId)
	if err != nil {
		return err
	}
	secret, err := decrypt(ts.AccessKeySecret)
	if err != nil {
		return err
	}
	tsc = tablestore.NewClient(ts.EndPoint, ts.InstanceName, id, secret)
	return nil
}

// TableStoreClient returns tsc
func TableStoreClient() *tablestore.TableStoreClient {
	return tsc
}
