package cmd

import (
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"github.com/dtamura/opentracing-tutorial-go/services/lesson03"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var lesson03Options = &lesson03.ConfigOptions{}

// lesson03Cmd represents the lesson03 command
var lesson03Cmd = &cobra.Command{
	Use:   "lesson03",
	Short: "start lesson03 program",
	Long:  "Start lesson03 Program",
	Run: func(cmd *cobra.Command, args []string) {

		zapLogger := logger.With(zap.String("service", "lesson03"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("lesson03", logger) // lesson03というサービス名のtracerを生成
		opentracing.SetGlobalTracer(tracer)                // to start the new spans, so we need to initialize that global variable to our instance of Jaeger tracer
		// Client
		client := lesson03.NewClient(
			lesson03Options,
			tracer,
			logger,
			closer,
		)
		defer closer.Close()
		client.RunE()
	},
}

func init() {
	rootCmd.AddCommand(lesson03Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lesson03Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lesson03Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lesson03Cmd.PersistentFlags().StringVarP(&lesson03Options.Message, "message", "m", "tamura", "Message") // -mオプションで文字列を取得する

}
