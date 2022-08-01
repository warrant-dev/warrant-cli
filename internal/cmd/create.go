package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create role|permission ID",
	Short: "Create a new resource given its type and desired id",
	Long:  "Create a new resource by type and id, specifically a role or permission.",
	Example: `
warrant create role new-role
warrant create permission new-perm`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("accepts 2 args, received %d", len(args))
		}
		entityType := args[0]
		if entityType != "role" && entityType != "permission" {
			return fmt.Errorf("entity to delete must be a role|permission")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := config.GetClient()
		if err != nil {
			return err
		}
		entityType := args[0]
		entityId := args[1]
		switch entityType {
		case "role":
			_, err = client.CreateRole(warrant.Role{
				RoleId: entityId,
			})
		case "permission":
			_, err = client.CreatePermission(warrant.Permission{
				PermissionId: entityId,
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
