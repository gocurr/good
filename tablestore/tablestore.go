package tablestore

import (
	"errors"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/pre"
	"reflect"
)

var err = errors.New("bad tablestore configuration")

// New returns *tablestore.TableStoreClient and error
func New(i interface{}) (*tablestore.TableStoreClient, error) {
	if i == nil {
		return nil, err
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

	tablestoreField := c.FieldByName(pre.TableStore)
	if !tablestoreField.IsValid() {
		return nil, err
	}

	endPointField := tablestoreField.FieldByName(consts.EndPoint)
	if !endPointField.IsValid() {
		return nil, err
	}
	endPoint := endPointField.String()

	instanceNameField := tablestoreField.FieldByName(consts.InstanceName)
	if !instanceNameField.IsValid() {
		return nil, err
	}
	instanceName := instanceNameField.String()

	accessKeyIdField := tablestoreField.FieldByName(consts.AccessKeyId)
	if !accessKeyIdField.IsValid() {
		return nil, err
	}
	accessKeyId := accessKeyIdField.String()

	accessKeySecretField := tablestoreField.FieldByName(consts.AccessKeySecret)
	if !accessKeySecretField.IsValid() {
		return nil, err
	}
	accessKeySecret := accessKeySecretField.String()

	var err error
	if key != "" {
		accessKeyId, err = crypto.Decrypt(key, accessKeyId)
		if err != nil {
			return nil, err
		}
		accessKeySecret, err = crypto.Decrypt(key, accessKeySecret)
		if err != nil {
			return nil, err
		}
	}

	return tablestore.NewClient(endPoint, instanceName, accessKeyId, accessKeySecret), nil
}
