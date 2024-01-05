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
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/printer"
	"github.com/warrant-dev/warrant-go/v6"
	"github.com/warrant-dev/warrant-go/v6/objecttype"
)

var listObjecttypeWarrantToken string
var typesFile string

func init() {
	listObjecttypeCmd.Flags().StringVarP(&listObjecttypeWarrantToken, "warrant-token", "w", "", "optional warrant token header value to include in list objecttypes request")
	applyObjecttypeCmd.Flags().StringVarP(&typesFile, "file", "f", "", "file containing object type definitions")

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
	Short: "List all object types in active environment",
	Long:  "List all object types in active environment.",
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

		// Fetch all objecttypes (paginate if necessary)
		var types []warrant.ObjectType
		for {
			typesResp, err := objecttype.ListObjectTypes(listParams)
			if err != nil {
				return err
			}
			types = append(types, typesResp.Results...)

			if typesResp.NextCursor == "" {
				break
			} else {
				listParams.NextCursor = typesResp.NextCursor
			}
		}
		printer.PrintJson(types)

		return nil
	},
}

var applyObjecttypeCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply updated object types configuration to active environment",
	Long:  "Apply updated object types configuration to active environment. New object type definitions can be provided via file (-f) or stdin.",
	Example: `
warrant objecttype apply -f types.json`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		GetConfigOrExit()

		var bytes []byte
		var err error
		if typesFile != "" {
			// Read from file if filename provided
			jsonFile, err := os.Open(typesFile)
			if err != nil {
				return err
			}
			defer jsonFile.Close()

			bytes, err = io.ReadAll(jsonFile)
			if err != nil {
				return err
			}
		} else {
			// Else read from stdin
			bytes, err = io.ReadAll(bufio.NewReader(os.Stdin))
			if err != nil {
				return err
			}
		}

		var objectTypes []warrant.ObjectTypeParams
		err = json.Unmarshal(bytes, &objectTypes)
		if err != nil {
			return err
		}

		_, err = objecttype.BatchUpdate(objectTypes)
		if err != nil {
			return err
		}

		fmt.Println("objecttypes updated")

		return nil
	},
}
