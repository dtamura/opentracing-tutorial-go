package main

import (
	"os"

	lesson02 "github.com/dtamura/opentracing-tutorial-go/lesson02/services"
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger  *zap.Logger
	options lesson02.ConfigOptions
)

func main() {
	// 引数チェック
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	options.Message = os.Args[1]

	// loggerの初期化
	logger, _ = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := logger.With(zap.String("service", "lesson02"))
	logger := log.NewFactory(zapLogger)

	// OpenTracingの初期化
	tracer, closer := tracing.Init("lesson02", logger)
	defer closer.Close()

	// lesson02 を実行するクライアントを初期化
	client := lesson02.NewClient(
		options,
		tracer,
		logger,
	)
	// 実行
	client.Run()
}
