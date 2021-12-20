package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocurr/good/consts"
	"io/ioutil"
	"net/http"
)

// handleResp handles http-response
// to return byte slice and reports error
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

// PostJSONRaw calls http.Post
// to return byte slice and reports error
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

// PostJSON posts JSON format data to the given url,
// then unmarshal body of response into out and reports error
//
// Assert out is a pointer
func PostJSON(url string, in interface{}, out interface{}) error {
	raw, err := PostJSONRaw(url, in)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, out)
}

// GetRaw calls http.Get to the given url
// to return byte slice and reports error
func GetRaw(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return handleResp(response)
}

// GetJSON calls http.Get to the given url,
// then unmarshal body of response into out and reports error
//
// Assert out is a pointer
func GetJSON(url string, out interface{}) error {
	raw, err := GetRaw(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}
