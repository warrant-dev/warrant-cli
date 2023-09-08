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
	"strings"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [type] [id]",
	Short: "Create a new resource given its type and desired id",
	Long:  "Create a new resource by type and id, specifically a user, tenant, role, permission, pricing-tier or feature.",
	Example: `
warrant create role new-role
warrant create permission new-perm`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		entityParts := strings.Split(args[0], ":")
		if len(entityParts) > 2 {
			return fmt.Errorf("Invalid object")
		}

		if len(args) == 2 {
			// TODO: Parse metadata
		}

		// TODO: create object

		//fmt.Printf("Created %s:%s\n", entityType, entityId)
		return nil
	},
}
