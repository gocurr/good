package sugar

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
)

// DefaultTimeFormat default time format
const DefaultTimeFormat = "2006-01-02 15:04:05"

// DefaultContentType default Content-Type
const DefaultContentType = "Content-Type:application/json"

// handleFloat handles float with fn
func handleFloat(f, n float64, fn func(float64) float64) float64 {
	mul := math.Pow(10, n)
	return fn(f*mul) / mul
}

// RoundFloat return rounded float64
func RoundFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Round)
}

// CeilFloat return rounded float64
func CeilFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Ceil)
}

// FloorFloat return rounded float64
func FloorFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Floor)
}

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

// handleResp handle response
func handleResp(r *http.Response) ([]byte, error) {
	defer func() { _ = r.Body.Close() }()
	return ioutil.ReadAll(r.Body)
}

// PostJSON posts JSON format data to the given url
func PostJSON(url string, data interface{}) ([]byte, error) {
	all, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, DefaultContentType, bytes.NewReader(all))
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}

// HttpGet get []byte via given url
func HttpGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}

// ErrMsg error message
type ErrMsg struct {
	Err string `json:"err"`
}

// HandleErr handle error
func HandleErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	log.Errorf("%v", err)
	if msg, err := json.Marshal(ErrMsg{Err: fmt.Sprintf("%v", err)}); err == nil {
		_, _ = w.Write(msg)
	}
}

// JSONHeader add JSON to response headers
func JSONHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
}
