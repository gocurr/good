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
	logrus := c.Logrus
	gl := logrus.GrayLog
	addr := gl.Host + ":" + strconv.Itoa(gl.Port)
	extra := gl.Extra
	hook := graylog.NewAsyncGraylogHook(addr, extra)
	defer hook.Flush()
	log.AddHook(hook)

	if !logrus.TTY {
		log.SetOutput(ioutil.Discard)
	}
}
