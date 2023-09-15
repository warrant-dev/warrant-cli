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
