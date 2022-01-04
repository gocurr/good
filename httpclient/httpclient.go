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

// handleResp returns byte-slice and reports error encountered
// by the given http-response.
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

// PostJSONRawTimeout marshals the input value and issues a POST to the specified URL,
// then it'll return byte-slice and report error encountered.
//
// Cancels the process when timeout.
func PostJSONRawTimeout(url string, in interface{}, timeout time.Duration) ([]byte, error) {
	all, errMarshal := json.Marshal(in)
	if errMarshal != nil {
		return nil, errMarshal
	}
	body := bytes.NewReader(all)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return handleResp(resp)
}

// PostJSONRaw marshals the input value and issues a POST to the specified URL,
// then it'll return byte-slice and report error encountered.
func PostJSONRaw(url string, in interface{}) ([]byte, error) {
	all, errMarshal := json.Marshal(in)
	if errMarshal != nil {
		return nil, errMarshal
	}
	body := bytes.NewReader(all)

	resp, err := http.Post(url, consts.JSONContentType, body)
	if err != nil {
		return nil, err
	}

	return handleResp(resp)
}

// PostJSON marshals the input value and issues a POST to the specified URL,
// then it'll unmarshal response into out(must be a pointer) and report error encountered.
//
// The timeout parameter is optional.
func PostJSON(url string, in interface{}, out interface{}) error {
	raw, err := PostJSONRaw(url, in)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, out)
}

// PostJSONTimeout marshals the input value and issues a POST to the specified URL,
// then it'll unmarshal response into out(must be a pointer) and report error encountered.
//
// Cancels the process when timeout.
func PostJSONTimeout(url string, in interface{}, out interface{}, timeout time.Duration) error {
	raw, err := PostJSONRawTimeout(url, in, timeout)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, out)
}

// GetRaw issues a GET to the specified URL,
// then it'll return byte-slice and report error encountered.
func GetRaw(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return handleResp(resp)
}

// GetRawTimeout issues a GET to the specified URL,
// then it'll return byte-slice and report error encountered.
//
// Cancels the process when timeout.
func GetRawTimeout(url string, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return handleResp(resp)
}

// GetJSON issues a GET to the specified URL,
// then it'll unmarshal response into out(must be a pointer) and report error encountered.
func GetJSON(url string, out interface{}) error {
	raw, err := GetRaw(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}

// GetJSONTimeout issues a GET to the specified URL,
// then it'll unmarshal response into out(must be a pointer) and report error encountered.
//
// Cancels the process when timeout.
func GetJSONTimeout(url string, out interface{}, timeout time.Duration) error {
	raw, err := GetRawTimeout(url, timeout)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}
