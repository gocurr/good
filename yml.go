package good

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Configuration represents a yaml configuration
type Configuration struct {
	Server struct {
		Port int `yaml:"port"`
	}

	Logrus *struct {
		TTY     bool `yaml:"tty"`
		GrayLog struct {
			Host  string                 `yaml:"host"`
			Port  int                    `yaml:"port"`
			Extra map[string]interface{} `yaml:"extra"`
		} `yaml:"graylog"`
	}

	DB *struct {
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

	Secure struct {
		Key string `yaml:"key"`
	}
}

// readYml yml-file to conf
func readYml(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, &conf)
}
