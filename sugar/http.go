package sugar

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// JSONBytes returns JSON []byte from http.Request
func JSONBytes(r *http.Request) ([]byte, error) {
	defer func() { _ = r.Body.Close() }()
	return ioutil.ReadAll(r.Body)
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
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
}

// handleResp handles response
func handleResp(r *http.Response) ([]byte, error) {
	defer func() { _ = r.Body.Close() }()
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(all))
	}
	return all, nil
}

// PostJSON posts JSON format data to the given url
func PostJSON(url string, data interface{}) ([]byte, error) {
	all, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, JSONContentType, bytes.NewReader(all))
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}

// HttpGet calls http.Get via given url, returns bytes of http.Response.Body and reports error
func HttpGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}
