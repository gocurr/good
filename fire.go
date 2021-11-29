package good

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	appYml         = "app.yml"
	applicationYml = "application.yml"
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

// tryConfig try to configure once more
func tryConfig() {
	filename := confFile()
	if filename == "" {
		log.Fatalln("cannot find config file")
	}
	Configure(filename, false)
	log.Warnf("app is configured by [%s]", filename)
}

// confFile returns a *.yml or *.yaml filename
func confFile() string {
	if _, err := os.Stat(appYml); err == nil {
		return appYml
	}
	if _, err := os.Stat(applicationYml); err == nil {
		return applicationYml
	}

	var filename string
	if err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		name := info.Name()
		if strings.HasSuffix(name, ".yml") ||
			strings.HasSuffix(name, ".yaml") {
			filename = path
		}
		return nil
	}); err != nil {
		return ""
	}

	return filename
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
