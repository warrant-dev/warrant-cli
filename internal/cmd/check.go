package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/warrant-dev/warrant-cli/internal/config"
	"github.com/warrant-dev/warrant-cli/internal/reader"
	"github.com/warrant-dev/warrant-go"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check subject relation object",
	Short: "Check if a given subject-relation-object warrant exists",
	Long:  "Check a warrant for a user",
	Example: `
warrant check user:23 member role:admin
warrant check role:admin editor document:45`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := config.GetClient()
		if err != nil {
			return err
		}
		subject, err := reader.ParseObject(args[0])
		if err != nil {
			return err
		}
		relation := args[1]
		object, err := reader.ParseObject(args[2])
		if err != nil {
			return err
		}
		warrantToCheck := warrant.Warrant{
			ObjectType: object.Type,
			ObjectId:   object.Id,
			Relation:   relation,
			Subject: warrant.Subject{
				ObjectType: subject.Type,
				ObjectId:   subject.Id,
			},
		}
		result, err := client.IsAuthorized(warrant.WarrantCheckParams{
			Warrants: []warrant.Warrant{warrantToCheck},
		})
		if err != nil {
			return err
		}
		fmt.Printf("%t\n", result)
		return nil
	},
}
