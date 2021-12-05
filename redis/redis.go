package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/crypto"
	"github.com/gocurr/good/vars"
	"reflect"
)

var redisErr = errors.New("bad redis configuration")

// New returns *redis.Client and error
func New(i interface{}) (*redis.Client, error) {
	if i == nil {
		return nil, redisErr
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

	redisField := c.FieldByName(vars.Redis)
	if !redisField.IsValid() {
		return nil, redisErr
	}

	hostField := redisField.FieldByName(consts.Host)
	if !hostField.IsValid() {
		return nil, redisErr
	}
	host := hostField.String()

	portField := redisField.FieldByName(consts.Port)
	if !portField.IsValid() {
		return nil, redisErr
	}
	port := portField.Int()

	passwordField := redisField.FieldByName(consts.Password)
	if !passwordField.IsValid() {
		return nil, redisErr
	}
	password := passwordField.String()

	dbField := redisField.FieldByName(consts.DB)
	if !dbField.IsValid() {
		return nil, redisErr
	}
	db := dbField.Int()

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
		DB:       int(db),
	})
	_, err = rdb.Ping(context.Background()).Result()
	return rdb, err
}
