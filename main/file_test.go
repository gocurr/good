package main

import (
	"github.com/gocurr/good/streaming"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	ss "strings"
	"testing"
	"time"
)

func Test_File(t *testing.T) {
	bytes, err := ioutil.ReadFile("no.txt")
	if err != nil {
		return
	}
	ls := string(bytes)

	var lines strings = ss.Split(ls, "\n")
	wordLen := 15

	//parallelStream := streaming.ParallelOf(lines)
	stream := streaming.Of(lines)

	/*since := time.Now()
	count := parallelStream.
		FlatMap(func(i interface{}) interface{} {
			return strings.Split(i.(string), " ")
		}).Filter(func(i interface{}) bool {
		return len(i.(string)) > wordLen
	}).Distinct().Count()
	logrus.Infof("%v took %v", count, time.Since(since))*/

	since := time.Now()
	count := stream.
		FlatMap(func(i interface{}) interface{} {
			return ss.Split(i.(string), " ")
		}).Filter(func(i interface{}) bool {
		return len(i.(string)) > wordLen
	}).Distinct().Count()
	logrus.Infof("%v took %v", count, time.Since(since))

	/*since := time.Now()
	var nothing struct{}
	var distinct = make(map[string]struct{})
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, w := range words {
			if len(w) > wordLen {
				distinct[w] = nothing
			}
		}
	}
	logrus.Infof("%v took %v", len(distinct), time.Since(since))*/
}
