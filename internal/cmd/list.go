package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go/v3"
	"github.com/warrant-dev/warrant-go/v3/permission"
	"github.com/warrant-dev/warrant-go/v3/role"
	"github.com/warrant-dev/warrant-go/v3/tenant"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list roles|tenants|permissions",
	Short: "List resources of built-in object types",
	Long:  "Get a list of all resources of a given built-in object type, specifically roles, tenants and permissions.",
	Example: `
warrant list roles
warrant list tenants
warrant list permissions`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"roles", "tenants", "permissions"},
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
		case "tenants":
			tenants, err := tenant.ListTenants(&warrant.ListTenantParams{
				ListParams: listParams,
			})
			if err != nil {
				return err
			}
			for _, tenant := range tenants {
				fmt.Printf("%s: %s\n", tenant.TenantId, tenant.Name)
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
		default:
			return fmt.Errorf("Invalid list request")
		}
	},
}
