package generators

import (
	"bytes"

	m "github.com/erkkah/margaid"

	"github.com/ivivanov/crypto-charts/pkg/types"
)

type SimpleLineChartGenerator struct {
	width     int
	height    int
	lineColor int
	lineWidth float32
	bgrColor  string
	margin    float64
}

func NewSimpleLineChartGenerator(width, height, lineColor int, lineWidth float32, bgrColor string, margin float64) *SimpleLineChartGenerator {
	return &SimpleLineChartGenerator{
		width:     width,
		height:    height,
		lineWidth: lineWidth,
		lineColor: lineColor,
		bgrColor:  bgrColor,
		margin:    margin,
	}
}

// NewLineChart returns svg representing the data in a simple line chart.
func (l *SimpleLineChartGenerator) NewLineChart(historicalData []types.MarketInfo) (string, error) {
	priceSeries := m.NewSeries()
	for i := 0; i < len(historicalData); i++ {
		priceSeries.Add(m.MakeValue(float64(i), historicalData[i].Price))
	}

	chart := m.New(l.width, l.height,
		m.WithAutorange(m.XAxis, priceSeries),
		m.WithAutorange(m.YAxis, priceSeries),
		m.WithAutorange(m.Y2Axis, priceSeries),
		// m.WithProjection(m.YAxis, m.Log),
		m.WithColorScheme(l.lineColor),
		m.WithBackgroundColor(l.bgrColor),
		m.WithInset(l.margin),
	)

	chart.Line(priceSeries,
		m.UsingAxes(m.XAxis, m.YAxis),
		m.UsingStrokeWidth(l.lineWidth),
	)

	buf := bytes.NewBuffer(nil)
	chart.Render(buf)

	return buf.String(), nil
}
