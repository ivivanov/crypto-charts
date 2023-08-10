package osmosis

import (
	"github.com/ivivanov/crypto-charts/pkg/types"
)

type OsmosisFetcher struct {
	client    *OsmosisClient
	symbols   []string
	timeframe int
}

func NewOsmosisFetcher(symbols []string, timeframe int) *OsmosisFetcher {
	return &OsmosisFetcher{
		client:    NewOsmosisClient(),
		symbols:   symbols,
		timeframe: timeframe,
	}
}

func (f *OsmosisFetcher) GetMarketInfo(symbol string) ([]types.MarketInfo, error) {
	data, err := f.client.GetOHLC(symbol, f.timeframe)
	if err != nil {
		return nil, err
	}

	return MapOHLCtoMarketInfo(data), nil
}

// GetAllMarketsInfo return market info for each pair.
func (f *OsmosisFetcher) GetAllMarketsInfo() (map[string][]types.MarketInfo, error) {
	res := make(map[string][]types.MarketInfo, len(f.symbols))
	for _, s := range f.symbols {
		marketInfo, err := f.GetMarketInfo(s)
		if err != nil {
			return nil, err
		}

		res[s] = marketInfo
	}

	return res, nil
}

func MapOHLCtoMarketInfo(ohlc []OHLC) []types.MarketInfo {
	res := make([]types.MarketInfo, len(ohlc))
	for i := 0; i < len(ohlc); i++ {
		res[i].Price = ohlc[i].Close
		res[i].Timestamp = ohlc[i].Time
	}

	return res
}
