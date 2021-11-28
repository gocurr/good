package good

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

var tsc *tablestore.TableStoreClient

// initTableStore inits tsc
func initTableStore() {
	ts := conf.TableStore
	if ts == nil {
		return
	}
	tsc = tablestore.NewClient(ts.EndPoint, ts.InstanceName, ts.AccessKeyId, ts.AccessKeySecret)
}

// TableStoreClient returns tsc
func TableStoreClient() *tablestore.TableStoreClient {
	return tsc
}
