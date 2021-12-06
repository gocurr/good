package server

import (
	"encoding/json"
	"github.com/gocurr/good/consts"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// Parameter returns string via name from http.Request
func Parameter(name string, r *http.Request) string {
	values := r.URL.Query()
	v, ok := values[name]
	if ok {
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

// Parameters returns []string via name from http.Request
func Parameters(name string, r *http.Request) []string {
	values := r.URL.Query()
	v, ok := values[name]
	if ok {
		return v
	}
	return nil
}

// JSON unmarshals body of http.Request into out
// Assert reflect.TypeOf(out).Kind() == reflect.Ptr
func JSON(r *http.Request, out interface{}) error {
	defer func() { _ = r.Body.Close() }()
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(all, out)
}

// ErrMsg error message
type ErrMsg struct {
	Err string `json:"err"`
}

// HandleErr handles error
// Assert err is none-nil or else it will PANIC
func HandleErr(err error, w http.ResponseWriter, status ...int) {
	if len(status) > 0 {
		w.WriteHeader(status[0])
	}
	log.Errorf("%v", err)
	if msg, err := json.Marshal(ErrMsg{Err: err.Error()}); err == nil {
		_, _ = w.Write(msg)
	}
}

// JSONHeader adds JSON to response headers
func JSONHeader(w http.ResponseWriter) {
	w.Header().Add(consts.ContentType, consts.ApplicationJSON)
}
