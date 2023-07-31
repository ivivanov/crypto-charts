package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	GET_METHOD = "GET"
)

var (
	ErrUnsuccessfulRequest = errors.New("unsuccessful http request")
)

type HttpRequester interface {
	DoGetRequest(url string, queryParams url.Values) ([]byte, error)
}

type HttpClient struct {
}

func (h *HttpClient) DoGetRequest(url string, queryParams url.Values) ([]byte, error) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	req, err := http.NewRequest(GET_METHOD, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	query := req.URL.Query()
	for k, v := range queryParams {
		query.Add(k, v[0])
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("%w: %v", ErrUnsuccessfulRequest, resp.StatusCode)
	}

	return body, nil
}
