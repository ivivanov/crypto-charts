package bitstamp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/ivivanov/crypto-charts/pkg/data"
)

const (
	BASE_ADDRESS = "https://www.bitstamp.net/api"
)

var (
	ErrUnmarshalErrorResponse = errors.New("unmarshal bitstamp error response")
)

type ErrorResponse struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}

type OHLCData struct {
	OHCL []OHLC `json:"ohlc"`
	Pair string `json:"pair"`
}

type OHLC struct {
	Close     float64 `json:"close,string"`
	High      float64 `json:"high,string"`
	Low       float64 `json:"low,string"`
	Open      float64 `json:"open,string"`
	Volume    float64 `json:"volume,string"`
	Timestamp string  `json:"timestamp"`
}

type BitstampClient struct {
	httpRequester data.HttpRequester
}

func NewBitstampClient() *BitstampClient {
	return &BitstampClient{
		httpRequester: &data.HttpClient{},
	}
}

// GetOHLC return array OHLC data.
func (bc *BitstampClient) GetOHLC(currencyPair string, step, limit int) ([]OHLC, error) {
	if currencyPair == "" || step == 0 {
		return nil, fmt.Errorf("all args are required")
	}

	ohlcEndpoint := fmt.Sprintf("%v/v2/ohlc/%s/", BASE_ADDRESS, currencyPair)
	queryParams := url.Values{}
	queryParams.Set("step", fmt.Sprint(step))
	queryParams.Set("limit", fmt.Sprint(limit))

	bodyJson, err := bc.httpRequester.DoGetRequest(ohlcEndpoint, queryParams)
	if err != nil {
		body := &ErrorResponse{}

		err = json.Unmarshal(bodyJson, body)
		if err != nil {
			return nil, fmt.Errorf("%w: %w, pair: %v, json: %v", ErrUnmarshalErrorResponse, err, currencyPair, bodyJson)
		}

		return nil, fmt.Errorf("error: %s, reason: %s, status: %s", body.Code, body.Reason, body.Status)
	}

	res := map[string]OHLCData{}
	err = json.Unmarshal(bodyJson, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal bitstamp response: %w", err)
	}

	return res["data"].OHCL, nil
}
