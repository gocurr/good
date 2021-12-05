package server

import (
	"errors"
	"fmt"
	"github.com/gocurr/good/consts"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

var serverErr = errors.New("bad server configuration")

// serverMux the global multiplexer
var serverMux *http.ServeMux

// Mux set serverMux
func Mux(mux *http.ServeMux) {
	serverMux = mux
}

// routeFns represents route-fn pairs
var routeFns = make(map[string]func(http.ResponseWriter, *http.Request))

// Route binds route-path to fn
func Route(route string, fn func(http.ResponseWriter, *http.Request)) {
	routeFns[route] = fn
}

// Fire http server entry
func Fire(i interface{}, callbacks ...func()) {
	if i == nil {
		panic(serverErr)
	}

	var c reflect.Value
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Ptr {
		c = reflect.ValueOf(i).Elem()
	} else {
		c = reflect.ValueOf(i)
	}

	var port int64
	switch kind {
	case reflect.Int64, reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8:
		port = c.Int()
	default:
		serverField := c.FieldByName(consts.Server)
		if !serverField.IsValid() {
			panic(serverErr)
		}

		portField := serverField.FieldByName(consts.Port)
		if !portField.IsValid() {
			panic(serverErr)
		}
		// port server bound port
		port = portField.Int()
	}

	if port < 0 || port > 1<<16-1 {
		log.Fatalf("port '%v' is invalid", port)
	} else {
		log.Infof(fmt.Sprintf("http server listening at: [::]: %v", port))
	}

	// invoke callbacks
	for _, callback := range callbacks {
		callback()
	}

	addr := fmt.Sprintf(":%v", port)
	if serverMux != nil {
		muxBoot(addr)
	} else {
		defaultBoot(addr)
	}
}

// defaultBoot boot by http
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
