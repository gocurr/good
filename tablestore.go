package good

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

var TSC *tablestore.TableStoreClient

func initTableStore() {
	ts := conf.TableStore
	if ts == nil {
		return
	}
	TSC = tablestore.NewClient(ts.EndPoint, ts.InstanceName, ts.AccessKeyId, ts.AccessKeySecret)
}

func TableStoreClient() *tablestore.TableStoreClient {
	return TSC
}
