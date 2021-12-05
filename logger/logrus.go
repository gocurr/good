package logger

import (
	"errors"
	"fmt"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gocurr/good/vars"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"reflect"
	"time"
)

var LogrusErr = errors.New("bad logrus configuration")
var GraylogErr = errors.New("bad graylog configuration")

// Set configures logrus
func Set(i interface{}) error {
	if i == nil {
		return LogrusErr
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	logrusField := c.FieldByName(vars.Logrus)
	if !logrusField.IsValid() {
		return LogrusErr
	}

	var format = "2006-01-02 15:04:05"
	formatField := logrusField.FieldByName(vars.Format)
	if formatField.IsValid() {
		f := formatField.String()
		if f != "" {
			format = f
		}
	}

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: format,
		FullTimestamp:   true,
	})

	ttyDiscardField := logrusField.FieldByName(vars.TTYDiscard)
	if ttyDiscardField.IsValid() {
		if ttyDiscardField.Bool() {
			// discard
			log.SetOutput(ioutil.Discard)
		}
	}

	graylogField := logrusField.FieldByName(vars.GrayLog)
	if graylogField.IsValid() {
		enableField := graylogField.FieldByName(vars.Enable)
		if enableField.IsValid() {
			enable := enableField.Bool()
			if enable {
				hostField := graylogField.FieldByName(vars.Host)
				if !hostField.IsValid() {
					return GraylogErr
				}
				portField := graylogField.FieldByName(vars.Port)
				if !portField.IsValid() {
					return GraylogErr
				}
				extraField := graylogField.FieldByName(vars.Extra)
				if !extraField.IsValid() {
					return GraylogErr
				}

				host := hostField.String()
				port := portField.Int()
				if host == "" || port == 0 {
					return GraylogErr
				}

				var extra = make(map[string]interface{})
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
		}
	}

	return nil
}
