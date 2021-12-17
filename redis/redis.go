package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/pre"
	"reflect"
)

var err = errors.New("bad redis configuration")

// New returns *redis.Client and error
func New(i interface{}, _db ...int) (*redis.Client, error) {
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

	redisField := c.FieldByName(pre.Redis)
	if !redisField.IsValid() {
		return nil, err
	}

	hostField := redisField.FieldByName(consts.Host)
	if !hostField.IsValid() {
		return nil, err
	}
	host := hostField.String()

	portField := redisField.FieldByName(consts.Port)
	if !portField.IsValid() {
		return nil, err
	}
	port := portField.Int()

	passwordField := redisField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		return nil, err
	}
	password := passwordField.String()

	var db int
	if len(_db) == 0 {
		dbField := redisField.FieldByName(consts.DB)
		if !dbField.IsValid() {
			return nil, err
		}
		db = int(dbField.Int())
	} else {
		db = _db[0]
	}

	var err error
	if key != "" {
		password, err = crypto.Decrypt(key, password)
		if err != nil {
			return nil, err
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})
	_, err = rdb.Ping(context.Background()).Result()
	return rdb, err
}
