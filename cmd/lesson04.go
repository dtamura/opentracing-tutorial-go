package cmd

import (
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"github.com/dtamura/opentracing-tutorial-go/services/lesson04"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var lesson04Options = &lesson04.ConfigOptions{}

// lesson04Cmd represents the lesson04 command
var lesson04Cmd = &cobra.Command{
	Use:   "lesson04",
	Short: "start lesson04 program",
	Long:  "Start lesson04 Program",
	Run: func(cmd *cobra.Command, args []string) {

		zapLogger := logger.With(zap.String("service", "lesson04"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("lesson04", logger) // lesson04というサービス名のtracerを生成
		opentracing.SetGlobalTracer(tracer)                // to start the new spans, so we need to initialize that global variable to our instance of Jaeger tracer
		// Client
		client := lesson04.NewClient(
			lesson04Options,
			tracer,
			logger,
			closer,
		)
		defer closer.Close()
		client.RunE()
	},
}

func init() {
	rootCmd.AddCommand(lesson04Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lesson04Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lesson04Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lesson04Cmd.PersistentFlags().StringVarP(&lesson04Options.Message, "message", "m", "tamura", "Message")    // -mオプションで文字列を取得する
	lesson04Cmd.PersistentFlags().StringVarP(&lesson04Options.Greeting, "greetng", "g", "Bonjour", "Greeting") // -gオプションで文字列を取得する

}
