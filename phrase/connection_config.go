package phrase

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type phraseConfig struct {
	AccessToken *string `cty:"access_token"`
	Datacenter  *string `cty:"datacenter"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"access_token": {
		Type: schema.TypeString,
	},
	"datacenter": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &phraseConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) phraseConfig {
	if connection == nil || connection.Config == nil {
		return phraseConfig{}
	}
	config, _ := connection.Config.(phraseConfig)
	return config
}
