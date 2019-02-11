// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/dtamura/opentracing-tutorial-go/lib/log"
	"github.com/dtamura/opentracing-tutorial-go/lib/tracing"
	"github.com/dtamura/opentracing-tutorial-go/services/formatter"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var formatterOptions formatter.ConfigOptions

// formatterCmd represents the formatter command
var formatterCmd = &cobra.Command{
	Use:   "formatter",
	Short: "Start formatter Service",
	Long:  "Start formatter service",
	Run: func(cmd *cobra.Command, args []string) {
		zapLogger := logger.With(zap.String("service", "formatter"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("formatter", logger) // lesson03というサービス名のtracerを生成
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer) // to start the new spans, so we need to initialize that global variable to our instance of Jaeger tracer

		formatter := formatter.NewServer(
			formatterOptions,
			tracer,
			logger,
		)
		formatter.RunE()
	},
}

func init() {
	rootCmd.AddCommand(formatterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	formatterOptions.Port = 8081
}
