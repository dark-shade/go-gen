package main

import (
	"flag"
	"os"

	"github.com/dark-shade/go-gen/pkg/controller"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	debugMode := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	var core zapcore.Core

	if *debugMode {
		core = ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	} else {
		core = ecszap.NewCore(encoderConfig, os.Stdout, zap.ErrorLevel)
	}

	logger := zap.New(core, zap.AddCaller())
	logger = logger.Named("go-gen")

	controller.CreateStructure(logger)
}
