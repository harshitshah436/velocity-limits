package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MaxLoadLimitPerDay  float64 `mapstructure:"MAX_LOAD_LIMIT_PER_DAY"`
	MaxLoadLimitPerWeek float64 `mapstructure:"MAX_LOAD_LIMIT_PER_WEEK"`
	MaxLoadPerDay       int     `mapstructure:"MAX_LOAD_PER_DAY"`
	InputFile           string  `mapstructure:"INPUT_FILE"`
	OutputFile          string  `mapstructure:"OUTPUT_FILE"`
}

type Configuration struct {
	Config `mapstructure:"config"`
}

func LoadConfig(path string) (config Configuration) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error - Reading config file: ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Error - Unmarshal config file content: ", err)
	}

	return
}
