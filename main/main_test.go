package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"testing"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type Custom struct {
	Server struct {
		Port int `yaml:"port"`
	}

	Logrus struct {
		Format  string `yaml:"format,omitempty"`
		TTY     bool   `yaml:"tty,omitempty"`
		GrayLog struct {
			Enable bool                   `yaml:"enable,omitempty"`
			Host   string                 `yaml:"host,omitempty"`
			Port   int                    `yaml:"port,omitempty"`
			Extra  map[string]interface{} `yaml:"extra,omitempty"`
		} `yaml:"graylog,omitempty"`
	}

	Mysql struct {
		Driver     string `yaml:"driver,omitempty"`
		User       string `yaml:"user,omitempty"`
		Password   string `yaml:"password,omitempty"`
		Datasource string `yaml:"datasource,omitempty"`
	}

	Redis struct {
		Host     string `yaml:"host,omitempty"`
		Port     int    `yaml:"port,omitempty"`
		Password string `yaml:"password,omitempty"`
		DB       int    `yaml:"db,omitempty"`
	} `yaml:"redis"`

	Crontab map[string]string `yaml:"crontab,omitempty"`

	Secure struct {
		Key string `yaml:"key,omitempty"`
	}

	Pg struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"pg"`

	Maria struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"maria"`

	Logic struct {
		Api   string   `yaml:"api"`
		Names []string `yaml:"names"`
	}
}

func Test_Main(t *testing.T) {
	var c Custom
	err := conf.Read("../app.yaml", &c)
	Panic(err)
	err = logger.Set(&c)
	Panic(err)
}
