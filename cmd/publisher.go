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
	"github.com/dtamura/opentracing-tutorial-go/services/publisher"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// publisherCmd represents the publisher command
var publisherCmd = &cobra.Command{
	Use:   "publisher",
	Short: "start publisher service",
	Long:  "start publisher service",
	Run: func(cmd *cobra.Command, args []string) {

		zapLogger := logger.With(zap.String("service", "publisher"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("publisher", logger) // publisherというサービス名のtracerを生成
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer) // to start the new spans, so we need to initialize that global variable to our instance of Jaeger tracer

		publisher := publisher.NewServer(
			publisherOptions,
			tracer,
			logger,
		)

		publisher.RunE()
	},
}

var publisherOptions publisher.ConfigOptions

func init() {
	rootCmd.AddCommand(publisherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publisherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publisherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	publisherOptions.Port = 8082
}
