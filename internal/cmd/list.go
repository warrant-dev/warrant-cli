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
		err := config.InitClient()
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
