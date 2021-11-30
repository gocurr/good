package conf

import (
	"errors"
	"fmt"
	"reflect"
)

// Configuration represents a yaml configuration
type Configuration struct {
	Server *struct {
		Port int `yaml:"port"`
	}

	Logrus *struct {
		Format  string `yaml:"format"`
		TTY     bool   `yaml:"tty"`
		GrayLog *struct {
			Enable bool                   `yaml:"enable"`
			Host   string                 `yaml:"host"`
			Port   int                    `yaml:"port"`
			Extra  map[string]interface{} `yaml:"extra"`
		} `yaml:"graylog"`
	}

	Oracle *struct {
		Driver     string `yaml:"driver"`
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Datasource string `yaml:"datasource"`
	}

	Mysql *struct {
		Driver     string `yaml:"driver"`
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Datasource string `yaml:"datasource"`
	}

	Redis *struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}

	RocketMq *struct {
		Addr      []string `yaml:"addr"`
		Retry     int      `yaml:"retry"`
		AccessKey string   `yaml:"access-key"`
		SecretKey string   `yaml:"secret-key"`
	} `yaml:"rocket-mq"`

	TableStore *struct {
		EndPoint        string `yaml:"end-point"`
		InstanceName    string `yaml:"instance-name"`
		AccessKeyId     string `yaml:"access-key-id"`
		AccessKeySecret string `yaml:"access-key-secret"`
	} `yaml:"table-store"`

	Crontab map[string]struct {
		Spec string `yaml:"spec"`
	}

	Secure *struct {
		Key string `yaml:"key"`
	}

	Custom map[string]interface{} // custom field for users
}

// String return string field in custom
func (c *Configuration) String(field string) string {
	i := c.Custom[field]
	if reflect.TypeOf(i).Kind() == reflect.String {
		return i.(string)
	}
	return fmt.Sprintf("%v", i)
}

// Int return int field in custom
func (c *Configuration) Int(field string) (int, error) {
	i := c.Custom[field]
	if reflect.TypeOf(i).Kind() == reflect.Int {
		return i.(int), nil
	}
	return 0, errors.New(fmt.Sprintf("%v is not 'int' type", i))
}

// Float64 return float64 field in custom
func (c *Configuration) Float64(field string) (float64, error) {
	i := c.Custom[field]
	if reflect.TypeOf(i).Kind() == reflect.Float64 {
		return i.(float64), nil
	}
	return 0, errors.New(fmt.Sprintf("%v is not 'float64' type", i))
}

// Float32 return float32 field in custom
func (c *Configuration) Float32(field string) (float32, error) {
	i := c.Custom[field]
	if reflect.TypeOf(i).Kind() == reflect.Float32 {
		return i.(float32), nil
	}
	return 0, errors.New(fmt.Sprintf("%v is not 'float32' type", i))
}

// Float return float64 field in custom
func (c *Configuration) Float(field string) (float64, error) {
	return c.Float64(field)
}

// Interface return interface{} field in custom
func (c *Configuration) Interface(field string) interface{} {
	return c.Custom[field]
}

// Slice return slice field in custom
func (c *Configuration) Slice(field string) ([]interface{}, error) {
	i := c.Custom[field]
	if reflect.TypeOf(i).Kind() == reflect.Slice {
		return i.([]interface{}), nil
	}
	return nil, errors.New(fmt.Sprintf("%v is not '[]interface{}' type", i))
}

// Map return map field in custom
func (c *Configuration) Map(field string) (map[string]interface{}, error) {
	i := c.Custom[field]
	if reflect.TypeOf(i).Kind() == reflect.Map {
		return i.(map[string]interface{}), nil
	}
	return nil, errors.New(fmt.Sprintf("%v is not 'map[string]interface{}' type", i))
}
