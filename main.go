package main

import (
	"github.com/turbot/steampipe-plugin-reddit/reddit"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: reddit.Plugin})
}
