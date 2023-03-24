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
		err := config.InitClient()
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
