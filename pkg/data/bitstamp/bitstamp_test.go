package bitstamp

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestOHLCData_correctStruct(t *testing.T) {
	resJson := `
	{
		"data": {
		  "ohlc": [
			{
			  "close": "1.00000",
			  "high": "1.00001",
			  "low": "1.00000",
			  "open": "1.00001",
			  "timestamp": "1667678400",
			  "volume": "5550.24470"
			},
			{
			  "close": "1.00000",
			  "high": "1.00002",
			  "low": "1.00000",
			  "open": "1.00001",
			  "timestamp": "1667682000",
			  "volume": "25753.43649"
			},
			{
			  "close": "1.00000",
			  "high": "1.00001",
			  "low": "1.00000",
			  "open": "1.00000",
			  "timestamp": "1667685600",
			  "volume": "4275.08018"
			}
		  ],
		  "pair": "USDT/USD"
		}
	  }
	`

	res := map[string]OHLCData{}
	err := json.Unmarshal([]byte(resJson), &res)
	if err != nil {
		t.Fatal(err)
	}

	if res["data"].OHCL[0].Close != 1.00000 {
		t.Error("unexpected value")
	}

	if res["data"].OHCL[0].Volume != 5550.24470 {
		t.Error("unexpected value")
	}
}

type MockHttpClient struct {
	t *testing.T
}

func (mock *MockHttpClient) DoGetRequest(url string, queryParams url.Values) ([]byte, error) {
	expURL := "https://www.bitstamp.net/api/v2/ohlc/usdtusd/"
	expQueryParams := 2

	if url != expURL {
		mock.t.Fatalf("exp: %v, act: %v", expURL, url)
	}

	if len(queryParams) != expQueryParams {
		mock.t.Fatalf("exp: %v, act: %v", expQueryParams, len(queryParams))
	}

	return []byte{}, nil
}

func TestGetOHLC(t *testing.T) {
	client := &BitstampClient{
		httpRequester: &MockHttpClient{
			t: t,
		},
	}

	currency := "usdtusd"
	limit := 1
	step := 3600

	client.GetOHLC(currency, step, limit)
}
