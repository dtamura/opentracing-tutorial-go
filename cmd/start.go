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
	"github.com/dtamura/hello-cobra/lib/log"
	"github.com/dtamura/hello-cobra/lib/tracing"
	"github.com/dtamura/hello-cobra/services/sandbox"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var options = &sandbox.ConfigOptions{}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		zapLogger := logger.With(zap.String("service", "sandbox"))
		logger := log.NewFactory(zapLogger)
		tracer, closer := tracing.Init("hello-world", logger)
		server := sandbox.NewServer(
			options,
			tracer,
			logger,
			closer,
		)
		defer closer.Close()
		err := server.RunE()
		return err
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.PersistentFlags().StringVarP(&options.Message, "message", "m", "tamura", "Message")

}
