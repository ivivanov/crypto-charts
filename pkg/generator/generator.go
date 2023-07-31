package generator

import (
	"bytes"

	"github.com/erkkah/margaid"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

const (
	WIDTH       = 800
	HEIGHT      = 400
	LINE_COLOR  = 47          // yellow
	BGR_COLOR   = "#00000000" // transparent
	LINE_WIDTH  = 2
	MARGIN_ZERO = 0
)

type LineChartGenerator struct{}

// NewLineChartSVG returns svg representing the data in a simple line chart.
func (l *LineChartGenerator) NewLineChartSVG(historicalData []types.MarketInfo) (string, error) {
	priceSeries := margaid.NewSeries()
	for i := 0; i < len(historicalData); i++ {
		priceSeries.Add(margaid.MakeValue(float64(i), historicalData[i].Price))
	}

	chart := margaid.New(WIDTH, HEIGHT,
		margaid.WithAutorange(margaid.XAxis, priceSeries),
		margaid.WithAutorange(margaid.YAxis, priceSeries),
		margaid.WithAutorange(margaid.Y2Axis, priceSeries),
		margaid.WithProjection(margaid.YAxis, margaid.Log),
		margaid.WithColorScheme(LINE_COLOR),
		margaid.WithBackgroundColor(BGR_COLOR),
		margaid.WithInset(MARGIN_ZERO),
	)

	chart.Line(priceSeries,
		margaid.UsingAxes(margaid.XAxis, margaid.YAxis),
		margaid.UsingStrokeWidth(LINE_WIDTH),
	)

	buf := bytes.NewBuffer(nil)
	chart.Render(buf)

	return buf.String(), nil
}
