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
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v4"
)

func init() {
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(assertCmd)
	rootCmd.AddCommand(assignCmd)
	rootCmd.AddCommand(removeCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check <subject> <relation> <object> [context]",
	Short: "Check if a subject has a given relation with an object",
	Long:  "Check if a subject (specified as 'type:id') has a given 'relation' with an object (also specified as 'type:id'). Returns 'true' if the relation exists, 'false' otherwise. Checks can also include an optional 'context' (as a json string) for policy evaluation.",
	Example: `
warrant check user:56 member role:admin
warrant check user:2 editor document:xyz
warrant check user:56 member tenant:x '{"clientIp": "192.168.0.1"}'`,
	Args: cobra.RangeArgs(3, 4),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		checkSpec, err := reader.ReadCheckArgs(args)
		if err != nil {
			return err
		}

		result, err := warrant.Check(checkSpec)
		if err != nil {
			return err
		}
		fmt.Printf("%t\n", result)

		return nil
	},
}

var assertCmd = &cobra.Command{
	Use:   "assert {true|false} <subject> <relation> <object> [context]",
	Short: "Assert whether a subject has a given relation with an object",
	Long:  "Assert whether a subject (specified as 'type:id') has a given 'relation' with an object (also specified as 'type:id'). Returns 'true' if the assertion matches the provided expectation, 'false' otherwise. Like 'check', an assertion can also include an optional 'context' (as a json string) for policy evaluation.",
	Example: `
warrant assert true user:56 member role:admin
warrant assert false user:2 editor document:xyz
warrant assert true user:56 member tenant:x '{"clientIp": "192.168.0.1"}'`,
	Args: cobra.RangeArgs(4, 5),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		expected, err := strconv.ParseBool(args[0])
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

		if result == expected {
			fmt.Printf("%t\n", true)
		} else {
			fmt.Printf("%t\n", false)
		}

		return nil
	},
}

var assignCmd = &cobra.Command{
	Use:   "assign <subject> <relation> <object> [policy]",
	Short: "Assign a subject to an object with given relation and an optional policy string",
	Long:  "Assign a subject (specified as 'type:id') to an object (also specified as 'type:id') with given 'relation' and optional 'policy' string.",
	Example: `
warrant assign user:1 editor document:xyz
warrant assign user:56 member role:admin 'domain == warrant.dev'`,
	Args: cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		warrantSpec, err := reader.ReadWarrantArgs(args)
		if err != nil {
			return err
		}

		newWarrant, err := warrant.Create(warrantSpec)
		if err != nil {
			return err
		}
		fmt.Printf("Created warrant %s:%s %s %s:%s\n", newWarrant.ObjectType, newWarrant.ObjectId, newWarrant.Relation, newWarrant.Subject.ObjectType, newWarrant.Subject.ObjectId)

		return nil
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove <subject> <relation> <object> [policy]",
	Short: "Remove an existing warrant defined as a subject associated with an object with given relation and an optional policy string",
	Long:  "Remove an existing warrant defined as a subject (specified as 'type:id') associated with an object (also specified as 'type:id') with given 'relation' and optional 'policy' string.",
	Example: `
warrant remove user:1 editor document:xyz
warrant remove user:56 member role:admin 'domain == warrant.dev'`,
	Args: cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

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
