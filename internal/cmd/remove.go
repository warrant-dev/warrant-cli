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
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove subject:id object:id",
	Short: "Remove an assigned 'subject' from an object (ex. remove a role from a user)",
	Long:  "Remove an assigned 'subject' from an object. Currently supports removing roles from users, permissions from users, permissions from roles and users from tenants.",
	Example: `
warrant remove role:admin user:user2
warrant remove permission:perm2 role:admin`,
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
			err = role.RemoveRoleFromUser(subject.Id, object.Id)
		} else if subject.Type == "permission" && object.Type == "user" {
			err = permission.RemovePermissionFromUser(subject.Id, object.Id)
		} else if subject.Type == "permission" && object.Type == "role" {
			err = permission.RemovePermissionFromRole(subject.Id, object.Id)
		} else if subject.Type == "user" && object.Type == "tenant" {
			err = user.RemoveUserFromTenant(subject.Id, object.Id, "member")
		} else {
			return fmt.Errorf("Invalid remove request")
		}
		if err != nil {
			return err
		}
		fmt.Printf("Removed %s:%s from %s:%s\n", subject.Type, subject.Id, object.Type, object.Id)
		return nil
	},
}
