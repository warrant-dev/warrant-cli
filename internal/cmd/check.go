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
	rootCmd.AddCommand(checkCmd)
}

// TODO: support 'debug' flag
var checkCmd = &cobra.Command{
	Use:   "check [subjectType:id] [relation] [objectType:id] [context]",
	Short: "Check if a subject has a given relation with an object",
	Long: `
Check if a subject (specified as 'type:id') has a given 'relation' with an object (also specified as 'type:id'). Checks can also include an optional context for policy evaluation. Example checks:

warrant check user:56 member role:admin
warrant check user:2 editor document:xyz
warrant check user:56 member tenant:x '{"clientIp": "192.168.0.1"}'`,
	Example: `
warrant check user:56 member role:admin
warrant check user:2 editor document:xyz
warrant check user:56 member tenant:x '{"clientIp": "192.168.0.1"}'`,
	Args: cobra.RangeArgs(3, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		checkSpec, err := reader.ReadCheckArgs(args)
		if err != nil {
			return err
		}

		result, err := warrant.Check(checkSpec)
		if err != nil {
			return err
		}

		if result {
			fmt.Printf("%s%t\n", Green, result)
		} else {
			fmt.Printf("%s%t\n", Red, result)
		}

		return nil
	},
}
