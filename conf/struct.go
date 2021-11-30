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

	Custom map[string]interface{} // custom field for users
}

// String return string field in custom
func (c *Configuration) String(field string) string {
	return c.Custom[field].(string)
}

// Int return int field in custom
func (c *Configuration) Int(field string) int {
	return c.Custom[field].(int)
}

// Float return float64 field in custom
func (c *Configuration) Float(field string) float64 {
	return c.Custom[field].(float64)
}

// Slice return slice field in custom
func (c *Configuration) Slice(field string) []interface{} {
	return c.Custom[field].([]interface{})
}

// Interface return interface{} field in custom
func (c *Configuration) Interface(field string) interface{} {
	return c.Custom[field]
}

// Map return map field in custom
func (c *Configuration) Map(field string) map[string]interface{} {
	return c.Custom[field].(map[string]interface{})
}
