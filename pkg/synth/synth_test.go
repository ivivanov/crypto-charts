package synth

import (
	"testing"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

// 1) bgn quote currency
// btceur:	1 btc = 26822 eur
// bgneur:	1 bgn = 0.5113 eur
// btcbgn:	1 btc = 26822/0.5113 = 52458.439272443 bgn
//
// 2) eur quote currency
// eurusd:	1 eur = 1.08704 usd
// eurbgn:	1 eur = 1.95583 bgn
// bgnusd:	1 bgn =  1.08704/1.95583 = 0.555794727 usd
//
// 3) bgn quote currency
// eurusd:	1 eur = 1.08704 usd
// eurbgn:	1 eur = 1.95583 bgn
// usdbgn:	1 bgn =  1.95583/1.08704 = 1.799225419 bgn
func TestGenerateSynth(t *testing.T) {
	allMarketInfos := map[string][]types.MarketInfo{
		"bgneur": {
			{Price: 0.5113},
			{Price: 0.5113},
		},
		"eurbgn": {
			{Price: 1.95583},
			{Price: 1.95583},
		},
		"btceur": {{Price: 26822}},
		"eurusd": {{Price: 1.08704}},
	}

	s := NewSynth()

	for _, tc := range []struct {
		name   string
		pairA  string
		pairB  string
		exp    []types.MarketInfo
		expErr bool
	}{
		{
			name:   "Successful generate of matching quotes currencies",
			pairA:  "btceur",
			pairB:  "bgneur",
			exp:    []types.MarketInfo{{Price: 52458.439272442796}},
			expErr: false,
		},
		{
			name:   "Successful generate of matching base currencies",
			pairA:  "eurusd",
			pairB:  "eurbgn",
			exp:    []types.MarketInfo{{Price: 0.5557947265355374}},
			expErr: false,
		},
		{
			name:   "Not matching base or quote currencies should return error",
			pairA:  "foobar",
			pairB:  "bgneur",
			exp:    nil,
			expErr: true,
		},
		{
			name:   "Missing pair data should return error",
			pairA:  "eurusd",
			pairB:  "usdusd",
			exp:    nil,
			expErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act, err := s.GenerateSynth(tc.pairA, tc.pairB, allMarketInfos)
			if tc.expErr {
				if err == nil {
					t.Fatal("error is expected")
				}

				return
			}

			expLen := len(tc.exp)
			actLen := len(act)
			if expLen != actLen {
				t.Fatalf("unexpected result length, exp: %v, act: %v", expLen, actLen)
			}

			for i := 0; i < expLen; i++ {
				if tc.exp[i].Price != act[i].Price {
					t.Fatalf("unexpected price, exp: %v, act: %v", tc.exp[i].Price, act[i].Price)
				}
			}
		})
	}
}
