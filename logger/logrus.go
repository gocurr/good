package logger

import (
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gocurr/good/conf"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
)

// Init inits logrus
func Init(c *conf.Configuration) {
	// set graylog
	logrus := c.Logrus
	gray := logrus.GrayLog
	hook := graylog.NewAsyncGraylogHook(gray.Host+":"+strconv.Itoa(gray.Port), gray.Extra)
	defer hook.Flush()
	log.AddHook(hook)

	var format = "2006-01-02 15:04:05"
	if logrus.Format != "" {
		format = logrus.Format
	}
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: format,
		FullTimestamp:   true,
	})

	// set tty
	if !logrus.TTY {
		log.SetOutput(ioutil.Discard)
	}
}
