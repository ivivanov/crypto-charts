package generators

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ivivanov/crypto-charts/pkg/types"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

var (
	VeryLightGrey = drawing.Color{R: 187, G: 190, B: 191, A: 100}
	LightGrey     = drawing.Color{R: 187, G: 190, B: 191, A: 250}
)

type AdvancedLineChartGenerator struct{}

func NewAdvancedLineChartGenerator() *AdvancedLineChartGenerator {
	return &AdvancedLineChartGenerator{}
}

func (l *AdvancedLineChartGenerator) NewLineChart(historicalData []types.MarketInfo) (string, error) {
	xv := xvalues(historicalData)
	yv := yvalues(historicalData)

	priceSeries := chart.ContinuousSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xv,
		YValues: yv,
	}

	// smaSeries := chart.SMASeries{
	// 	Style: chart.Style{
	// 		StrokeColor:     drawing.ColorRed,
	// 		StrokeDashArray: []float64{5.0, 5.0},
	// 	},
	// 	InnerSeries: priceSeries,
	// }

	// bbSeries := &chart.BollingerBandsSeries{
	// 	Style: chart.Style{
	// 		StrokeColor: drawing.ColorFromHex("efefef"),
	// 		FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
	// 	},
	// 	InnerSeries: priceSeries,
	// }

	graph := chart.Chart{
		// maintain ratio of 16:9
		Width:  960,
		Height: 540,
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionBetweenTicks,
			TickStyle: chart.Style{
				FontSize:  8,
				FontColor: LightGrey,
			},
			Style: chart.Style{
				StrokeColor: VeryLightGrey,
			},
			GridMajorStyle: chart.Style{
				StrokeColor: VeryLightGrey,
				StrokeWidth: 1.0,
			},
			GridMinorStyle: chart.Style{
				StrokeColor: VeryLightGrey,
				StrokeWidth: 1.0,
			},
			ValueFormatter: func(v interface{}) string {
				vf, isFloat := v.(float64)
				if !isFloat {
					return ""
				}
				return time.Unix(int64(vf), 0).Format("01/06")
			},
		},
		YAxis: chart.YAxis{
			TickStyle: chart.Style{
				FontSize:  8,
				FontColor: LightGrey,
			},
			Style: chart.Style{
				StrokeColor: VeryLightGrey,
			},
			GridMajorStyle: chart.Style{
				StrokeColor: VeryLightGrey,
				StrokeWidth: 1.0,
			},
			GridMinorStyle: chart.Style{
				StrokeColor: VeryLightGrey,
				StrokeWidth: 1.0,
			},
			ValueFormatter: func(v interface{}) string {
				vf, isFloat := v.(float64)
				if !isFloat {
					return ""
				}

				return formatPrice(vf)
			},
		},
		YAxisSecondary: chart.YAxis{
			Style: chart.Style{
				StrokeColor: chart.ColorTransparent,
			},
		},
		Series: []chart.Series{
			priceSeries,
			// bbSeries,
			// smaSeries,
			// fancy label of the current price
			// chart.AnnotationSeries{
			// 	Style: chart.Style{
			// 		Padding:     chart.Box{Top: 2, Left: -1, Right: 8, Bottom: 4},
			// 		FillColor:   chart.ColorAlternateBlue,
			// 		StrokeColor: drawing.ColorWhite,
			// 	},

			// 	Annotations: []chart.Value2{
			// 		{XValue: xv[len(xv)-1], YValue: yv[len(yv)-1], Label: formatPrice(yv[len(yv)-1])},
			// 	},
			// },
		},
		Canvas: chart.Style{
			FillColor: chart.ColorTransparent,
		},
	}

	buf := bytes.NewBuffer(nil)
	graph.Render(chart.SVG, buf)

	return buf.String(), nil
}

func xvalues(historicalData []types.MarketInfo) []float64 {
	epochs := make([]float64, len(historicalData))

	for i := 0; i < len(historicalData); i++ {
		epochs[i] = float64(historicalData[i].Timestamp)
	}

	return epochs
}

func yvalues(historicalData []types.MarketInfo) []float64 {
	prices := make([]float64, len(historicalData))

	for i := 0; i < len(historicalData); i++ {
		prices[i] = historicalData[i].Price
	}

	return prices
}

func formatPrice(value float64) string {
	switch {
	case value < 0.0001: //6
		return fmt.Sprintf("%0.6f", value)
	case value < 0.001: //4
		return fmt.Sprintf("%0.4f", value)
	case value < 1000: //2
		return fmt.Sprintf("%0.2f", value)
	}

	return fmt.Sprintf("%0.0f", value)
}
