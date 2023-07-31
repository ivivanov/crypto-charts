package types

import "github.com/ivivanov/crypto-charts/pkg/data/bitstamp"

type MarketInfo struct {
	Price     float64
	Volume    float64
	Timestamp string
}

func MapOHLCtoHistoricalData(ohlc []bitstamp.OHLC) []MarketInfo {
	res := make([]MarketInfo, len(ohlc))
	for i := 0; i < len(ohlc); i++ {
		res[i].Price = ohlc[i].Close
		res[i].Volume = ohlc[i].Volume
		res[i].Timestamp = ohlc[i].Timestamp
	}

	return res
}
