package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-go"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import users filename",
	Short: "Import entities from a csv file",
	Long:  "Import entities (currently supports users) from a csv file.",
	Example: `
warrant import users file.csv`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 && len(args) != 2 {
			return fmt.Errorf("must provide 2 args for import: 'entityType' and 'filename'")
		}
		if len(args) > 0 {
			entityType := args[0]
			if entityType != "users" {
				return fmt.Errorf("first arg must be one of: users")
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Return info if run with no args
		if len(args) == 0 {
			fmt.Printf(`
This cmd can be used to bulk import entities from a csv file directly into Warrant.

1. Start by creating a csv file containing the entities you'd like to import (one entity per row). Each file must only contain 1 type of entity and all required attributes must be provided in column order (see columns below).
Please include the columns for your entityType as the first row (header) in your file:

entityType --> [columns] (optional attributes indicated by *)

users --> [userId, email*]

2. Once you have created your csv file, you can import it via this cmd:

warrant import [entityType] [filename]

For example:

warrant import users users.csv
`)
			return nil
		}

		// Import from csv
		client, err := config.GetClient()
		if err != nil {
			return err
		}
		entityType := args[0]
		fileName := args[1]

		// open & read entities from csv file
		fmt.Printf("Reading file...\n")
		f, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			return err
		}

		// import entities
		switch entityType {
		case "users":
			importUsers(data, client)
		default:
			return fmt.Errorf("Invalid import")
		}
		return nil
	},
}

func importUsers(data [][]string, client warrant.WarrantClient) {
	fmt.Printf("Creating users...\n")
	rowsProcessed := 0
	usersCreated := 0
	usersFailed := 0
	for i, line := range data {
		if i > 0 { // omit header line
			var newUser warrant.User
			for j, field := range line {
				if j == 0 {
					newUser.UserId = strings.TrimSpace(field)
				} else if j == 1 {
					newUser.Email = strings.TrimSpace(field)
				}
			}
			_, err := client.CreateUser(newUser)
			if err != nil {
				usersFailed++
				fmt.Printf("Error processing row %d: %s\n", i, err)
			} else {
				usersCreated++
			}
			rowsProcessed++
			// Take some time between calls
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Printf("Import complete.\nRows processed: %d\nUsers created: %d\nRows failed: %d\n", rowsProcessed, usersCreated, usersFailed)
}