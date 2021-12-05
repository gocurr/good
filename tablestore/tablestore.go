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
		panic(tablestoreErr)
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
		panic(tablestoreErr)
	}

	endPointField := tablestoreField.FieldByName(vars.EndPoint)
	if !endPointField.IsValid() {
		panic(tablestoreErr)
	}
	endPoint := endPointField.String()

	instanceNameField := tablestoreField.FieldByName(vars.InstanceName)
	if !instanceNameField.IsValid() {
		panic(tablestoreErr)
	}
	instanceName := instanceNameField.String()

	accessKeyIdField := tablestoreField.FieldByName(vars.AccessKeyId)
	if !accessKeyIdField.IsValid() {
		panic(tablestoreErr)
	}
	accessKeyId := accessKeyIdField.String()

	accessKeySecretField := tablestoreField.FieldByName(vars.AccessKeySecret)
	if !accessKeySecretField.IsValid() {
		panic(tablestoreErr)
	}
	accessKeySecret := accessKeySecretField.String()

	var err error
	if key != "" {
		accessKeyId, err = crypto.Decrypt(key, accessKeyId)
		if err != nil {
			panic(err)
		}
		accessKeySecret, err = crypto.Decrypt(key, accessKeySecret)
		if err != nil {
			panic(err)
		}
	}

	return tablestore.NewClient(endPoint, instanceName, accessKeyId, accessKeySecret), nil
}
