package good

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// port server bound port
var port int

// serverMux the global multiplexer
var serverMux *http.ServeMux

// Fire http server fire entry
func Fire() {
	if !configured {
		log.Fatalln("configure the app first")
	}

	if serverMux == nil {
		log.Info("default handler has been set")
	}

	if port < 0 || port > 1<<16-1 {
		log.Fatalln("illegal server port")
	} else {
		log.Infof(fmt.Sprintf("http server listening at [::%v]", port))
	}

	if serverMux != nil {
		if err := http.ListenAndServe(":"+strconv.Itoa(port), serverMux); err != nil {
			log.Errorf("http server: %v", err)
		}
	} else {
		if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
			log.Errorf("http server: %v", err)
		}
	}
}

// ServerMux set serverMux
func ServerMux(mux *http.ServeMux) {
	serverMux = mux
}

// Route binds route path to fn
func Route(route string, fn func(http.ResponseWriter, *http.Request)) {
	if serverMux != nil {
		serverMux.HandleFunc(route, fn)
	} else {
		http.HandleFunc(route, fn)
	}
}
