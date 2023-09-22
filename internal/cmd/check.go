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
	"os"
	"strconv"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v5"
)

var assertFlagVal string
var debug bool
var checkWarrantToken string

func init() {
	checkCmd.Flags().StringVarP(&assertFlagVal, "assert", "a", "", "execute check in 'assert' mode with an expected result. Returns 'pass' (exit:0) if the check result matches the expected result, 'fail' (exit:1) otherwise.")
	checkCmd.Flags().BoolVarP(&debug, "debug", "d", false, "run check in debug mode")
	checkCmd.Flags().StringVarP(&checkWarrantToken, "warrant-token", "w", "", "optional warrant token header value to include in check request")

	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check <subject> <relation> <object> [context]",
	Short: "Check if a subject has a given relation with an object",
	Long:  "Check if a subject (specified as 'type:id') has a given 'relation' with an object (also specified as 'type:id'). Returns 'true' if the relation exists, 'false' otherwise. Checks can also include an optional 'context' (as a json string) for policy evaluation.",
	Example: `
warrant check user:56 member role:admin
warrant check user:2 editor document:xyz
warrant check user:56 member tenant:x '{"clientIp": "192.168.0.1"}'
warrant check user:56 member role:admin --assert true`,
	Args: cobra.RangeArgs(3, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		var assertVal bool
		if assertFlagVal != "" {
			var err error
			assertVal, err = strconv.ParseBool(assertFlagVal)
			if err != nil {
				return err
			}
		}

		checkSpec, err := reader.ReadCheckArgs(args)
		if err != nil {
			return err
		}
		checkSpec.Debug = debug

		if checkWarrantToken != "" {
			checkSpec.WarrantToken = checkWarrantToken
		}

		checkResult, err := warrant.Check(checkSpec)
		if err != nil {
			return err
		}

		if assertFlagVal != "" {
			// Assert
			if checkResult == assertVal {
				fmt.Println(termenv.String(printer.Checkmark + " passed").Foreground(printer.Green))
			} else {
				fmt.Println(termenv.String(printer.Cross + " failed").Foreground(printer.Red))
				os.Exit(1)
			}
		} else {
			// Check
			if checkResult {
				fmt.Println(termenv.String(printer.Checkmark + " true").Foreground(printer.Green))
			} else {
				fmt.Println(termenv.String(printer.Cross + " false").Foreground(printer.Red))
			}
		}

		return nil
	},
}
