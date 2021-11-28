package good

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var serverMux *http.ServeMux

// Fire http server fire entry
func Fire() {
	if !configured {
		log.Fatalln("Configure the application first!")
	}

	if serverMux == nil {
		log.Fatalln("Set ServerMux first!")
	}

	port := conf.Server.Port
	if port < 0 || port > 1<<16-1 {
		log.Fatalln("Illegal server port!")
	} else {
		log.Infof(fmt.Sprintf("http server listening at [::%v]", port))
	}

	if err := http.ListenAndServe(":"+strconv.Itoa(port), serverMux); err != nil {
		log.Errorf("http server: %v", err)
	}
}

// ServerMux set serverMux
func ServerMux(mux *http.ServeMux) {
	serverMux = mux
}

// Route binds route path to fn
func Route(route string, fn func(http.ResponseWriter, *http.Request)) {
	serverMux.HandleFunc(route, fn)
}
