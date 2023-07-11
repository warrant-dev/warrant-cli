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
	"github.com/warrant-dev/warrant-go/v3"
	"github.com/warrant-dev/warrant-go/v3/feature"
	"github.com/warrant-dev/warrant-go/v3/permission"
	"github.com/warrant-dev/warrant-go/v3/pricingtier"
	"github.com/warrant-dev/warrant-go/v3/role"
	"github.com/warrant-dev/warrant-go/v3/tenant"
	"github.com/warrant-dev/warrant-go/v3/user"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list [type]",
	Short: "List resources of built-in object types",
	Long:  "Get a list of all resources of a given built-in object type, specifically users, tenants, roles, permissions, pricing-tiers and features.",
	Example: `
warrant list roles
warrant list tenants
warrant list permissions`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"users", "tenants", "roles", "permissions", "pricing-tiers", "features"},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}
		listParams := warrant.ListParams{
			Page:  1,
			Limit: 50,
		}
		requestedType := args[0]
		switch requestedType {
		case "users":
			users, err := user.ListUsers(&warrant.ListUserParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, user := range users {
				fmt.Println(user.UserId)
			}
			return nil
		case "tenants":
			tenants, err := tenant.ListTenants(&warrant.ListTenantParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, tenant := range tenants {
				fmt.Println(tenant.TenantId)
			}
			return nil
		case "roles":
			roles, err := role.ListRoles(&warrant.ListRoleParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, role := range roles {
				fmt.Println(role.RoleId)
			}
			return nil
		case "permissions":
			permissions, err := permission.ListPermissions(&warrant.ListPermissionParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, permission := range permissions {
				fmt.Println(permission.PermissionId)
			}
			return nil
		case "pricing-tiers":
			pricingTiers, err := pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, pricingTier := range pricingTiers {
				fmt.Println(pricingTier.PricingTierId)
			}
			return nil
		case "features":
			features, err := feature.ListFeatures(&warrant.ListFeatureParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, feature := range features {
				fmt.Println(feature.FeatureId)
			}
			return nil
		default:
			return fmt.Errorf("Invalid list request")
		}
	},
}
