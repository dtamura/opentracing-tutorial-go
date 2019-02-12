package main

import (
	"os"

	lesson01 "github.com/dtamura/opentracing-tutorial-go/lesson01/services"
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger  *zap.Logger
	options lesson01.ConfigOptions
)

func main() {
	// 引数チェック
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	options.Message = os.Args[1]

	// loggerの初期化
	logger, _ = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := logger.With(zap.String("service", "lesson01"))
	logger := log.NewFactory(zapLogger)

	// OpenTracingの初期化
	tracer, closer := tracing.Init("lesson01", logger) // lesson01というService名のtracerを生成
	defer closer.Close()

	// lesson01 を実行するクライアントを初期化
	client := lesson01.NewClient(
		options,
		tracer,
		logger,
	)
	// 実行
	client.Run()
}
