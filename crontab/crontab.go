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

var errCrontab = errors.New("crontab: bad crontab configuration")
var errStart = errors.New("crontab: start failed")

type Crontab struct {
	crontab *cronctl.Crontab // the crontab entity
	enable  bool             // enable to Start
	discard bool             // discards logs

	mu      sync.Mutex             // protects the remaining fields
	jobs    map[string]cronctl.Job // name-job mapping
	started bool                   // state of the crontab
}

// New returns a new Crontab.
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

	var logDiscard bool
	logDiscardField := crontabField.FieldByName(consts.LogDiscard)
	if logDiscardField.IsValid() {
		logDiscard = logDiscardField.Bool()
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

	return &Crontab{enable: enable, discard: logDiscard, jobs: jobs}, nil
}

// Start starts up the crontab.
func (c *Crontab) Start() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}

	if c.started {
		return errors.New("crontab: started already")
	}

	c.started = true // Set the state.

	var goodJobs = make(map[string]cronctl.Job)
	for k, v := range c.jobs {
		if k != "" && v.Spec != "" && v.Fn != nil {
			goodJobs[k] = v
		}
	}

	// Create a crontab.
	var crontab *cronctl.Crontab
	var err error
	if c.discard {
		crontab, err = cronctl.Create(goodJobs, cronctl.Discard)
	} else {
		crontab, err = cronctl.Create(goodJobs, cronctl.Logrus)
	}
	if err != nil {
		return err
	}

	// Startup the crontab.
	if err = crontab.Startup(); err != nil {
		return err
	}

	// Set crontab entity.
	c.crontab = crontab
	return nil
}

// Bind binds the specific name to the given function.
func (c *Crontab) Bind(name string, fn func()) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}
	if c.started {
		return errors.New("crontab: cannot Bind after Start")
	}

	job, ok := c.jobs[name]
	if !ok {
		return fmt.Errorf("crontab: cannot find '%s' in configuration", name)
	}

	c.jobs[name] = cronctl.Job{
		Spec: job.Spec,
		Fn:   fn,
	}
	return nil
}

// Register registers a new cron by the given name, spec and function.
func (c *Crontab) Register(name, spec string, fn func()) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}
	if c.started {
		return errors.New("crontab: cannot Register after Start")
	}

	c.jobs[name] = cronctl.Job{
		Spec: spec,
		Fn:   fn,
	}
	return nil
}

// Suspend suspends the crontab.
func (c *Crontab) Suspend() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}
	if !c.started {
		return errors.New("crontab: cannot Suspend before Start")
	}

	if c.crontab == nil {
		return errStart
	}

	return c.crontab.Suspend()
}

// Disable stops the given job.
func (c *Crontab) Disable(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}
	if !c.started {
		return errors.New("crontab: cannot Disable before Start")
	}

	if c.crontab == nil {
		return errStart
	}

	return c.crontab.Disable(name)
}

// Enable restarts the given job.
func (c *Crontab) Enable(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}
	if !c.started {
		return errors.New("crontab: cannot Enable before Start")
	}

	if c.crontab == nil {
		return errStart
	}

	return c.crontab.Enable(name)
}

// Continue restarts the crontab.
func (c *Crontab) Continue() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enable {
		return nil
	}

	if !c.started {
		return errors.New("crontab: cannot Continue before Start")
	}

	if c.crontab == nil {
		return errStart
	}

	return c.crontab.Continue()
}
