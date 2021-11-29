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
func Fire(callbacks ...func()) {
	if !configured {
		tryConfig()
	}

	// invoke callbacks
	for _, callback := range callbacks {
		callback()
	}

	if serverMux == nil {
		log.Info("default handler has been set")
	}

	if port < 0 || port > 1<<16-1 {
		log.Fatalln("illegal server port")
	} else {
		log.Infof(fmt.Sprintf("http server listening at: [::]: %v", port))
	}

	if serverMux != nil {
		muxboot()
	} else {
		defaultBoot()
	}
}

// defaultBoot boot by default
func defaultBoot() {
	for route, fn := range routeFns {
		http.HandleFunc(route, fn)
	}
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Errorf("http server: %v", err)
	}
}

// muxboot boot by serverMux
func muxboot() {
	for route, fn := range routeFns {
		serverMux.HandleFunc(route, fn)
	}
	if err := http.ListenAndServe(":"+strconv.Itoa(port), serverMux); err != nil {
		log.Errorf("http server: %v", err)
	}
}

// ServerMux set serverMux
func ServerMux(mux *http.ServeMux) {
	serverMux = mux
}

// routeFns represents route-fn pairs
var routeFns = make(map[string]func(http.ResponseWriter, *http.Request))

// Route binds route path to fn
func Route(route string, fn func(http.ResponseWriter, *http.Request)) {
	routeFns[route] = fn
}
