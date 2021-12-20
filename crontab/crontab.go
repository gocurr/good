package crontab

import (
	"errors"
	"fmt"
	"github.com/gocurr/cronctl"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/pre"
	"reflect"
	"sync"
)

var errCrontab = errors.New("bad crontab configuration")

// Crontab jobs wrapper
type Crontab struct {
	enable  bool                   // enable to Start
	jobs    map[string]cronctl.Job // name:job mapping
	once    sync.Once              // start once
	done    bool                   // reports Start invoked
	discard bool                   // discard log
}

// New returns a new Crontab
func New(i interface{}, discard ...bool) (*Crontab, error) {
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

	var _discard bool
	if len(discard) > 0 {
		_discard = discard[0]
	}

	return &Crontab{enable: enable, jobs: jobs, discard: _discard}, nil
}

// Start starts up crontab
func (c *Crontab) Start() error {
	if !c.enable {
		return nil
	}

	var err error
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
		var crontab *cronctl.Crontab
		if c.discard {
			crontab, err = cronctl.Create(goodJobs, cronctl.Discard)
		} else {
			crontab, err = cronctl.Create(goodJobs, cronctl.Logrus)
		}

		if err == nil {
			// startup crontab
			err = crontab.Startup()
		}
	})

	return err
}

// Bind binds name to function
func (c *Crontab) Bind(name string, fn func()) error {
	if !c.enable {
		return nil
	}
	if c.done {
		return errors.New("cannot Bind after Start")
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
		return errors.New("cannot Register after Start")
	}

	c.jobs[name] = cronctl.Job{
		Spec: spec,
		Fn:   fn,
	}
	return nil
}
