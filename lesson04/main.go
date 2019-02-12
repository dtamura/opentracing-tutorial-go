package main

import (
	"os"

	lesson04 "github.com/dtamura/opentracing-tutorial-go/lesson04/services"
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger  *zap.Logger
	options lesson04.ConfigOptions
)

func main() {

	if len(os.Args) != 3 {
		panic("ERROR: Expecting two argument")
	}
	options.Message = os.Args[1]
	options.Greeting = os.Args[2]

	logger, _ = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := logger.With(zap.String("service", "lesson04"))
	logger := log.NewFactory(zapLogger)
	tracer, closer := tracing.Init("lesson04", logger)
	client := lesson04.NewClient(
		options,
		tracer,
		logger,
	)
	defer closer.Close()
	client.Run()
}
