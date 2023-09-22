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
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-go/v5"
)

var queryWarrantToken string

func init() {
	queryCmd.Flags().StringVarP(&queryWarrantToken, "warrant-token", "w", "", "optional warrant token header value to include in query request")

	rootCmd.AddCommand(queryCmd)
}

var queryCmd = &cobra.Command{
	Use:   "query <queryString>",
	Short: "Run a provided Warrant query",
	Long:  "Run a provided Warrant query.",
	Example: `
warrant query 'select explicit *'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		queryParams := &warrant.QueryParams{}
		if queryWarrantToken != "" {
			queryParams.RequestOptions = warrant.RequestOptions{
				WarrantToken: queryWarrantToken,
			}
		}
		result, err := warrant.Query(args[0], &warrant.QueryParams{})
		if err != nil {
			return err
		}

		printer.PrintJson(result)

		return nil
	},
}
