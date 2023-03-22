package reddit

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type redditConfig struct {
	ClientID     *string `cty:"client_id"`
	ClientSecret *string `cty:"client_secret"`
	Username     *string `cty:"username"`
	Password     *string `cty:"password"`
	AccessToken  *string `cty:"access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"client_id": {
		Type: schema.TypeString,
	},
	"client_secret": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"access_token": {
		Type: schema.TypeString,
	},
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
