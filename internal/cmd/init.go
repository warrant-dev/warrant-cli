package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

type ConfigFile struct {
	Key string `json:"key"`
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI for use",
	Long:  "Initialize the CLI for use, including configuring the Warrant API key.",
	Example: `
warrant init`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Please navigate to https://app.warrant.dev/account in your browser, login and retrieve your API key (prod or test) and enter it here:")
		fmt.Print("> ")
		buf := bufio.NewReader(os.Stdin)
		input, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		fmt.Println("Creating ~/.warrant.json")
		key := strings.TrimSpace(string(input))
		config := ConfigFile{
			Key: key,
		}
		fileContents, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(homeDir+"/.warrant.json", fileContents, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Setup complete.")
		return nil
	},
}
