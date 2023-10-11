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
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI for use",
	Long:  "Initialize the CLI for use, including configuring an environment and API key.",
	Example: `
warrant init`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		envName, env, err := reader.ReadEnvFromConsole()
		if err != nil {
			return err
		}

		fmt.Println("creating ~/.warrant.json")
		envMap := make(map[string]config.Environment)
		envMap[envName] = *env
		newConfig := config.Config{
			ActiveEnvironment: envName,
			Environments:      envMap,
		}
		err = newConfig.Write()
		if err != nil {
			return err
		}
		fmt.Println("setup complete")

		return nil
	},
}
