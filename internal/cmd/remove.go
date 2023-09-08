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
	"github.com/warrant-dev/warrant-go/v5"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [subjectType:id] [relation] [objectType:id] [policy]",
	Short: "Remove an existing warrant defined as a subject associated with an object with given relation and optional policy",
	Long: `
Remove an existing warrant defined as a subject (specified as 'type:id') associated with an object (also specified as 'type:id') with given 'relation' and optional policy. Examples:

warrant remove user:1 editor document:xyz
warrant remove user:56 member role:admin 'domain == warrant.dev'`,
	Example: `
warrant remove user:1 editor document:xyz
warrant remove user:56 member role:admin 'domain == warrant.dev'`,
	Args: cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		warrantSpec, err := reader.ReadWarrantArgs(args)
		if err != nil {
			return err
		}

		err = warrant.Delete(warrantSpec)
		if err != nil {
			return err
		}
		fmt.Printf("Deleted warrant\n")

		return nil
	},
}
