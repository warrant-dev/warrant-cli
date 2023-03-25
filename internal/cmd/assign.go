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
	rootCmd.AddCommand(assignCmd)
}

var assignCmd = &cobra.Command{
	Use:   "assign [object:id] [assignTo:id]",
	Short: "Assign an object to another object (e.g. assign a role to a user)",
	Long: `
Assign an object to another object. Objects are specified as [type:id] tuples. Supported assignments include:

warrant assign user:id tenant:id
warrant assign role:id user:id
warrant assign permission:id role:id
warrant assign permission:id user:id
warrant assign feature:id pricing-tier:id
warrant assign feature:id user:id
warrant assign pricing-tier:id tenant:id
warrant assign pricing-tier:id user:id`,
	Example: `
warrant assign role:admin user:user2
warrant assign permission:perm2 role:admin`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		object, err := reader.ParseObject(args[0])
		if err != nil {
			return err
		}
		assignTo, err := reader.ParseObject(args[1])
		if err != nil {
			return err
		}
		err = config.Init()
		if err != nil {
			return err
		}

		if object.Type == "user" && assignTo.Type == "tenant" {
			_, err = user.AssignUserToTenant(object.Id, assignTo.Id, "member")
		} else if object.Type == "role" && assignTo.Type == "user" {
			_, err = role.AssignRoleToUser(object.Id, assignTo.Id)
		} else if object.Type == "permission" && assignTo.Type == "role" {
			_, err = permission.AssignPermissionToRole(object.Id, assignTo.Id)
		} else if object.Type == "permission" && assignTo.Type == "user" {
			_, err = permission.AssignPermissionToUser(object.Id, assignTo.Id)
		} else if object.Type == "feature" && assignTo.Type == "pricing-tier" {
			_, err = feature.AssignFeatureToPricingTier(object.Id, assignTo.Id)
		} else if object.Type == "feature" && assignTo.Type == "user" {
			_, err = feature.AssignFeatureToUser(object.Id, assignTo.Id)
		} else if object.Type == "pricing-tier" && assignTo.Type == "tenant" {
			_, err = pricingtier.AssignPricingTierToTenant(object.Id, assignTo.Id)
		} else if object.Type == "pricing-tier" && assignTo.Type == "user" {
			_, err = pricingtier.AssignPricingTierToUser(object.Id, assignTo.Id)
		} else {
			return fmt.Errorf("Invalid assign request")
		}
		if err != nil {
			return err
		}
		fmt.Printf("Assigned %s:%s to %s:%s\n", object.Type, object.Id, assignTo.Type, assignTo.Id)
		return nil
	},
}
