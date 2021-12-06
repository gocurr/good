package main

import (
	"github.com/gocurr/good/httpclient"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_HttpClient(t *testing.T) {
	type msg struct {
		Text string
	}
	var m = msg{Text: "great"}
	var out msg
	err := httpclient.PostJSON("http://127.0.0.1:9091", m, &out)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(out)
}
