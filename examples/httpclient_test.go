package examples

import (
	"github.com/gocurr/good/httpclient"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_HttpClient(t *testing.T) {
	type msg struct {
		Text string
	}
	var m = msg{Text: "great"}
	var out msg
	url := "http://127.0.0.1:9091"
	err := httpclient.PostJSON(url, m, &out, time.Second)
	if err != nil {
		log.Error(err)
	}
	log.Info(out)

	raw, err := httpclient.GetRaw(url, time.Millisecond)
	if err != nil {
		log.Error(err)
	}

	log.Info(string(raw))
}
