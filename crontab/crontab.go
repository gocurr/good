package crontab

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/pre"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

var (
	errCrontab  = errors.New("bad crontab configuration")
	errBind     = errors.New("cannot Bind after Start")
	errRegister = errors.New("cannot Register after Start")
)

// Crontab jobs wrapper
type Crontab struct {
	enable bool                   // enable to Start
	jobs   map[string]cronctl.Job // name-job mapping
	once   sync.Once              // for Start
	done   bool                   // reports Start invoked
}

// New Crontab constructor
func New(i interface{}) (*Crontab, error) {
	if i == nil {
		return nil, errCrontab
	}

	var c reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	crontabField := c.FieldByName(pre.Crontab)
	if !crontabField.IsValid() {
		return nil, errCrontab
	}

	var enable bool
	enableField := crontabField.FieldByName(consts.Enable)
	if enableField.IsValid() {
		enable = enableField.Bool()
	}

	var jobs = make(map[string]cronctl.Job)
	specsField := crontabField.FieldByName(consts.Specs)
	if specsField.IsValid() {
		iter := specsField.MapRange()
		for iter.Next() {
			name := iter.Key().String()
			spec := iter.Value().String()
			jobs[name] = cronctl.Job{
				Spec: spec,
			}
		}
	}
	return &Crontab{enable: enable, jobs: jobs}, nil
}

// Start starts up crontab
func (c *Crontab) Start() {
	if !c.enable {
		return
	}
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
		crontab, err := cronctl.Create(goodJobs, cronctl.Logrus)
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

// Bind binds name to function
func (c *Crontab) Bind(name string, fn func()) error {
	if !c.enable {
		return nil
	}
	if c.done {
		return errBind
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
func (c *Crontab) Register(name, spec string, fn func()) error {
	if !c.enable {
		return nil
	}
	if c.done {
		return errRegister
	}

	c.jobs[name] = cronctl.Job{
		Spec: spec,
		Fn:   fn,
	}
	return nil
}
