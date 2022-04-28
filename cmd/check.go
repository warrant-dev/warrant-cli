package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warrant-dev/warrant-go"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check a warrant for a user",
	Long: `Check a warrant for a user
	
Example:
	warrant check userId1 owner document doc23`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("key")
		if apiKey == "" {
			fmt.Println("Missing API key. Please define it in your config file or using the --key flag.")
			return
		}
		client := warrant.NewClient(warrant.ClientConfig{
			ApiKey: apiKey,
		})
		result, err := client.IsAuthorized(warrant.Warrant{
			User: warrant.WarrantUser{
				UserId: args[0],
			},
			Relation:   args[1],
			ObjectType: args[2],
			ObjectId:   args[3],
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Result--> %t\n", result)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
