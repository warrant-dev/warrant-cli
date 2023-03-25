package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v3/feature"
	"github.com/warrant-dev/warrant-go/v3/permission"
	"github.com/warrant-dev/warrant-go/v3/pricingtier"
	"github.com/warrant-dev/warrant-go/v3/role"
	"github.com/warrant-dev/warrant-go/v3/user"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [object:id] [removeFrom:id]",
	Short: "Remove an assigned object from another object (e.g. remove a role from a user)",
	Long: `
Remove an assigned object from another object. Objects are specified as [type:id] tuples. Supported removal requests include:

warrant remove user:id tenant:id
warrant remove role:id user:id
warrant remove permission:id role:id
warrant remove permission:id user:id
warrant remove feature:id pricing-tier:id
warrant remove feature:id user:id
warrant remove pricing-tier:id tenant:id
warrant remove pricing-tier:id user:id`,
	Example: `
warrant remove role:admin user:user2
warrant remove permission:perm2 role:admin`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		object, err := reader.ParseObject(args[0])
		if err != nil {
			return err
		}
		removeFrom, err := reader.ParseObject(args[1])
		if err != nil {
			return err
		}
		err = config.Init()
		if err != nil {
			return err
		}

		if object.Type == "user" && removeFrom.Type == "tenant" {
			err = user.RemoveUserFromTenant(object.Id, removeFrom.Id, "member")
		} else if object.Type == "role" && removeFrom.Type == "user" {
			err = role.RemoveRoleFromUser(object.Id, removeFrom.Id)
		} else if object.Type == "permission" && removeFrom.Type == "role" {
			err = permission.RemovePermissionFromRole(object.Id, removeFrom.Id)
		} else if object.Type == "permission" && removeFrom.Type == "user" {
			err = permission.RemovePermissionFromUser(object.Id, removeFrom.Id)
		} else if object.Type == "feature" && removeFrom.Type == "pricing-tier" {
			err = feature.RemoveFeatureFromPricingTier(object.Id, removeFrom.Id)
		} else if object.Type == "feature" && removeFrom.Type == "user" {
			err = feature.RemoveFeatureFromUser(object.Id, removeFrom.Id)
		} else if object.Type == "pricing-tier" && removeFrom.Type == "tenant" {
			err = pricingtier.RemovePricingTierFromTenant(object.Id, removeFrom.Id)
		} else if object.Type == "pricing-tier" && removeFrom.Type == "user" {
			err = pricingtier.RemovePricingTierFromUser(object.Id, removeFrom.Id)
		} else {
			return fmt.Errorf("Invalid remove request")
		}
		if err != nil {
			return err
		}
		fmt.Printf("Removed %s:%s from %s:%s\n", object.Type, object.Id, removeFrom.Type, removeFrom.Id)
		return nil
	},
}
