// Copyright 2023 Forerunner Labs, Inc.
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
	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-go/v5"
)

var cmdConfig *config.Config

var rootCmd = &cobra.Command{
	Use:   "warrant",
	Short: "Warrant CLI",
	Long:  `The Warrant CLI is a tool to interact with Warrant via the command line.`,
}

func SetVersion(version string) {
	rootCmd.Version = version
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.warrant.json)")
	// rootCmd.PersistentFlags().StringVarP(&EnvName, "env", "e", "", "environment")
	// viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	cmdConfig = config.LoadConfig()
	warrant.ApiKey = cmdConfig.Environments[cmdConfig.ActiveEnvironment].ApiKey
	warrant.ApiEndpoint = cmdConfig.Environments[cmdConfig.ActiveEnvironment].ApiEndpoint
}

func GetConfigOrExit() *config.Config {
	if cmdConfig.ActiveEnvironment == "" {
		printer.PrintErrAndExit("no active environment configured. Run 'warrant init'")
	}
	if len(cmdConfig.Environments) == 0 {
		printer.PrintErrAndExit("no environments configured. Run 'warrant init'")
	}
	if _, ok := cmdConfig.Environments[cmdConfig.ActiveEnvironment]; !ok {
		printer.PrintErrAndExit("invalid active environment configured. Run 'warrant init'")
	}
	return cmdConfig
}
