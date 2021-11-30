package crontab

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
	"github.com/gocurr/good/conf"
	log "github.com/sirupsen/logrus"
	"sync"
)

var crontabErr = errors.New("cannot Bind after Start()")

// jobs global crontab
var jobs = make(map[string]cronctl.Job)

// Init inits crontab
func Init(c *conf.Configuration) {
	for name, c := range c.Crontab {
		jobs[name] = cronctl.Job{
			Spec: c.Spec,
		}
	}
}

// once for StartCrontab
var once sync.Once

// done reports StartCrontab invoked
var done bool

// Start starts up crontab
func Start() {
	once.Do(func() {
		done = true // set done

		// filter bad jobs
		var goodJobs = make(map[string]cronctl.Job)
		for k, v := range jobs {
			if k != "" && v.Spec != "" && v.Fn != nil {
				goodJobs[k] = v
			}
		}

		// create a crontab
		crontab, err := cronctl.Create(goodJobs, cronctl.DefaultLogger{})
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		// startup crontab
		if err := crontab.Startup(); err != nil {
			log.Errorf("%v", err)
		}
	})

}

// Bind binds cron to function fn
func Bind(name string, fn func()) error {
	if done {
		return crontabErr
	}
	job, ok := jobs[name]
	if !ok {
		return errors.New(fmt.Sprintf("cron '%s' does not exist", name))
	}

	jobs[name] = cronctl.Job{
		Spec: job.Spec,
		Fn:   fn,
	}
	return nil
}

// Register registers a new cron
func Register(name, spec string, fn func()) {
	if done {
		return
	}
	jobs[name] = cronctl.Job{
		Spec: spec,
		Fn:   fn,
	}
}
