package sugar

import (
	"errors"
	"fmt"
	"github.com/gocurr/good/conf"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var serverErr = errors.New("bad server configuration")

// serverMux the global multiplexer
var serverMux *http.ServeMux

// Fire http server entry
func Fire(c *conf.Configuration, callbacks ...func()) {
	if c == nil || c.Server == nil {
		panic(serverErr)
	}

	// invoke callbacks
	for _, callback := range callbacks {
		callback()
	}

	// port server bound port
	port := c.Server.Port
	if port < 0 || port > 1<<16-1 {
		log.Fatalf("port '%v' is invalid", port)
	} else {
		log.Infof(fmt.Sprintf("http server listening at: [::]: %v", port))
	}

	addr := ":" + strconv.Itoa(port)
	if serverMux != nil {
		muxBoot(addr)
	} else {
		defaultBoot(addr)
	}
}

// defaultBoot boot by default
func defaultBoot(addr string) {
	for route, fn := range routeFns {
		http.HandleFunc(route, fn)
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Errorf("http server: %v", err)
	}
}

// muxBoot boot by serverMux
func muxBoot(addr string) {
	for route, fn := range routeFns {
		serverMux.HandleFunc(route, fn)
	}
	if err := http.ListenAndServe(addr, serverMux); err != nil {
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
