package conf

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

// default configuration names
const (
	appYml  = "app.yml"
	appYaml = "app.yaml"

	applicationYml  = "application.yml"
	applicationYaml = "application.yaml"

	confAppYml  = "conf/app.yml"
	confAppYaml = "conf/app.yaml"

	confApplicationYml  = "conf/application.yml"
	confApplicationYaml = "conf/application.yaml"
)

// Filename returns a configuration name
func Filename() string {
	if _, err := os.Stat(appYml); err == nil {
		return appYml
	}
	if _, err := os.Stat(appYaml); err == nil {
		return appYaml
	}
	if _, err := os.Stat(applicationYml); err == nil {
		return applicationYml
	}
	if _, err := os.Stat(applicationYaml); err == nil {
		return applicationYaml
	}
	if _, err := os.Stat(confAppYml); err == nil {
		return confAppYml
	}
	if _, err := os.Stat(confAppYaml); err == nil {
		return confAppYaml
	}
	if _, err := os.Stat(confApplicationYml); err == nil {
		return confApplicationYml
	}
	if _, err := os.Stat(confApplicationYaml); err == nil {
		return confApplicationYaml
	}
	return ""
}

// Read returns *Configuration and error
func Read(filename string, raw ...interface{}) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(raw) > 0 {
		if reflect.TypeOf(raw[0]).Kind() == reflect.Ptr {
			err = yaml.Unmarshal(bytes, raw[0])
			return nil, err
		}
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)

	// cache bytes
	c.cache = bytes
	return &c, err
}

var confErr = errors.New("configuration not found")

// ReadDefault reads default configurations
func ReadDefault(raw ...interface{}) (*Configuration, error) {
	filename := Filename()
	if filename == "" {
		return nil, confErr
	}
	return Read(filename, raw...)
}
