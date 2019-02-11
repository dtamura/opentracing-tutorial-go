package cmd

import (
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"github.com/dtamura/opentracing-tutorial-go/services/lesson01"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var lesson01Options = &lesson01.ConfigOptions{}

// lesson01Cmd represents the lesson01 command
var lesson01Cmd = &cobra.Command{
	Use:   "lesson01",
	Short: "start lesson01 program",
	Long:  "Start Lesson01 Program",
	Run: func(cmd *cobra.Command, args []string) {

		zapLogger := logger.With(zap.String("service", "lesson01"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("lesson01", logger) // lesson01というサービス名のtracerを生成
		client := lesson01.NewClient(
			lesson01Options,
			tracer,
			logger,
		)
		defer closer.Close()
		client.RunE()
	},
}

func init() {
	rootCmd.AddCommand(lesson01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lesson01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lesson01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lesson01Cmd.PersistentFlags().StringVarP(&lesson01Options.Message, "message", "m", "tamura", "Message") // -mオプションで文字列を取得する

}
