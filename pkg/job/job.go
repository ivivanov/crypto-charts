package job

import (
	"fmt"
	"log"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

type Job struct {
	fetchers  []types.Fetcher
	synthGen  types.SynthGen
	generator types.Generator
	uploader  types.Uploader
	synths    map[string][]string
}

func NewJob(
	fetchers []types.Fetcher,
	synthGen types.SynthGen,
	generator types.Generator,
	uploader types.Uploader,
	synths map[string][]string) *Job {
	return &Job{
		fetchers:  fetchers,
		synthGen:  synthGen,
		generator: generator,
		uploader:  uploader,
		synths:    synths,
	}
}

func (j *Job) Run() error {
	allMarketInfos := make(map[string][]types.MarketInfo)

	// fetch HistoricalData
	for _, fetcher := range j.fetchers {
		marketInfos, err := fetcher.GetAllMarketsInfo()
		if err != nil {
			return err
		}

		for k, v := range marketInfos {
			if _, ok := allMarketInfos[k]; !ok {
				allMarketInfos[k] = v
			}
		}
	}

	// generate synthetic pairs
	for k, v := range j.synths {
		synthMarketInfo, err := j.synthGen.GenerateSynth(v[0], v[1], allMarketInfos)
		if err != nil {
			return fmt.Errorf("failed to generate %v: %w", k, err)
		}

		if _, ok := allMarketInfos[k]; !ok {
			allMarketInfos[k] = synthMarketInfo
		}
	}

	// Generate charts and upload
	for pair, mi := range allMarketInfos {
		if len(mi) == 0 {
			continue // skip pairs with no updates
		}

		// generate line chart svg
		svg, err := j.generator.NewLineChart(mi)
		if err != nil {
			return err
		}

		// store svg into bucket
		j.uploader.UploadSVG(pair, svg)
	}

	log.Println("Successful update")

	return nil
}
