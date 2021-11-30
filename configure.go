package good

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/mysql"
	"github.com/gocurr/good/oracle"
	"github.com/gocurr/good/redis"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

// custom represents the same filed in configuration
var custom = make(map[string]interface{})

// reports Configure has been invoked
var configured bool

// ConfigDefault config by default file
func ConfigDefault() {
	tryConfig()
}

// Configure configures the application
func Configure(filename string, fastFail bool) {
	// tag configured
	configured = true

	c, err := conf.Read(filename)
	if err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("conf: %v", err)
		}
	}

	if c.Logrus != nil {
		if err := logger.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("logrus: %v", err)
			}
		}

	}

	if len(c.Crontab) > 0 {
		crontab.Init(c)
	}

	if c.TableStore != nil {
		if err := tablestore.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("tablestore: %v", err)
			}
		}
	}

	if c.Oracle != nil {
		if err := oracle.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("oracle: %v", err)
			}
		}
	}

	if c.Mysql != nil {
		if err := mysql.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("mysql: %v", err)
			}
		}
	}

	if c.Redis != nil {
		if err := redis.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("redis: %v", err)
			}
		}
	}

	if c.RocketMq != nil {
		if err := rocketmq.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("rocketmq: %v", err)
			}
		}
	}

	// set custom field
	if len(c.Custom) > 0 {
		custom = c.Custom
	}
}

// tryOnce for tryConfig
var tryOnce sync.Once

// tryConfig try to configure once more
func tryConfig() {
	tryOnce.Do(func() {
		f := filename()
		if f == "" {
			log.Fatalln("cannot find config file")
		}
		Configure(f, false)
		log.Infof("app is configured by '%s'", f)
	})
}

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

// filename returns a configuration name
func filename() string {
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
