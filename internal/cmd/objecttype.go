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
	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/objecttype"
)

var typesFile string
var listObjecttypeWarrantToken string

func init() {
	applyObjecttypeCmd.Flags().StringVarP(&typesFile, "file", "f", "", "file containing object type definitions")
	listObjecttypeCmd.Flags().StringVarP(&listObjecttypeWarrantToken, "warrant-token", "w", "", "optional warrant token header value to include in list objecttypes request")

	objecttypeCmd.AddCommand(listObjecttypeCmd)
	objecttypeCmd.AddCommand(applyObjecttypeCmd)
	rootCmd.AddCommand(objecttypeCmd)
}

var objecttypeCmd = &cobra.Command{
	Use:   "objecttype",
	Short: "Operate on object type definitions",
	Long:  "Operate on object type definitions, including listing existing object types and applying new configuration.",
	Example: `
warrant objecttype list
warrant objecttype apply -f types.json`,
}

var listObjecttypeCmd = &cobra.Command{
	Use:   "list",
	Short: "List all object types in environment",
	Long:  "List all object types in environment.",
	Example: `
warrant objecttype list`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		listParams := &warrant.ListObjectTypeParams{}
		if listObjecttypeWarrantToken != "" {
			listParams.RequestOptions = warrant.RequestOptions{
				WarrantToken: listObjecttypeWarrantToken,
			}
		}
		types, err := objecttype.ListObjectTypes(listParams)
		if err != nil {
			return err
		}
		printer.PrintJson(types)

		return nil
	},
}

var applyObjecttypeCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply updated object types configuration",
	Long:  "Apply updated object types configuration. New object type definitions can be provided via file (-f) or stdin.",
	Example: `
warrant objecttype apply -f types.json`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		return nil
	},
}
