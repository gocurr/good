package conf

import (
	"errors"
	"fmt"
	"github.com/gocurr/good/vars"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

var confErr = errors.New("configuration not found")

// Filename returns a configuration name
func Filename() string {
	if _, err := os.Stat(vars.AppYml); err == nil {
		return vars.AppYml
	}
	if _, err := os.Stat(vars.AppYaml); err == nil {
		return vars.AppYaml
	}
	if _, err := os.Stat(vars.ApplicationYml); err == nil {
		return vars.ApplicationYml
	}
	if _, err := os.Stat(vars.ApplicationYaml); err == nil {
		return vars.ApplicationYaml
	}
	if _, err := os.Stat(vars.ConfAppYml); err == nil {
		return vars.ConfAppYml
	}
	if _, err := os.Stat(vars.ConfAppYaml); err == nil {
		return vars.ConfAppYaml
	}
	if _, err := os.Stat(vars.ConfApplicationYml); err == nil {
		return vars.ConfApplicationYml
	}
	if _, err := os.Stat(vars.ConfApplicationYaml); err == nil {
		return vars.ConfApplicationYml
	}
	return ""
}

// ReadDefault read default configuration into custom
func ReadDefault(custom interface{}) error {
	filename := Filename()
	if filename == "" {
		return confErr
	}
	return Read(filename, custom)
}

// Read filename-configuration into custom
func Read(filename string, custom interface{}) error {
	if custom == nil {
		return errors.New("input is nil")
	}

	if reflect.TypeOf(custom).Kind() != reflect.Ptr {
		return fmt.Errorf("%s is not a pointer", reflect.TypeOf(custom).Name())
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, custom)
}

// NewDefault returns a default configuration
func NewDefault() (*Configuration, error) {
	filename := Filename()
	if filename == "" {
		return nil, confErr
	}
	return New(filename)
}

// New returns *Configuration and error
func New(filename string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)

	// cache bytes
	c.cache = bytes
	return &c, err
}
