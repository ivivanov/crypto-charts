package job

import (
	"log"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

type Job struct {
	fetchers  []types.Fetcher
	generator types.Generator
	uploader  types.Uploader
}

func NewJob(fetchers []types.Fetcher, generator types.Generator, uploader types.Uploader) *Job {
	return &Job{
		fetchers:  fetchers,
		generator: generator,
		uploader:  uploader,
	}
}

func (j *Job) Run() error {
	// fetch HistoricalData
	for _, fetcher := range j.fetchers {
		marketInfos, err := fetcher.GetAllMarketsInfo()
		if err != nil {
			return err
		}

		for pair, mi := range marketInfos {
			if len(mi) == 0 {
				continue // skip pairs which has no updates
			}

			// generate line chart svg
			svg, err := j.generator.NewLineChart(mi)
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
