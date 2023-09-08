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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go/v5"
)

func init() {
	rootCmd.AddCommand(queryCmd)
}

var queryCmd = &cobra.Command{
	Use:   "query [queryString]",
	Short: "Run provided Warrant query",
	Long: `
Run provided Warrant query. Examples:

warrant query 'select explicit *'`,
	Example: `
warrant query 'select explicit *'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		// TODO: Support --limit=100 --sort=createdAt --afterId=1 --beforeId=45 --afterValue=sdf --beforeValue=sdf --sortBy=sdf --sortOrder=dfsdf --limit=50 --warrantToken
		result, err := warrant.Query(args[0], &warrant.ListWarrantParams{})
		if err != nil {
			return err
		}

		jsonQueryResult, err := json.Marshal(result)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(jsonQueryResult))

		return nil
	},
}
