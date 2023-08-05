package fetchers

import (
	"errors"

	"github.com/ivivanov/crypto-charts/pkg/data/bitstamp"
	"github.com/ivivanov/crypto-charts/pkg/types"
)

type BitstampFetcher struct {
	client *bitstamp.BitstampClient
	pairs  []string
	step   int
	limit  int
}

func NewBitstampFetcher(pairs []string, step, limit int) *BitstampFetcher {
	return &BitstampFetcher{
		client: bitstamp.NewBitstampClient(),
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

	return types.MapOHLCtoHistoricalData(data), nil
}

// GetAllMarketsInfo return market info for each pair.
func (f *BitstampFetcher) GetAllMarketsInfo() (map[string][]types.MarketInfo, error) {
	res := make(map[string][]types.MarketInfo, len(f.pairs))
	for _, pair := range f.pairs {
		marketInfo, err := f.GetMarketInfo(pair)
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
