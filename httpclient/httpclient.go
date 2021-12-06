package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocurr/good/consts"
	"io/ioutil"
	"net/http"
)

// handleResp handles response to return []byte and error
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

// PostJSON posts JSON format data to the given url
// Unmarshal body of response into out and reports error
// Assert out is a pointer
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

// HttpGetJSON calls http.Get via given url
// Unmarshal body of response into out and reports error
// Assert out is a pointer
func HttpGetJSON(url string, out interface{}) error {
	raw, err := HttpGetRaw(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}
