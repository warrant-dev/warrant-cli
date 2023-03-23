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
	Key         string `json:"key"`
	ApiEndpoint string `json:"apiEndpoint"`
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI for use",
	Long:  "Initialize the CLI for use, including configuring server endpoint and API key.",
	Example: `
warrant init`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Warrant endpoint override (leave blank to use https://api.warrant.dev default):")
		fmt.Print("> ")
		buf := bufio.NewReader(os.Stdin)
		input, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		endpoint := strings.TrimSpace(string(input))
		if endpoint == "" {
			endpoint = "https://api.warrant.dev"
		}
		fmt.Println("API Key:")
		fmt.Print("> ")
		buf = bufio.NewReader(os.Stdin)
		input, err = buf.ReadBytes('\n')
		if err != nil {
			return err
		}
		key := strings.TrimSpace(string(input))

		fmt.Println("Creating ~/.warrant.json")
		config := ConfigFile{
			ApiEndpoint: endpoint,
			Key:         key,
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
