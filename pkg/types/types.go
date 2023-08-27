package types

type Fetcher interface {
	GetMarketInfo(currencyPair string) ([]MarketInfo, error)
	GetAllMarketsInfo() (map[string][]MarketInfo, error)
}

type SynthGen interface {
	GenerateSynth(pairA, pairB string, allMarketInfos map[string][]MarketInfo) ([]MarketInfo, error)
}

type Generator interface {
	NewLineChart(historicalData []MarketInfo) (string, error)
}

type Uploader interface {
	UploadSVG(pair, svg string) error
}

type MarketInfo struct {
	Price     float64
	Volume    float64
	Timestamp int64
}
