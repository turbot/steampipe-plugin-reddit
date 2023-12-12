package reddit

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type redditConfig struct {
	ClientID     *string `hcl:"client_id"`
	ClientSecret *string `hcl:"client_secret"`
	Username     *string `hcl:"username"`
	Password     *string `hcl:"password"`
	AccessToken  *string `hcl:"access_token"`
}

func ConfigInstance() interface{} {
	return &redditConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) redditConfig {
	if connection == nil || connection.Config == nil {
		return redditConfig{}
	}
	config, _ := connection.Config.(redditConfig)
	return config
}
