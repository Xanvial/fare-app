package main

import (
	"flag"
	"os"

	"github.com/Xanvial/fare-app/internal/model"
	"github.com/Xanvial/fare-app/internal/presenter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Read input flag, currently only to use specific config file
	inputFile := flag.String("input", "", "data input that will be processed by the application")
	cfgPath := flag.String("config", "", "input config path, use absolute path or path relative to binary")
	flag.Parse()

	// read config file if exist, if not this will return base config specified inside model
	cfg := model.ReadConfig(*cfgPath)

	// init zap logger using production config
	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	// possible to override log level if needed, default to info on higher
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(cfg.LogConfig.LogLevel))
	// possible to override log path if needed, default to "./fare-app.log"
	config.OutputPaths = cfg.LogConfig.OutputPaths

	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	mainApp := presenter.New(cfg)

	if inputFile == nil || len(*inputFile) == 0 {
		// if there's no input file, read input from standard IO
		mainApp.Run(os.Stdin, true)
	} else {
		// if input file exist, try to open it and process the data
		file, err := os.Open(*inputFile)
		if err != nil {
			logger.Error("unable to open inputfile",
				zap.String("input_file", *inputFile))
			return
		}
		defer file.Close()
		mainApp.Run(file, false)
	}
}
