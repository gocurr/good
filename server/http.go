package server

import (
	"encoding/json"
	"github.com/gocurr/good/consts"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// Parameter returns a string parameter in URL
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

// Parameters returns string slice in URL
func Parameters(name string, r *http.Request) []string {
	values := r.URL.Query()
	v, ok := values[name]
	if ok {
		return v
	}
	return nil
}

// JSON unmarshal body of http.Request into out.
//
// Assert out is a pointer or else it will PANIC
func JSON(r *http.Request, out interface{}) error {
	defer func() { _ = r.Body.Close() }()
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(all, out)
}

// ErrMsg represents an error message
type ErrMsg struct {
	Err string `json:"err"`
}

// HandleErr handles error
//
// Assert err is non-nil or else it will PANIC
func HandleErr(err error, w http.ResponseWriter, status ...int) {
	if len(status) > 0 {
		w.WriteHeader(status[0])
	}
	log.Errorf("%v", err)
	if msg, err := json.Marshal(ErrMsg{Err: err.Error()}); err == nil {
		_, _ = w.Write(msg)
	}
}

// JSONHeader sets response header content-type:"application/json; charset=UTF-8"
func JSONHeader(w http.ResponseWriter) {
	w.Header().Set(consts.ContentType, consts.ApplicationJSON)
}

// WriteJSON marshals data and write to the connection
func WriteJSON(w http.ResponseWriter, data interface{}) (int, error) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	return w.Write(marshal)
}
