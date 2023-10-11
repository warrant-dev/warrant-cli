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
	"strings"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/object"
)

func init() {
	objectCmd.AddCommand(createCmd)
	objectCmd.AddCommand(getCmd)
	objectCmd.AddCommand(updateCmd)
	objectCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(objectCmd)
}

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Operate on objects (create, get, update, delete)",
	Long:  "Operate on objects (create, get, update, delete), including their metadata.",
	Example: `
warrant object create role:admin
warrant object get role:admin
warrant object update role:admin '{"name": "New name"}'
warrant object delete role:admin`,
}

var createCmd = &cobra.Command{
	Use:   "create <object> [meta]",
	Short: "Create a new object of specified type with optional id and optional meta",
	Long:  "Create a new object of specified type with optional id and optional meta. If an id is provided (e.g. 'role:123'), it will be assigned to the newly created object. The optional 'meta' is provided as a json string and will be attached to the newly created object.",
	Example: `
warrant object create role
warrant object create user:123
warrant object create permission:edit-users '{"name": "Edit Users"}'`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		typeAndId := strings.Split(args[0], ":")
		if len(typeAndId) > 2 {
			printer.PrintErrAndExit("invalid object provided, must be 'type' or 'type:id'")
		}
		objectType := typeAndId[0]
		objectId := ""
		if len(typeAndId) == 2 {
			objectId = typeAndId[1]
		}

		var meta map[string]interface{}
		var err error
		if len(args) == 2 {
			meta, err = reader.ReadObjectMetaArg(args[1])
			if err != nil {
				return err
			}
		}

		newObj, err := object.Create(&warrant.ObjectParams{
			ObjectType: objectType,
			ObjectId:   objectId,
			Meta:       meta,
		})
		if err != nil {
			return err
		}

		fmt.Printf("created %s:%s\n", newObj.ObjectType, newObj.ObjectId)
		if len(newObj.Meta) > 0 {
			printer.PrintJson(newObj.Meta)
		}

		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get <object>",
	Short: "Get an object specified by type:id",
	Long:  "Get an object specified by type:id. Also returns the object's 'meta', if present.",
	Example: `
warrant object get role:123`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		objectType, objectId, err := reader.ReadObjectArg(args[0])
		if err != nil {
			return err
		}

		obj, err := object.Get(objectType, objectId, &warrant.ObjectParams{})
		if err != nil {
			return err
		}

		fmt.Printf("%s:%s\n", obj.ObjectType, obj.ObjectId)
		if len(obj.Meta) > 0 {
			printer.PrintJson(obj.Meta)
		}

		return nil
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <object> <meta>",
	Short: "Update an object's (specified as type:id) meta",
	Long:  "Update an object's (specified as type:id) meta. Object 'meta' must be passed as a json string. Note that an object's existing type and id cannot be updated.",
	Example: `
warrant object update role:123 '{"name": "New name"}'`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		objectType, objectId, err := reader.ReadObjectArg(args[0])
		if err != nil {
			return err
		}

		meta, err := reader.ReadObjectMetaArg(args[1])
		if err != nil {
			return err
		}

		updatedObj, err := object.Update(objectType, objectId, &warrant.ObjectParams{
			Meta: meta,
		})
		if err != nil {
			return err
		}

		fmt.Printf("updated %s:%s\n", updatedObj.ObjectType, updatedObj.ObjectId)
		if len(updatedObj.Meta) > 0 {
			printer.PrintJson(updatedObj.Meta)
		}

		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <object>",
	Short: "Delete the object with specified type:id",
	Long:  "Delete the object with specified type:id. The entire object, including its 'meta', will be deleted.",
	Example: `
warrant object delete role:admin`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		objectType, objectId, err := reader.ReadObjectArg(args[0])
		if err != nil {
			return err
		}

		err = object.Delete(objectType, objectId)
		if err != nil {
			return err
		}

		fmt.Printf("deleted %s:%s\n", objectType, objectId)

		return nil
	},
}
