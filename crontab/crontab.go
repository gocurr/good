package crontab

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
	"github.com/gocurr/good/conf"
	log "github.com/sirupsen/logrus"
	"sync"
)

var crontabErr = errors.New("cannot Bind after Start")

// Crontab cron wrapper
type Crontab struct {
	jobs map[string]cronctl.Job // name-job mapping
	once sync.Once              // for Start
	done bool                   // reports Start invoked
}

// New Crontab constructor
func New(c *conf.Configuration) *Crontab {
	var jobs = make(map[string]cronctl.Job)
	for name, c := range c.Crontab {
		jobs[name] = cronctl.Job{
			Spec: c.Spec,
		}
	}
	return &Crontab{jobs: jobs}
}

// Start starts up crontab
func (c *Crontab) Start() {
	c.once.Do(func() {
		c.done = true // set done

		// filter bad jobs
		var goodJobs = make(map[string]cronctl.Job)
		for k, v := range c.jobs {
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

// Bind binds cron-name to function-fn
func (c *Crontab) Bind(name string, fn func()) error {
	if c.done {
		return crontabErr
	}

	job, ok := c.jobs[name]
	if !ok {
		return fmt.Errorf("cannot find '%s' in configuration", name)
	}

	c.jobs[name] = cronctl.Job{
		Spec: job.Spec,
		Fn:   fn,
	}
	return nil
}

// Register registers a new cron
func (c *Crontab) Register(name, spec string, fn func()) {
	if c.done {
		return
	}

	c.jobs[name] = cronctl.Job{
		Spec: spec,
		Fn:   fn,
	}
}
