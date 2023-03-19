package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go/v3/permission"
	"github.com/warrant-dev/warrant-go/v3/role"
	"github.com/warrant-dev/warrant-go/v3/tenant"
	"github.com/warrant-dev/warrant-go/v3/user"
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
		err := config.InitClient()
		if err != nil {
			return err
		}
		entityType := args[0]
		entityId := args[1]
		switch entityType {
		case "role":
			err = role.Delete(entityId)
		case "tenant":
			err = tenant.Delete(entityId)
		case "permission":
			err = permission.Delete(entityId)
		case "user":
			err = user.Delete(entityId)
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
