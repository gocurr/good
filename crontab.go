package good

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
)

var jobs = make(map[string]cronctl.Job)

// initCrontab inits crontab
func initCrontab() {
	for name, c := range conf.Crontab {
		jobs[name] = cronctl.Job{
			Spec: c.Spec,
		}
	}
}

// NameFns name-function pairs
type NameFns []struct {
	Name string
	Fn   func()
}

func StartCrontab(nameFns NameFns) error {
	for _, pair := range nameFns {
		if err := register(pair.Name, pair.Fn); err != nil {
			return err
		}
	}

	// create a crontab
	crontab, err := cronctl.Create(jobs, cronctl.DefaultLogger{})
	if err != nil {
		return err
	}

	// startup crontab
	return crontab.Startup()
}

// register binds cron to function fn
func register(name string, fn func()) error {
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
