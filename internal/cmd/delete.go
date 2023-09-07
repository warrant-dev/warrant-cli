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
	"github.com/warrant-dev/warrant-go/v5/feature"
	"github.com/warrant-dev/warrant-go/v5/permission"
	"github.com/warrant-dev/warrant-go/v5/pricingtier"
	"github.com/warrant-dev/warrant-go/v5/role"
	"github.com/warrant-dev/warrant-go/v5/tenant"
	"github.com/warrant-dev/warrant-go/v5/user"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [type] [id]",
	Short: "Delete a resource given its type and id",
	Long:  "Delete a resource by id, specifically a user, tenant, role, permission, pricing-tier or feature.",
	Example: `
warrant delete role admin-1
warrant delete tenant tenant-2
warrant delete permission perm-1
warrant delete user user-1`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}
		entityType := args[0]
		entityId := args[1]
		switch entityType {
		case "user":
			err = user.Delete(entityId)
		case "tenant":
			err = tenant.Delete(entityId)
		case "role":
			err = role.Delete(entityId)
		case "permission":
			err = permission.Delete(entityId)
		case "pricing-tier":
			err = pricingtier.Delete(entityId)
		case "feature":
			err = feature.Delete(entityId)
		default:
			return fmt.Errorf("Invalid delete request")
		}
		if err != nil {
			return err
		}
		fmt.Printf("Deleted %s:%s\n", entityType, entityId)
		return nil
	},
}
