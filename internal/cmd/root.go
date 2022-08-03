package cmd

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "warrant",
	Short: "Warrant CLI",
	Long:  `The Warrant CLI is a tool to interact with Warrant via the command line.`,
}

func Execute() {
	// Execute requested cmd and handle any errors
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Flags (including persistent) definition
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.warrant.json)")
	// Declare and bind API key via config
	// rootCmd.PersistentFlags().StringP("key", "k", "", "Warrant API key")
	// viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))

	// Cobra also supports local flags, which will only run when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// Look for .warrant.json in HOME dir and create an empty version if it doesn't exist
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	_, err = os.Stat(home + "/.warrant.json")
	if os.IsNotExist(err) {
		emptyJson := []byte("{}")
		err = ioutil.WriteFile(home+"/.warrant.json", emptyJson, 0644)
		if err != nil {
			cobra.CheckErr(err)
		}
	}

	// Load config from ~/.warrant.json
	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.SetConfigName(".warrant")
	viper.AutomaticEnv() // read in environment variables that match
	err = viper.ReadInConfig()
	cobra.CheckErr(err)
}
