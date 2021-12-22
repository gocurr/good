package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/tablestore"
	"testing"
)

func Test_Tablestore(t *testing.T) {
	c, err := conf.New("../app.yaml")
	Panic(err)

	_, err = tablestore.New(c)
	Panic(err)
}
