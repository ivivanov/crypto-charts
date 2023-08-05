/*
Copyright © 2023 Ivan ivanivanov12@gmail.com
*/
package cmd

import (
	"log"
	"time"

	"github.com/briandowns/spinner"

	"github.com/ivivanov/crypto-charts/pkg/fetchers"
	"github.com/ivivanov/crypto-charts/pkg/generators"
	"github.com/ivivanov/crypto-charts/pkg/job"
	"github.com/ivivanov/crypto-charts/pkg/types"
	"github.com/ivivanov/crypto-charts/pkg/uploader"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// fetchers
	LIMIT_DEFAULT = 168  // 1 week of data
	STEP_DEFAULT  = 3600 // 1 hour candles

	// uploader
	BUCKET_DEFAULT      = "crypto_charts"
	OBJECT_PATH_DEFAULT = "charts"

	// simple generator
	WIDTH_DEFAULT      = 800
	HEIGHT_DEFAULT     = 400
	LINE_COLOR_DEFAULT = 47 // yellow
	LINE_WIDTH_DEFAULT = 2.0
	MARGIN_DEFAULT     = 0.0
	BGR_COLOR_DEFAULT  = "#00000000" // transparent
)

var (
	// Config
	config          Config
	cfgFilePathFlag string
	printCfg        bool

	// uploader
	bucketFlag string
	pathFlag   string

	// simple generator
	widthFlag     int
	heightFlag    int
	lineColorFlag int
	lineWidthFlag float32
	marginFlag    float64
	bgrColorFlag  string
	advancedFlag  bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "crypto-charts",
		Short: "Creates and uploads SVG line charts",
		Long: `Periodically pulls OHLC data from a set of data sources. 
Generates simple line chart in SVG format and uploads 
it to a place of your choice.`,
		PreRun: preRun,
		Run: func(cmd *cobra.Command, args []string) {
			run := func() {
				bitstamp := config.Fetchers["bitstamp"]
				job := job.NewJob(
					[]types.Fetcher{
						fetchers.NewBitstampFetcher(bitstamp.Pairs, bitstamp.Step, bitstamp.Limit),
					},
					getGenerator(config.Generator.IsAdvanced),
					uploader.NewGoogleBucketUploader(config.Uploader.Bucket, config.Uploader.Path),
				)

				err := job.Run()
				if err != nil {
					log.Fatal(err)
				}
			}

			execWithLoading(run)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFilePathFlag, "config", "", "config file (default is $HOME/.crypto-charts.yaml or current dir)")
	rootCmd.PersistentFlags().BoolVar(&printCfg, "print", false, "print config file")

	// uploader
	rootCmd.PersistentFlags().StringVar(&bucketFlag, "bucket", BUCKET_DEFAULT, "Uploader: GCP cloud storage bucket name")
	viper.BindPFlag("uploader.bucket", rootCmd.PersistentFlags().Lookup("bucket"))

	rootCmd.PersistentFlags().StringVar(&pathFlag, "path", OBJECT_PATH_DEFAULT, "Uploader: object path")
	viper.BindPFlag("uploader.path", rootCmd.PersistentFlags().Lookup("path"))

	// generator
	rootCmd.PersistentFlags().BoolVar(&advancedFlag, "advanced", false, "Generator: SVG chart with time and price visualization")
	viper.BindPFlag("generator.is-advanced", rootCmd.PersistentFlags().Lookup("advanced"))

	// generator: simple
	rootCmd.PersistentFlags().IntVar(&widthFlag, "width", WIDTH_DEFAULT, "Generator: SVG chart width")
	viper.BindPFlag("generator.simple.width", rootCmd.PersistentFlags().Lookup("width"))

	rootCmd.PersistentFlags().IntVar(&heightFlag, "height", HEIGHT_DEFAULT, "Generator: SVG chart height")
	viper.BindPFlag("generator.simple.height", rootCmd.PersistentFlags().Lookup("height"))

	rootCmd.PersistentFlags().IntVar(&lineColorFlag, "line-color", LINE_COLOR_DEFAULT, "Generator: SVG chart line color HUE (default is yellow)")
	viper.BindPFlag("generator.simple.line-color", rootCmd.PersistentFlags().Lookup("line-color"))

	rootCmd.PersistentFlags().Float32Var(&lineWidthFlag, "line-width", LINE_WIDTH_DEFAULT, "Generator: SVG chart line thickness")
	viper.BindPFlag("generator.simple.line-width", rootCmd.PersistentFlags().Lookup("line-width"))

	rootCmd.PersistentFlags().Float64Var(&marginFlag, "margin", MARGIN_DEFAULT, "Generator: SVG chart margin")
	viper.BindPFlag("generator.simple.margin", rootCmd.PersistentFlags().Lookup("margin"))

	rootCmd.PersistentFlags().StringVar(&bgrColorFlag, "bgr-color", BGR_COLOR_DEFAULT, "Generator: SVG chart background color HEX (default is transparent)")
	viper.BindPFlag("generator.simple.bgr-color", rootCmd.PersistentFlags().Lookup("bgr-color"))

	// generator: advanced
	// todo
}

func execWithLoading(f func()) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Loading ... "
	start := time.Now()
	s.Start()
	f()
	s.Stop()
	log.Printf("Execution time %v\n", time.Since(start))
}

func getGenerator(isAdvanced bool) types.Generator {
	if isAdvanced {
		return generators.NewAdvancedLineChartGenerator()
	}

	return generators.NewSimpleLineChartGenerator(
		config.Generator.Simple.Width,
		config.Generator.Simple.Height,
		config.Generator.Simple.LineColor,
		config.Generator.Simple.LineWidth,
		config.Generator.Simple.BgrColor,
		config.Generator.Simple.Margin,
	)
}
