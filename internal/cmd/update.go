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
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update [type] [id]",
	Short: "Delete a resource given its type and id",
	Long:  "Delete a resource by id, specifically a user, tenant, role, permission, pricing-tier or feature.",
	Example: `
warrant delete role admin-1
warrant delete tenant tenant-2
warrant delete permission perm-1
warrant delete user user-1`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		_, err = reader.ParseObject(args[0])
		if err != nil {
			return err
		}

		// TODO: get object by id
		//fmt.Printf("Deleted %s:%s\n", entityType, entityId)
		return nil
	},
}
