package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	IsAdvanced bool              `mapstructure:"is_advanced"`
	Simple     SimpleGenerator   `mapstructure:"simple"`
	Advanced   AdvancedGenerator `mapstructure:"advanced"`
}

type Config struct {
	Fetchers  map[string]Fetcher `mapstructure:"fetchers"`
	Uploader  Uploader           `mapstructure:"uploader"`
	Generator Generator          `mapstructure:"generator"`
}

func initConfig() {
	if cfgFilePathFlag != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFilePathFlag)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.AddConfigPath("/")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".crypto-charts")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func preRun(ccmd *cobra.Command, args []string) {
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to read Viper options into configuration: %v", err)
	}

	if printCfg {
		fmt.Printf("%+v\n", config)
	}
}
