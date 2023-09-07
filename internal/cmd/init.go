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
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

type ConfigFile struct {
	Key         string `json:"key"`
	ApiEndpoint string `json:"apiEndpoint"`
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI for use",
	Long:  "Initialize the CLI for use, including configuring server endpoint and API key.",
	Example: `
warrant init`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Warrant endpoint override (leave blank to use https://api.warrant.dev default):")
		fmt.Print("> ")
		buf := bufio.NewReader(os.Stdin)
		input, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		endpoint := strings.TrimSpace(string(input))
		if endpoint == "" {
			endpoint = "https://api.warrant.dev"
		}
		fmt.Println("API Key:")
		fmt.Print("> ")
		buf = bufio.NewReader(os.Stdin)
		input, err = buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		key := strings.TrimSpace(string(input))

		fmt.Println("Creating ~/.warrant.json")
		config := ConfigFile{
			ApiEndpoint: endpoint,
			Key:         key,
		}
		fileContents, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		err = os.WriteFile(homeDir+"/.warrant.json", fileContents, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Setup complete.")
		return nil
	},
}
