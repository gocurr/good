package logger

import (
	"errors"
	"fmt"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gocurr/good/consts"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"reflect"
	"time"
)

var logrusErr = errors.New("bad logrus configuration")

// Set configures logrus
func Set(i interface{}) error {
	if i == nil {
		panic(logrusErr)
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	logrusField := c.FieldByName(consts.Logrus)
	if !logrusField.IsValid() {
		panic(logrusErr)
	}

	f := logrusField.FieldByName(consts.Format).String()
	// set logrus output format
	var format = "2006-01-02 15:04:05"
	if f != "" {
		format = f
	}
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: format,
		FullTimestamp:   true,
	})

	tty := logrusField.FieldByName(consts.TTY).Bool()
	// set tty
	if !tty {
		log.SetOutput(ioutil.Discard)
	}

	graylogField := logrusField.FieldByName(consts.GrayLog)
	enable := graylogField.FieldByName(consts.Enable).Bool()
	if enable {
		host := graylogField.FieldByName(consts.Host).String()
		port := graylogField.FieldByName(consts.Port).Int()
		if host == "" || port == 0 {
			return logrusErr
		}

		var extra = make(map[string]interface{})
		extraField := graylogField.FieldByName(consts.Extra)
		iter := extraField.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			val := iter.Value().Interface()
			extra[key] = val
		}
		extra["timestamp"] = time.Now().Format("2006-01-02 15:04:05")

		hook := graylog.NewAsyncGraylogHook(fmt.Sprintf("%s:%v", host, port), extra)
		defer hook.Flush()
		log.AddHook(hook)
	}
	return nil
}
