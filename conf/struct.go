package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
)

// Configuration represents a yaml configuration
type Configuration struct {
	cache []byte `yaml:"-"` // cached yaml-bytes

	Server struct {
		Port int `yaml:"port,omitempty"`
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

	Oracle struct {
		Driver     string `yaml:"driver,omitempty"`
		User       string `yaml:"user,omitempty"`
		Password   string `yaml:"password,omitempty"`
		Datasource string `yaml:"datasource,omitempty"`
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
	}

	RocketMq struct {
		Addr      []string `yaml:"addr,omitempty"`
		Retry     int      `yaml:"retry,omitempty"`
		AccessKey string   `yaml:"access-key,omitempty"`
		SecretKey string   `yaml:"secret-key,omitempty"`
	} `yaml:"rocket-mq,omitempty"`

	TableStore struct {
		EndPoint        string `yaml:"end-point,omitempty"`
		InstanceName    string `yaml:"instance-name,omitempty"`
		AccessKeyId     string `yaml:"access-key-id,omitempty"`
		AccessKeySecret string `yaml:"access-key-secret,omitempty"`
	} `yaml:"table-store,omitempty"`

	Crontab map[string]string `yaml:"crontab,omitempty"`

	Secure struct {
		Key string `yaml:"key,omitempty"`
	}

	Reserved map[string]interface{} `yaml:"reserved,omitempty"` // reserved area for users
}

// Fill fills custom struct
func (c *Configuration) Fill(custom interface{}) error {
	if reflect.TypeOf(custom).Kind() != reflect.Ptr {
		return fmt.Errorf("%s is not a pointer", reflect.TypeOf(custom).Name())
	}
	return yaml.Unmarshal(c.cache, custom)
}

// ReservedString return string field in reserved
func (c *Configuration) ReservedString(field string) (string, error) {
	i := c.Reserved[field]
	if i == nil {
		return "", fmt.Errorf("'%s' not found in configuration", field)
	}
	if reflect.TypeOf(i).Kind() == reflect.String {
		return i.(string), nil
	}

	return fmt.Sprintf("%v", i), nil
}

// ReservedInt return int field in reserved
func (c *Configuration) ReservedInt(field string) (int, error) {
	i := c.Reserved[field]
	if i == nil {
		return 0, fmt.Errorf("'%s' not found in configuration", field)
	}
	if reflect.TypeOf(i).Kind() == reflect.Int {
		return i.(int), nil
	}
	return 0, fmt.Errorf("%v is not 'int' type", i)
}

// ReservedFloat64 return float64 field in reserved
func (c *Configuration) ReservedFloat64(field string) (float64, error) {
	i := c.Reserved[field]
	if i == nil {
		return 0, fmt.Errorf("'%s' not found in configuration", field)
	}
	if reflect.TypeOf(i).Kind() == reflect.Float64 {
		return i.(float64), nil
	}
	return 0, fmt.Errorf("%v is not 'float64' type", i)
}

// ReservedFloat32 return float32 field in reserved
func (c *Configuration) ReservedFloat32(field string) (float32, error) {
	i := c.Reserved[field]
	if i == nil {
		return 0, fmt.Errorf("'%s' not found in configuration", field)
	}
	if reflect.TypeOf(i).Kind() == reflect.Float32 {
		return i.(float32), nil
	}
	return 0, fmt.Errorf("%v is not 'float32' type", i)
}

// ReservedFloat return float64 field in reserved
func (c *Configuration) ReservedFloat(field string) (float64, error) {
	return c.ReservedFloat64(field)
}

// ReservedInterface return interface{} field in reserved
func (c *Configuration) ReservedInterface(field string) interface{} {
	return c.Reserved[field]
}

// ReservedSlice return slice field in reserved
func (c *Configuration) ReservedSlice(field string) ([]interface{}, error) {
	i := c.Reserved[field]
	if i == nil {
		return nil, fmt.Errorf("'%s' not found in configuration", field)
	}
	if reflect.TypeOf(i).Kind() == reflect.Slice {
		return i.([]interface{}), nil
	}
	return nil, fmt.Errorf("%v is not '[]interface{}' type", i)
}

// ReservedMap return map field in reserved
func (c *Configuration) ReservedMap(field string) (map[string]interface{}, error) {
	i := c.Reserved[field]
	if i == nil {
		return nil, fmt.Errorf("'%s' not found in configuration", field)
	}
	if reflect.TypeOf(i).Kind() == reflect.Map {
		return i.(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("%v is not 'map[string]interface{}' type", i)
}
