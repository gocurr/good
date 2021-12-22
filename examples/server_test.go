package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/server"
	"net/http"
	"testing"
)

func Test_Server(t *testing.T) {
	_, err := conf.New("../app.yaml")
	Panic(err)

	type Msg struct {
		Text string
	}
	server.Route("/", func(w http.ResponseWriter, r *http.Request) {
		server.JSONHeader(w)

		msg := Msg{Text: "hello, better"}
		if _, err := server.WriteJSON(w, &msg); err != nil {
			panic(err)
		}
	})

	//server.Fire(c)
}
