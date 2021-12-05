package conf

import (
	"errors"
	"fmt"
	"github.com/gocurr/good/consts"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

var confErr = errors.New("configuration not found")

// Filename returns a configuration name
func Filename() string {
	if _, err := os.Stat(consts.AppYml); err == nil {
		return consts.AppYml
	}
	if _, err := os.Stat(consts.AppYaml); err == nil {
		return consts.AppYaml
	}
	if _, err := os.Stat(consts.ApplicationYml); err == nil {
		return consts.ApplicationYml
	}
	if _, err := os.Stat(consts.ApplicationYaml); err == nil {
		return consts.ApplicationYaml
	}
	if _, err := os.Stat(consts.ConfAppYml); err == nil {
		return consts.ConfAppYml
	}
	if _, err := os.Stat(consts.ConfAppYaml); err == nil {
		return consts.ConfAppYaml
	}
	if _, err := os.Stat(consts.ConfApplicationYml); err == nil {
		return consts.ConfApplicationYml
	}
	if _, err := os.Stat(consts.ConfApplicationYaml); err == nil {
		return consts.ConfApplicationYml
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
