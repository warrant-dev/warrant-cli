package cmd

import (
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.warrant.json)")
	// Declare and bind API key via config
	rootCmd.PersistentFlags().StringP("key", "k", "", "Warrant API key")
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".warrant.json"
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".warrant")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// Read config file
	err := viper.ReadInConfig()
	cobra.CheckErr(err)
}
