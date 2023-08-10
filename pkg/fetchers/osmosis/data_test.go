package osmosis

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestOHLCData_correctStruct(t *testing.T) {
	resJson := `
	[
		{
		  "time": 1690725600,
		  "close": 27.18943631,
		  "high": 27.18943631,
		  "low": 26.97994062,
		  "open": 26.97994062
		},
		{
		  "time": 1690729200,
		  "close": 27.19212248,
		  "high": 27.19212248,
		  "low": 27.18932529,
		  "open": 27.18943631
		},
		{
		  "time": 1690732800,
		  "close": 27.1077099,
		  "high": 27.2016217,
		  "low": 27.10591602,
		  "open": 27.19212248
		}
	]
	`

	res := []OHLC{}
	err := json.Unmarshal([]byte(resJson), &res)
	if err != nil {
		t.Fatal(err)
	}

	if res[0].Close != 27.18943631 {
		t.Error("unexpected value")
	}

	if res[0].Time != 1690725600 {
		t.Error("unexpected value")
	}
}

type MockHttpClient struct {
	t *testing.T
}

func (mock *MockHttpClient) DoGetRequest(url string, queryParams url.Values) ([]byte, error) {
	expURL := "https://api-osmosis.imperator.co/tokens/v2/historical/NLS/chart"
	expQueryParams := 1

	if url != expURL {
		mock.t.Fatalf("exp: %v, act: %v", expURL, url)
	}

	if len(queryParams) != expQueryParams {
		mock.t.Fatalf("exp: %v, act: %v", expQueryParams, len(queryParams))
	}

	return []byte{}, nil
}

func TestGetOHLC(t *testing.T) {
	client := &OsmosisClient{
		httpRequester: &MockHttpClient{
			t: t,
		},
	}

	currency := "NLS"
	step := 3600

	client.GetOHLC(currency, step)
}
