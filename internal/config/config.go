package config

import (
	"github.com/spf13/viper"
	"github.com/warrant-dev/warrant-go/v3"
)

func InitClient() error {
	warrant.ApiKey = viper.GetString("key")
	endpoint := viper.GetString("apiEndpoint")
	warrant.ApiEndpoint = endpoint
	warrant.AuthorizeEndpoint = endpoint
	return nil
}
