// ECB - european central bank
// This is static fetcher generating historical
// price from the fixed bgneur rate
// https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/eurofxref-graph-bgn.en.html

package ecb

import (
	"fmt"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

const (
	BGNEUR_PAIR  string = "bgneur"
	EURBGN_PAIR  string = "eurbgn"
	BGNEUR_PRICE        = 0.5113 // BGN 1 = EUR 0.5113
	EURBGN_PRICE        = 1.9558 // EUR 1 = BGN 1.9558
)

type ECBFetcher struct {
	pairs []string
	limit int
}

func NewECBFetcher(pairs []string, limit int) *ECBFetcher {
	return &ECBFetcher{
		pairs: pairs,
		limit: limit,
	}
}

func (f *ECBFetcher) GetMarketInfo(pair string) ([]types.MarketInfo, error) {
	if pair != BGNEUR_PAIR && pair != EURBGN_PAIR {
		return nil, fmt.Errorf("ecb fetcher unsupported pair: %v", pair)
	}

	var price float64
	if pair == BGNEUR_PAIR {
		price = BGNEUR_PRICE
	} else {
		price = EURBGN_PRICE
	}

	res := make([]types.MarketInfo, f.limit)
	for i := 0; i < f.limit; i++ {
		res[i].Price = price
	}

	return res, nil
}

func (f *ECBFetcher) GetAllMarketsInfo() (map[string][]types.MarketInfo, error) {
	res := make(map[string][]types.MarketInfo, len(f.pairs))
	for _, s := range f.pairs {
		marketInfo, err := f.GetMarketInfo(s)
		if err != nil {
			return nil, err
		}

		res[s] = marketInfo
	}

	return res, nil
}
