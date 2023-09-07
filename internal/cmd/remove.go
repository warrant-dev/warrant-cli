// Copyright 2023 Forerunner Labs, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/feature"
	"github.com/warrant-dev/warrant-go/v5/permission"
	"github.com/warrant-dev/warrant-go/v5/pricingtier"
	"github.com/warrant-dev/warrant-go/v5/role"
	"github.com/warrant-dev/warrant-go/v5/user"
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
warrant remove pricing-tier:id user:id
warrant remove subject:id relation object:id`,
	Example: `
warrant remove role:admin user:user2
warrant remove permission:perm2 role:admin
warrant remove user:1 editor document:xyz`,
	Args: cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}

		if len(args) == 3 {
			// Delete warrant (subject, relation, object)
			subject, err := reader.ParseObject(args[0])
			if err != nil {
				return err
			}
			relation := args[1]
			object, err := reader.ParseObject(args[2])
			if err != nil {
				return err
			}
			err = warrant.Delete(&warrant.WarrantParams{
				ObjectType: object.Type,
				ObjectId:   object.Id,
				Relation:   relation,
				Subject: warrant.Subject{
					ObjectType: subject.Type,
					ObjectId:   subject.Id,
				},
			})
			if err != nil {
				return err
			}
			fmt.Printf("Deleted warrant %s:%s %s %s:%s\n", subject.Type, subject.Id, relation, object.Type, object.Id)
		} else {
			// Remove built-in type associations
			object, err := reader.ParseObject(args[0])
			if err != nil {
				return err
			}
			removeFrom, err := reader.ParseObject(args[1])
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
		}
		return nil
	},
}
