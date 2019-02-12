package main

import (
	"os"

	lesson03 "github.com/dtamura/opentracing-tutorial-go/lesson03/services"
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger  *zap.Logger
	options lesson03.ConfigOptions
)

func main() {
	// 引数チェック
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	options.Message = os.Args[1]

	// loggerの初期化
	logger, _ = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapLogger := logger.With(zap.String("service", "lesson03"))
	logger := log.NewFactory(zapLogger)

	// OpenTracingの初期化
	tracer, closer := tracing.Init("lesson03", logger) // lesson03というサービス名のtracerを生成
	defer closer.Close()

	// lesson03 を実行するクライアントの初期化
	client := lesson03.NewClient(
		options,
		tracer,
		logger,
	)
	// 実行
	client.Run()
}
