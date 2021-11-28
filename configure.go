package good

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/db"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/redis"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
)

// custom represents the same filed in configuration
var custom map[string]interface{}

// reports Configure has been invoked
var configured bool

// Configure configures the application
func Configure(filename string, fastFail bool) {
	// tag configured
	configured = true

	c, err := conf.ReadYml(filename)
	if err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("readYml: %v", err)
		}
	}

	if c.Logrus != nil {
		logger.Init(c)
	}

	if len(c.Crontab) > 0 {
		crontab.Init(c)
	}

	if c.TableStore != nil {
		if err := tablestore.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("initDb: %v", err)
			}
		}
	}

	if c.DB != nil {
		if err := db.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("initDb: %v", err)
			}
		}
	}

	if c.Redis != nil {
		if err := redis.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("initRedis: %v", err)
			}
		}
	}

	if c.RocketMq != nil {
		if err := rocketmq.Init(c); err != nil {
			if fastFail {
				panic(err)
			} else {
				log.Errorf("initRocketMq: %v", err)
			}
		}
	}

	// set server bound port
	port = c.Server.Port
	// set custom field
	custom = c.Custom
}
