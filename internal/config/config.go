package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/warrant-dev/warrant-go/v3"
)

func InitClient() error {
	apiKey := viper.GetString("key")
	if apiKey == "" {
		return fmt.Errorf("missing API key. Please run `warrant init`")
	}
	warrant.ApiKey = apiKey
	return nil
}
