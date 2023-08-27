package cmd

type Fetcher struct {
	Limit int      `mapstructure:"limit"`
	Step  int      `mapstructure:"step"`
	Pairs []string `mapstructure:"pairs"`
}

type Uploader struct {
	Bucket string `mapstructure:"bucket"`
	Path   string `mapstructure:"path"`
}

type SimpleGenerator struct {
	Width     int     `mapstructure:"width"`
	Height    int     `mapstructure:"height"`
	LineColor int     `mapstructure:"line-color"`
	LineWidth float32 `mapstructure:"line-width"`
	Margin    float64 `mapstructure:"margin"`
	BgrColor  string  `mapstructure:"bgr-color"`
}

type AdvancedGenerator struct {
}

type Generator struct {
	IsAdvanced bool              `mapstructure:"is-advanced"`
	Simple     SimpleGenerator   `mapstructure:"simple"`
	Advanced   AdvancedGenerator `mapstructure:"advanced"`
}

type Config struct {
	Fetchers  map[string]Fetcher  `mapstructure:"fetchers"`
	Synths    map[string][]string `mapstructure:"synths"`
	Uploader  Uploader            `mapstructure:"uploader"`
	Generator Generator           `mapstructure:"generator"`
}
