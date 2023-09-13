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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-cli/internal/reader"
)

func init() {
	envCmd.AddCommand(listEnvCmd)
	envCmd.AddCommand(addEnvCmd)
	envCmd.AddCommand(removeEnvCmd)
	envCmd.AddCommand(switchEnvCmd)
	rootCmd.AddCommand(envCmd)
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Get the name of the current active environment",
	Long:  "Get the name of the current active environment.",
	Example: `
warrant env`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := GetConfigOrExit()
		fmt.Println(config.ActiveEnvironment)

		return nil
	},
}

var listEnvCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured environments",
	Long:  "List all configured environments, including the current active environment denoted by a * prefix.",
	Example: `
warrant list`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := GetConfigOrExit()
		for env := range config.Environments {
			if env == config.ActiveEnvironment {
				fmt.Println("* " + env)
			} else {
				fmt.Println("  " + env)
			}
		}

		return nil
	},
}

var addEnvCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new environment to config",
	Long:  "Add a new environment to config, including its API key and API endpoint.",
	Example: `
warrant add`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := GetConfigOrExit()
		envToAdd, newEnv, err := reader.ReadEnvFromConsole()
		if err != nil {
			return err
		}
		config.Environments[envToAdd] = *newEnv
		err = config.Write()
		if err != nil {
			return err
		}
		fmt.Printf("Added environment '%s'\n", envToAdd)

		return nil
	},
}

var removeEnvCmd = &cobra.Command{
	Use:   "remove <envName>",
	Short: "Remove an existing environment from config",
	Long:  "Remove an existing environment from config, provided it exists and is not currently active.",
	Example: `
warrant remove test`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := GetConfigOrExit()
		envToRemove := args[0]
		if envToRemove == config.ActiveEnvironment {
			printer.PrintErrAndExit("cannot remove active environment")
		}
		if _, ok := config.Environments[envToRemove]; !ok {
			printer.PrintErrAndExit(fmt.Sprintf("environment '%s' does not exist", envToRemove))
		}
		delete(config.Environments, envToRemove)
		err := config.Write()
		if err != nil {
			return err
		}
		fmt.Printf("Removed environment '%s'\n", envToRemove)

		return nil
	},
}

var switchEnvCmd = &cobra.Command{
	Use:   "switch <envName>",
	Short: "Switch to the given environment",
	Long:  "Switch to the given environment, provided it exists in config.",
	Example: `
warrant switch prod`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config := GetConfigOrExit()
		env := args[0]
		if _, ok := config.Environments[env]; !ok {
			printer.PrintErrAndExit(fmt.Sprintf("environment '%s' does not exist", env))
		}
		config.ActiveEnvironment = env
		err := config.Write()
		if err != nil {
			return err
		}
		fmt.Printf("Switched to environment '%s'\n", env)

		return nil
	},
}