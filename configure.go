package good

import (
	log "github.com/sirupsen/logrus"
)

// conf the global Configuration
var conf Configuration

// Configure configures the application
func Configure(filename string, fastFail bool) {
	if err := readYml(filename); err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("readYml: %v", err)
		}
	}

	initLogurs()

	initCrontab()

	if err := initTableStore(); err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("initDb: %v", err)
		}
	}

	if err := initDb(); err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("initDb: %v", err)
		}
	}

	if err := initRedis(); err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("initRedis: %v", err)
		}
	}

	if err := initRocketMq(); err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("initRocketMq: %v", err)
		}
	}
}
