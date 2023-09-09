package config

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigFileName = ".warrant.json"

type Config struct {
	ActiveEnvironment string                 `mapstructure:"activeEnvironment" json:"activeEnvironment"`
	Environments      map[string]Environment `mapstructure:"environments" json:"environments"`
}

type Environment struct {
	ApiKey      string `mapstructure:"apiKey" json:"apiKey"`
	ApiEndpoint string `mapstructure:"apiEndpoint" json:"apiEndpoint"`
}

func (c Config) Write() error {
	fileContents, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = os.WriteFile(homeDir+"/"+ConfigFileName, fileContents, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() *Config {
	// Look for .warrant.json in HOME dir and create an empty version if it doesn't exist
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	_, err = os.Stat(homeDir + "/" + ConfigFileName)
	if errors.Is(err, fs.ErrNotExist) {
		emptyJson := []byte("{}")
		err = os.WriteFile(homeDir+"/"+ConfigFileName, emptyJson, 0644)
		cobra.CheckErr(err)
	}

	// Load config from ~/.warrant.json
	viper.AddConfigPath(homeDir)
	viper.SetConfigType("json")
	viper.SetConfigName(".warrant")
	viper.AutomaticEnv() // read in environment variables that match
	err = viper.ReadInConfig()
	cobra.CheckErr(err)

	// Unmarshal config & set warrant client vals
	var config Config
	err = viper.Unmarshal(&config)
	cobra.CheckErr(err)
	return &config
}
