package logger

import (
	"errors"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gocurr/good/conf"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"time"
)

var logrusErr = errors.New("bad logrus configuration")

// Set configures logrus
func Set(c *conf.Configuration) error {
	if c == nil {
		return logrusErr
	}
	// set graylog
	l := c.Logrus
	if l == nil {
		return logrusErr
	}

	// set logrus output format
	var format = "2006-01-02 15:04:05"
	if l.Format != "" {
		format = l.Format
	}
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: format,
		FullTimestamp:   true,
	})

	// set tty
	if !l.TTY {
		log.SetOutput(ioutil.Discard)
	}

	gray := l.GrayLog
	if gray != nil {
		if gray.Enable {
			host := gray.Host
			port := gray.Port
			if host == "" || port == 0 {
				return logrusErr
			}

			if gray.Extra == nil {
				gray.Extra = make(map[string]interface{})
			}
			gray.Extra["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
			hook := graylog.NewAsyncGraylogHook(host+":"+strconv.Itoa(port), gray.Extra)
			defer hook.Flush()
			log.AddHook(hook)
		}
	}
	return nil
}
