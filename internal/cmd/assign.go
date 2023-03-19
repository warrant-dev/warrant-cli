package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v3/permission"
	"github.com/warrant-dev/warrant-go/v3/role"
	"github.com/warrant-dev/warrant-go/v3/user"
)

func init() {
	rootCmd.AddCommand(assignCmd)
}

var assignCmd = &cobra.Command{
	Use:   "assign subject:id object:id",
	Short: "Assign a 'subject' to another object (ex. assign a role to a user)",
	Long:  "Assign a 'subject' to another object. Currently supported assignments include assigning roles to users, permissions to users, permissions to roles and users to tenants.",
	Example: `
warrant assign role:admin user:user2
warrant assign permission:perm2 role:admin`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		subject, err := reader.ParseObject(args[0])
		if err != nil {
			return err
		}
		object, err := reader.ParseObject(args[1])
		if err != nil {
			return err
		}
		err = config.InitClient()
		if err != nil {
			return err
		}
		if subject.Type == "role" && object.Type == "user" {
			_, err = role.AssignRoleToUser(subject.Id, object.Id)
		} else if subject.Type == "permission" && object.Type == "user" {
			_, err = permission.AssignPermissionToUser(subject.Id, object.Id)
		} else if subject.Type == "permission" && object.Type == "role" {
			_, err = permission.AssignPermissionToRole(subject.Id, object.Id)
		} else if subject.Type == "user" && object.Type == "tenant" {
			_, err = user.AssignUserToTenant(subject.Id, object.Id, "member")
		} else {
			return fmt.Errorf("Invalid assign request")
		}
		if err != nil {
			return err
		}
		fmt.Printf("Assigned %s:%s to %s:%s\n", subject.Type, subject.Id, object.Type, object.Id)
		return nil
	},
}
