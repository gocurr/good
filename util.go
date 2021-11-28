package good

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocurr/good/crypto"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

// DefaultTimeFormat default time format
const DefaultTimeFormat = "2006-01-02 15:04:05"
const DefaultContentType = "Content-Type:application/json"

// GenPasswd generates encrypted string via pw
func GenPasswd(pw string) string {
	var secret string
	filename := "secret.txt"
	all, err := ioutil.ReadFile(filename)
	if err != nil {
		secret = crypto.CreateSecret()
		fmt.Println(secret)
		if err := ioutil.WriteFile(filename, []byte(secret), os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		secret = string(all)
	}

	encrypter, err := crypto.Encrypt(secret, pw)
	if err != nil {
		panic(err)
	}
	return encrypter
}

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
func JSONBytes(r *http.Request) []byte {
	body := r.Body
	defer func() { _ = body.Close() }()

	all, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}
	return all
}

// PostJSON posts JSON format data to the given url
func PostJSON(url string, data []interface{}) []byte {
	all, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	response, err := http.Post(url, DefaultContentType, bytes.NewReader(all))
	if err != nil {
		return nil
	}
	return handleResp(response)
}

// HttpGet get []byte via given url
func HttpGet(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	return handleResp(response)
}

// handleResp handle response
func handleResp(response *http.Response) []byte {
	body := response.Body
	defer func() { _ = body.Close() }()

	respBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}
	return respBytes
}
