package bitstamp

import (
	"errors"
	"strconv"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

type BitstampFetcher struct {
	client *BitstampClient
	pairs  []string
	step   int
	limit  int
}

func NewBitstampFetcher(pairs []string, step, limit int) *BitstampFetcher {
	return &BitstampFetcher{
		client: NewBitstampClient(),
		pairs:  pairs,
		step:   step,
		limit:  limit,
	}
}

func (f *BitstampFetcher) GetMarketInfo(currencyPair string) ([]types.MarketInfo, error) {
	data, err := f.client.GetOHLC(currencyPair, f.step, f.limit)
	if err != nil {
		return nil, err
	}

	return MapOHLCtoMarketInfo(data), nil
}

// GetAllMarketsInfo return market info for each pair.
func (f *BitstampFetcher) GetAllMarketsInfo() (map[string][]types.MarketInfo, error) {
	res := make(map[string][]types.MarketInfo, len(f.pairs))
	for _, pair := range f.pairs {
		marketInfo, err := f.GetMarketInfo(pair)
		if err != nil {
			if errors.Is(err, ErrUnmarshalErrorResponse) {
				continue // do not break the entire fetch because such error might be just a single failed request
			}

			return nil, err
		}

		res[pair] = marketInfo
	}

	return res, nil
}

func MapOHLCtoMarketInfo(ohlc []OHLC) []types.MarketInfo {
	res := make([]types.MarketInfo, len(ohlc))
	for i := 0; i < len(ohlc); i++ {
		res[i].Price = ohlc[i].Close
		res[i].Volume = ohlc[i].Volume
		time, _ := strconv.ParseInt(ohlc[i].Timestamp, 10, 64)
		res[i].Timestamp = time
	}

	return res
}
