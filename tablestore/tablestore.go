package tablestore

import (
	"errors"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/vars"
	"reflect"
)

var tablestoreErr = errors.New("bad tablestore configuration")

// New returns *tablestore.TableStoreClient and error
func New(i interface{}) (*tablestore.TableStoreClient, error) {
	if i == nil {
		return nil, tablestoreErr
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	var key string
	secureField := c.FieldByName(vars.Secure)
	if secureField.IsValid() {
		keyField := secureField.FieldByName(vars.Key)
		if keyField.IsValid() {
			key = keyField.String()
		}
	}

	tablestoreField := c.FieldByName(vars.TableStore)
	if !tablestoreField.IsValid() {
		return nil, tablestoreErr
	}

	endPointField := tablestoreField.FieldByName(vars.EndPoint)
	if !endPointField.IsValid() {
		return nil, tablestoreErr
	}
	endPoint := endPointField.String()

	instanceNameField := tablestoreField.FieldByName(vars.InstanceName)
	if !instanceNameField.IsValid() {
		return nil, tablestoreErr
	}
	instanceName := instanceNameField.String()

	accessKeyIdField := tablestoreField.FieldByName(vars.AccessKeyId)
	if !accessKeyIdField.IsValid() {
		return nil, tablestoreErr
	}
	accessKeyId := accessKeyIdField.String()

	accessKeySecretField := tablestoreField.FieldByName(vars.AccessKeySecret)
	if !accessKeySecretField.IsValid() {
		return nil, tablestoreErr
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
