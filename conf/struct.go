package conf

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

	Custom map[string]interface{}
}
