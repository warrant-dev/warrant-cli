package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/warrant-dev/warrant-go"
)

func GetClient() (warrant.WarrantClient, error) {
	apiKey := viper.GetString("key")
	if apiKey == "" {
		return warrant.WarrantClient{}, fmt.Errorf("Missing API key. Please define it in your config file or using the --key flag.")
	}
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: apiKey,
	})
	return client, nil
}
