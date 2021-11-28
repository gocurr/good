package good

import (
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
)

func initLogurs() {
	logrus := conf.Logrus
	if logrus == nil {
		return
	}

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
