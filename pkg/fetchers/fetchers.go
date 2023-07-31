package fetchers

import (
	"errors"

	"github.com/ivivanov/crypto-charts/pkg/data/bitstamp"
	"github.com/ivivanov/crypto-charts/pkg/types"
)

type BitstampFetcher struct {
	client *bitstamp.BitstampClient
	pairs  []string
}

func NewBitstampFetcher(pairs []string) *BitstampFetcher {
	return &BitstampFetcher{
		client: bitstamp.NewBitstampClient(),
		pairs:  pairs,
	}
}

func (f *BitstampFetcher) GetMarketInfo(currencyPair string, step, limit int) ([]types.MarketInfo, error) {
	data, err := f.client.GetOHLC(currencyPair, step, limit)
	if err != nil {
		return nil, err
	}

	return types.MapOHLCtoHistoricalData(data), nil
}

// GetAllMarketInfo return market info for each pair.
func (f *BitstampFetcher) GetAllMarketInfo(step, limit int) (map[string][]types.MarketInfo, error) {
	res := make(map[string][]types.MarketInfo, len(f.pairs))
	for _, pair := range f.pairs {
		marketInfo, err := f.GetMarketInfo(pair, step, limit)
		if err != nil {
			if errors.Is(err, bitstamp.ErrUnmarshalErrorResponse) {
				continue // do not break the entire fetch because such error might be just a single failed request
			}

			return nil, err
		}

		res[pair] = marketInfo
	}

	return res, nil
}
