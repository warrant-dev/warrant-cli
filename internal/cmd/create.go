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
	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/feature"
	"github.com/warrant-dev/warrant-go/v5/permission"
	"github.com/warrant-dev/warrant-go/v5/pricingtier"
	"github.com/warrant-dev/warrant-go/v5/role"
	"github.com/warrant-dev/warrant-go/v5/tenant"
	"github.com/warrant-dev/warrant-go/v5/user"
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
		entityType := args[0]
		entityId := args[1]
		switch entityType {
		case "user":
			_, err = user.Create(&warrant.UserParams{
				UserId: entityId,
			})
		case "tenant":
			_, err = tenant.Create(&warrant.TenantParams{
				TenantId: entityId,
			})
		case "role":
			_, err = role.Create(&warrant.RoleParams{
				RoleId: entityId,
			})
		case "permission":
			_, err = permission.Create(&warrant.PermissionParams{
				PermissionId: entityId,
			})
		case "pricing-tier":
			_, err = pricingtier.Create(&warrant.PricingTierParams{
				PricingTierId: entityId,
			})
		case "feature":
			_, err = feature.Create(&warrant.FeatureParams{
				FeatureId: entityId,
			})
		default:
			return fmt.Errorf("Invalid create request")
		}
		if err != nil {
			return err
		}
		fmt.Printf("Created %s:%s\n", entityType, entityId)
		return nil
	},
}
