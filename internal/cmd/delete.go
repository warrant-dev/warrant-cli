package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete role|tenant|permission|user ID",
	Short: "Delete a resource given its type and id",
	Long:  "Delete a resource by id, specifically a role, tenant, permission or user.",
	Example: `
warrant delete role admin-1
warrant delete tenant tenant-2
warrant delete permission perm-1
warrant delete user user-1`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("accepts 2 args, received %d", len(args))
		}
		entityType := args[0]
		if entityType != "role" && entityType != "tenant" && entityType != "permission" && entityType != "user" {
			return fmt.Errorf("entity to delete must be a role|tenant|permission|user")
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
			err = client.DeleteRole(entityId)
		case "tenant":
			err = client.DeleteTenant(entityId)
		case "permission":
			err = client.DeletePermission(entityId)
		case "user":
			err = client.DeleteUser(entityId)
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
