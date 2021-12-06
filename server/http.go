package server

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// PostJSONRaw calls http.Post to return []byte and error
func PostJSONRaw(url string, in interface{}) ([]byte, error) {
	all, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, consts.JSONContentType, bytes.NewReader(all))
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}

// PostJSON posts JSON format data to the given url, unmarshals body of response into out and reports error
// Assert reflect.TypeOf(out).Kind() == reflect.Ptr
func PostJSON(url string, in interface{}, out interface{}) error {
	raw, err := PostJSONRaw(url, in)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, out)
}

// HttpGetRaw calls http.Get to return []byte and error
func HttpGetRaw(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}

// HttpGetJSON calls http.Get via given url, unmarshals body of response into out and reports error
// Assert reflect.TypeOf(out).Kind() == reflect.Ptr
func HttpGetJSON(url string, out interface{}) error {
	raw, err := HttpGetRaw(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}
