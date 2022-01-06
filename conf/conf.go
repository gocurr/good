package conf

import (
	"errors"
	"github.com/gocurr/good/consts"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

var errConf = errors.New("conf: configuration not found")
var errFilename = errors.New("conf: filename must be non-empty")

// Filename returns a builtin configuration filename.
func Filename() string {
	if _, err := os.Stat(consts.AppYaml); err == nil {
		return consts.AppYaml
	}
	if _, err := os.Stat(consts.AppYml); err == nil {
		return consts.AppYml
	}
	if _, err := os.Stat(consts.ApplicationYaml); err == nil {
		return consts.ApplicationYaml
	}
	if _, err := os.Stat(consts.ApplicationYml); err == nil {
		return consts.ApplicationYml
	}
	if _, err := os.Stat(consts.ConfAppYaml); err == nil {
		return consts.ConfAppYaml
	}
	if _, err := os.Stat(consts.ConfAppYml); err == nil {
		return consts.ConfAppYml
	}
	if _, err := os.Stat(consts.ConfApplicationYaml); err == nil {
		return consts.ConfApplicationYaml
	}
	if _, err := os.Stat(consts.ConfApplicationYml); err == nil {
		return consts.ConfApplicationYml
	}
	return ""
}

// ReadDefault reads default configuration into custom.
//
// Note: custom must be a pointer.
func ReadDefault(custom interface{}) error {
	filename := Filename()
	if filename == "" {
		return errConf
	}
	return Read(filename, custom)
}

// Read reads configuration into custom by the given filename.
//
// Note: custom must be a pointer.
func Read(filename string, custom interface{}) error {
	if filename == "" {
		return errFilename
	}

	if custom == nil {
		return errors.New("conf: custom must be non-nil")
	}

	if reflect.TypeOf(custom).Kind() != reflect.Ptr {
		return errors.New("conf: custom must be a pointer")
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, custom)
}

// NewDefault returns default configuration.
func NewDefault() (*Configuration, error) {
	filename := Filename()
	if filename == "" {
		return nil, errConf
	}
	return New(filename)
}

// New returns a configuration by the given filename.
func New(filename string) (*Configuration, error) {
	if filename == "" {
		return nil, errFilename
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c Configuration
	err = yaml.Unmarshal(file, &c)

	// cache file bytes
	c.cache = file
	return &c, err
}
