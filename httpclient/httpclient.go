package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gocurr/good/consts"
	"io/ioutil"
	"net/http"
	"time"
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
func PostJSONRaw(url string, in interface{}, timeout ...time.Duration) ([]byte, error) {
	all, errMarshal := json.Marshal(in)
	if errMarshal != nil {
		return nil, errMarshal
	}
	body := bytes.NewReader(all)

	var err error
	var resp *http.Response
	if len(timeout) > 0 {
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout[0])
		defer cancelFunc()

		var req *http.Request
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, body)
		if err != nil {
			return nil, err
		}
		resp, err = http.DefaultClient.Do(req)
	} else {
		resp, err = http.Post(url, consts.JSONContentType, body)
	}
	if err != nil {
		return nil, err
	}

	return handleResp(resp)
}

// PostJSON posts JSON format data to the given url,
// then unmarshal body of response into out and reports error
//
// Assert out is a pointer
func PostJSON(url string, in interface{}, out interface{}, timeout ...time.Duration) error {
	raw, err := PostJSONRaw(url, in, timeout...)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, out)
}

// GetRaw calls http.Get to the given url
// to return byte slice and reports error
func GetRaw(url string, timeout ...time.Duration) ([]byte, error) {
	var err error
	var resp *http.Response
	if len(timeout) > 0 {
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout[0])
		defer cancelFunc()

		var req *http.Request
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		resp, err = http.DefaultClient.Do(req)
	} else {
		resp, err = http.Get(url)
	}
	if err != nil {
		return nil, err
	}

	return handleResp(resp)
}

// GetJSON calls http.Get to the given url,
// then unmarshal body of response into out and reports error
//
// Assert out is a pointer
func GetJSON(url string, out interface{}, timeout ...time.Duration) error {
	raw, err := GetRaw(url, timeout...)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}
