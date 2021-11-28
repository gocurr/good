package crontab

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
	"github.com/gocurr/good/conf"
	"github.com/robfig/cron/v3"
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
func StartCrontab(customLogger cron.Logger) error {
	// filter bad jobs
	var goodJobs = make(map[string]cronctl.Job)
	for k, v := range jobs {
		if k != "" && v.Spec != "" && v.Fn != nil {
			goodJobs[k] = v
		}
	}

	// create a crontab
	crontab, err := cronctl.Create(goodJobs, customLogger)
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
