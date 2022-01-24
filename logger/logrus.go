package logger

import (
	"errors"
	"fmt"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/pre"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"
)

const timestamp = "timestamp"

var (
	ErrLogrus  = errors.New("logger: bad logrus configuration")
	ErrGraylog = errors.New("logger: bad graylog configuration")
)

// Set configures logrus.
func Set(i interface{}) error {
	if i == nil {
		return ErrLogrus
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	logrusField := c.FieldByName(pre.Logrus)
	if !logrusField.IsValid() {
		return ErrLogrus
	}

	var timeFormat = consts.DefaultTimeFormat
	timeFormatField := logrusField.FieldByName(consts.TimeFormat)
	if timeFormatField.IsValid() {
		f := timeFormatField.String()
		if f != "" {
			timeFormat = f
		}
	}

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: timeFormat,
		FullTimestamp:   true,
	})

	ttyDiscardField := logrusField.FieldByName(consts.TTYDiscard)
	if ttyDiscardField.IsValid() {
		if ttyDiscardField.Bool() {
			// Discard logs.
			log.SetOutput(ioutil.Discard)
		}
	}

	graylogField := logrusField.FieldByName(consts.Graylog)
	if graylogField.IsValid() {
		enableField := graylogField.FieldByName(consts.Enable)
		if enableField.IsValid() {
			enable := enableField.Bool()
			if enable {
				hostField := graylogField.FieldByName(consts.Host)
				if !hostField.IsValid() {
					return ErrGraylog
				}
				portField := graylogField.FieldByName(consts.Port)
				if !portField.IsValid() {
					return ErrGraylog
				}
				extraField := graylogField.FieldByName(consts.Extra)
				if !extraField.IsValid() {
					return ErrGraylog
				}

				host := hostField.String()
				port := portField.Int()
				if host == "" || port == 0 {
					return ErrGraylog
				}

				var extra = make(map[string]interface{})
				iter := extraField.MapRange()
				for iter.Next() {
					key := iter.Key().String()
					val := iter.Value().Interface()
					extra[key] = val
				}
				extra[timestamp] = time.Now().Format(timeFormat)

				overrideExtraByEnv(extra)

				hook := graylog.NewAsyncGraylogHook(fmt.Sprintf("%s:%d", host, port), extra)
				defer hook.Flush()
				log.AddHook(hook)
			}
		}
	}

	return nil
}

// overrideExtraByEnv overrides extra by environment variables.
func overrideExtraByEnv(extra map[string]interface{}) {
	const separator = "_"
	for name := range extra {
		key := strings.ToUpper(strings.Join([]string{consts.Graylog, consts.Extra, name}, separator))
		newVal := os.Getenv(key)
		if newVal != "" {
			extra[name] = newVal
		}
	}
}
