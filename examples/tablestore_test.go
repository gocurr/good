package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Tablestore(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		log.Error(err)
		return
	}

	_, err = tablestore.New(c)
	if err != nil {
		log.Error(err)
	}
}
