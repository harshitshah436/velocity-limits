// Package config reads configurations from config.toml file and create the struct.
package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct contains all velocity limits and input, output file names.
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

// LoadConfig loads velocity configs and filenames from config.toml file.
func LoadConfig(path string) (config Configuration) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")

	// Try reading a config.toml file.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error - Reading config file: ", err)
	}

	// Unmarshal configs into a configuration struct.
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Error - Unmarshal config file content: ", err)
	}

	return
}
