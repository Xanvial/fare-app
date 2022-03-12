package model

import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	Fare      FareConfig `json:"fare"`
	LogConfig Log        `json:"log"`
}

type Log struct {
	LogLevel    int      `json:"log_level"`    // default InfoLevel
	OutputPaths []string `json:"output_paths"` // default "./fare-app.log"
}

type FareConfig struct {
	MinData        int `json:"min_data"`         // default 2
	MaxIntervalSec int `json:"max_interval_sec"` // default 600 seconds (5 minutes)

	BaseFare         int     `json:"base_fare"`      // default 400
	BaseFareDistance float64 `json:"base_fare_dist"` // default 1000

	LimitDistance          float64 `json:"limit_dist"`            // default 10000
	UnderLimitFare         int     `json:"under_limit_fare"`      // default 40
	UnderLimitFareDistance float64 `json:"under_limit_fare_dist"` // default 400
	OverLimitFare          int     `json:"over_limit_fare"`       // default 40
	OverLimitFareDistance  float64 `json:"over_limit_fare_dist"`  // default 350
}

func ReadConfig(file string) Config {
	// Base Config
	config := Config{
		LogConfig: Log{
			LogLevel: int(zapcore.InfoLevel),
			OutputPaths: []string{
				"./fare-app.log",
			},
		},
		Fare: FareConfig{
			MinData:        2,
			MaxIntervalSec: 600,

			BaseFare:         400,
			BaseFareDistance: 1000,

			LimitDistance:          10000,
			UnderLimitFare:         40,
			UnderLimitFareDistance: 400,
			OverLimitFare:          40,
			OverLimitFareDistance:  350,
		},
	}

	// if no path provided, use base config
	if file == "" {
		return config
	}

	configFile, err := os.Open(file)
	if err != nil {
		// log the error using default logger because Zap is not initialized yet
		// and return default config
		log.Println("error opening config file, path:", file, " err:", err)
		return config
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		// log the error using default logger because Zap is not initialized yet
		// and return default config
		log.Println("error decoding config file, path:", file, " err:", err)
		return config
	}

	return config
}
