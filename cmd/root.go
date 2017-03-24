// Copyright © 2016 Phil Estes
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
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logLevel string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "elastistack",
	Short: "Import Golang stack trace data into Elasticsearch",
	Long: `This program can parse standard Go stack trace dump
text data and import it as structed text to a running Elasticsearch
instance.

The parsing code will ignore non-stack trace content in the input
file, so there is no need to clean up a larger logfile or standard
output that includes both a Go stack trace as well as other data.`,
	SilenceErrors: false,
	SilenceUsage: false,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initLogLevel)
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "warn", "set the logging level (info,warn,err,debug)")
}

func initLogLevel() {
	switch logLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
		return
	case "warn":
		log.SetLevel(log.WarnLevel)
		return
	case "err":
		log.SetLevel(log.ErrorLevel)
		return
	case "debug":
		log.SetLevel(log.DebugLevel)
		return
	}
	fmt.Printf("Invalid log level specified: %q\n", logLevel)
	os.Exit(1)
}
