package good

import (
	log "github.com/sirupsen/logrus"
)

var conf Configuration

func Configure(file string, fastFail bool) {
	if err := read(file); err != nil {
		if fastFail {
			panic(err)
		} else {
			log.Errorf("read: %v", err)
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

	initTableStore()

	initLogurs()
}
