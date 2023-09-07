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
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check [type:id] [hasPermission|hasRole|hasFeature|relation] [id|type:id]",
	Short: "Check if an object (specified as type:id) has a given permission, role, feature or relationship to another object",
	Long: `
Check if an object (specified as type:id) has a given permission, role, feature or relationship to another object. Supported checks include:

warrant check user:id hasPermission perm1
warrant check user:id hasRole admin
warrant check user:id hasFeature feature1
warrant check user:id member tenant:id`,
	Example: `
warrant check user:56 hasPermission view-dashboards
warrant check user:45 hasRole admin
warrant check user:1 hasFeature dashboard
warrant check user:2 editor document:xyz`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.Init()
		if err != nil {
			return err
		}
		subject, err := reader.ParseObject(args[0])
		if err != nil {
			return err
		}
		toCheck := args[1]

		result := false
		if toCheck == "hasPermission" && subject.Type == "user" {
			result, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
				PermissionId: args[2],
				UserId:       subject.Id,
			})
		} else if toCheck == "hasRole" && subject.Type == "user" {
			result, err = warrant.CheckUserHasRole(&warrant.RoleCheckParams{
				RoleId: args[2],
				UserId: subject.Id,
			})
		} else if toCheck == "hasFeature" {
			result, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
				FeatureId: args[2],
				Subject: warrant.Subject{
					ObjectType: subject.Type,
					ObjectId:   subject.Id,
				},
			})
		} else {
			object, e := reader.ParseObject(args[2])
			if e != nil {
				return e
			}
			result, err = warrant.Check(&warrant.WarrantCheckParams{
				WarrantCheck: warrant.WarrantCheck{
					Object: warrant.Object{
						ObjectType: object.Type,
						ObjectId:   object.Id,
					},
					Relation: toCheck,
					Subject: warrant.Subject{
						ObjectType: subject.Type,
						ObjectId:   subject.Id,
					},
				},
			})
		}
		if err != nil {
			return err
		}
		fmt.Printf("%t\n", result)
		return nil
	},
}
