package job

import (
	"fmt"
	"log"
	"os"

	"github.com/ivivanov/crypto-charts/pkg/fetchers"
	"github.com/ivivanov/crypto-charts/pkg/generator"
	"github.com/ivivanov/crypto-charts/pkg/types"
	"github.com/ivivanov/crypto-charts/pkg/uploader"
)

const (
	LIMIT = 168  // 1 week of data
	STEP  = 3600 // 1 hour candles
)

type Fetcher interface {
	GetMarketInfo(currencyPair string, step, limit int) ([]types.MarketInfo, error)
	GetAllMarketInfo(step, limit int) (map[string][]types.MarketInfo, error)
}

type Generator interface {
	NewLineChartSVG(historicalData []types.MarketInfo) (string, error)
}

type Uploader interface {
	UploadSVG(pair, svg string) error
}

type Job struct {
	fetchers  []Fetcher
	generator Generator
	uploader  Uploader
}

func NewJob(envPath string) (*Job, error) {
	err := InitDotEnv(envPath)
	if err != nil {
		return nil, err
	}

	bsPairs := os.Getenv("BITSTAMP_PAIRS_LIST")
	if bsPairs == "" {
		return nil, fmt.Errorf("BITSTAMP_PAIRS_LIST is empty")
	}

	bsPairsArr := GetArrayFrom(bsPairs)

	return &Job{
		fetchers:  append([]Fetcher{}, fetchers.NewBitstampFetcher(bsPairsArr)),
		generator: &generator.LineChartGenerator{},
		uploader:  &uploader.GoogleBucketUploader{},
	}, nil
}

func (j *Job) Run() error {
	// fetch HistoricalData
	for _, fetcher := range j.fetchers {
		marketInfos, err := fetcher.GetAllMarketInfo(STEP, LIMIT)
		if err != nil {
			return err
		}

		for pair, mi := range marketInfos {
			if len(mi) == 0 {
				continue // skip pairs which has no updates
			}

			// generate line chart svg
			svg, err := j.generator.NewLineChartSVG(mi)
			if err != nil {
				return err
			}

			// store svg into bucket
			j.uploader.UploadSVG(pair, svg)
		}
	}

	log.Println("Successful update")

	return nil
}
