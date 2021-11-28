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
	gray := c.Logrus.GrayLog
	hook := graylog.NewAsyncGraylogHook(gray.Host+":"+strconv.Itoa(gray.Port), gray.Extra)
	defer hook.Flush()
	log.AddHook(hook)

	// set tty
	if !c.Logrus.TTY {
		log.SetOutput(ioutil.Discard)
	}
}
