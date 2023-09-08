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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v5"
)

func init() {
	rootCmd.AddCommand(assertCmd)
}

var assertCmd = &cobra.Command{
	Use:   "assert [true|false] [subjectType:id] [relation] [objectType:id] [context]",
	Short: "Assert whether a given check is true or false",
	Long: `
Assert whether a given check with subject (specified as 'type:id') has a 'relation' with an object (also specified as 'type:id'). Checks can also include an optional context for policy evaluation. Example assertions:

warrant assert true user:56 member role:admin # returns true if check passes
warrant assert false user:2 editor document:xyz # returns true if check does not pass
warrant assert true user:56 member tenant:x '{"clientIp": "192.168.0.1"}' # returns true if check passes`,
	Example: `
warrant assert true user:56 member role:admin
warrant assert false user:2 editor document:xyz
warrant assert true user:56 member tenant:x '{"clientIp": "192.168.0.1"}'`,
	Args: cobra.RangeArgs(4, 5),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		assertVal, err := strconv.ParseBool(args[0])
		if err != nil {
			return err
		}

		checkSpec, err := reader.ReadCheckArgs(args[1:])
		if err != nil {
			return err
		}

		result, err := warrant.Check(checkSpec)
		if err != nil {
			return err
		}

		if result == assertVal {
			fmt.Printf("%s%t\n", Green, true)
		} else {
			fmt.Printf("%s%t\n", Red, false)
		}

		return nil
	},
}
