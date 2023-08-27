package synth

import (
	"fmt"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

type Synth struct {
}

func NewSynth() *Synth {
	return &Synth{}
}

// GenerateSynth calculates and creates series for new pair
func (s *Synth) GenerateSynth(pairA, pairB string, allMarketInfos map[string][]types.MarketInfo) ([]types.MarketInfo, error) {
	// Ensure that the quote OR base currency on both a & b is the same
	// Get base currency assuming its 3 chars
	prefix1 := pairA[:3]
	prefix2 := pairB[:3]

	// Get quote currency assuming its 3 chars
	suffix1 := pairA[len(pairA)-3:]
	suffix2 := pairB[len(pairB)-3:]

	baseAreEq := prefix1 == prefix2
	quoteAreEq := suffix1 == suffix2

	if !baseAreEq && !quoteAreEq {
		return nil, fmt.Errorf("neither quote or base currencies match in: %v, %v", pairA, pairB)
	}

	// Get the historical data form MarketInfos map
	infoA, ok := allMarketInfos[pairA]
	if !ok {
		return nil, fmt.Errorf("missing info for: %v", pairA)
	}

	infoB, ok := allMarketInfos[pairB]
	if !ok {
		return nil, fmt.Errorf("missing info for: %v", pairB)
	}

	length := min(len(infoA), len(infoB))
	result := make([]types.MarketInfo, length)
	for i := 0; i < length; i++ {
		result[i].Price = infoA[i].Price / infoB[i].Price
		// TODO: think of better way to provide timestamp
		// This is dummy fix for missing timestamp
		// ECB fetcher generates artificial price series without timestamp
		// We have to take the timestamp from the exchange fetcher
		timestamp := infoA[i].Timestamp
		if timestamp == 0 {
			timestamp = infoB[i].Timestamp
		}

		result[i].Timestamp = timestamp
	}

	return result, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
