package osmosis

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/ivivanov/crypto-charts/pkg/fetchers"
)

const (
	// 	BASE_ADDRESS           = "https://api.osmosis.zone"
	BASE_ADDRESS = "https://api-osmosis.imperator.co"
)

type OHLC struct {
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
	Time  int64   `json:"time"`
}

type OsmosisClient struct {
	httpRequester fetchers.HttpRequester
	retry         int
}

func NewOsmosisClient() *OsmosisClient {
	return &OsmosisClient{
		httpRequester: fetchers.NewHttpClient(3 * time.Second),
		retry:         3,
	}
}

// GetOHLC return array OHLC data.
func (oc *OsmosisClient) GetOHLC(symbol string, timeFrame int) ([]OHLC, error) {
	ohlcEndpoint := fmt.Sprintf("%v/tokens/v2/historical/%v/chart", BASE_ADDRESS, symbol)
	queryParams := url.Values{}
	queryParams.Set("tf", fmt.Sprint(timeFrame))

	var bodyJson []byte
	var err error
	r := 0
	for  {
		if bodyJson, err = oc.httpRequester.DoGetRequest(ohlcEndpoint, queryParams); err == nil {
			break
		} else if r == oc.retry {
			return nil, fmt.Errorf("osmosis error fetching OHLC for %v: %w", symbol, err)
		}

		r++
	}

	res := []OHLC{}
	err = json.Unmarshal(bodyJson, &res)
	if err != nil {
		return nil, fmt.Errorf("osmosis error unmarshal response: %w", err)
	}

	return res, nil
}
