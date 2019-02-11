package cmd

import (
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"github.com/dtamura/opentracing-tutorial-go/services/lesson02"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var lesson02Options = &lesson02.ConfigOptions{}

// lesson02Cmd represents the lesson02 command
var lesson02Cmd = &cobra.Command{
	Use:   "lesson02",
	Short: "start lesson02 program",
	Long:  "Start lesson02 Program",
	Run: func(cmd *cobra.Command, args []string) {

		zapLogger := logger.With(zap.String("service", "lesson02"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("lesson02", logger) // lesson02というサービス名のtracerを生成
		opentracing.SetGlobalTracer(tracer)                // to start the new spans, so we need to initialize that global variable to our instance of Jaeger tracer
		client := lesson02.NewClient(
			lesson02Options,
			tracer,
			logger,
		)
		defer closer.Close()
		client.RunE()
	},
}

func init() {
	rootCmd.AddCommand(lesson02Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lesson02Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lesson02Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lesson02Cmd.PersistentFlags().StringVarP(&lesson02Options.Message, "message", "m", "tamura", "Message") // -mオプションで文字列を取得する

}
