package crontab

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
	"github.com/gocurr/good/conf"
)

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

// StartCrontab starts up crontab
func StartCrontab() error {
	// create a crontab
	crontab, err := cronctl.Create(jobs, cronctl.DefaultLogger{})
	if err != nil {
		return err
	}

	// startup crontab
	return crontab.Startup()
}

// Register binds cron to function fn
func Register(name string, fn func()) error {
	job, ok := jobs[name]
	if !ok {
		return errors.New(fmt.Sprintf("cron [%s] does not exist", name))
	}

	jobs[name] = cronctl.Job{
		Spec: job.Spec,
		Fn:   fn,
	}

	return nil
}
