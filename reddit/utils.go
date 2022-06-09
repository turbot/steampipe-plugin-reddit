package reddit

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*reddit.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "reddit"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*reddit.Client), nil
	}

	// Default to using env vars
	clientID := os.Getenv("REDDIT_CLIENT_ID")
	clientSecret := os.Getenv("REDDIT_CLIENT_SECRET")
	username := os.Getenv("REDDIT_USERNAME")
	password := os.Getenv("REDDIT_PASSWORD")

	// But prefer the config
	redditConfig := GetConfig(d.Connection)
	if redditConfig.ClientID != nil {
		clientID = *redditConfig.ClientID
	}
	if redditConfig.ClientSecret != nil {
		clientSecret = *redditConfig.ClientSecret
	}
	if redditConfig.Username != nil {
		username = *redditConfig.Username
	}
	if redditConfig.Password != nil {
		password = *redditConfig.Password
	}

	if clientID == "" || clientSecret == "" || username == "" || password == "" {
		// Credentials not set
		return nil, errors.New("client_id, client_secret, username and password must be configured")
	}

	credentials := reddit.Credentials{ID: clientID, Secret: clientSecret, Username: username, Password: password}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Resource not found")
}

func timeToRfc3339(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value != nil {
		t := d.Value.(*reddit.Timestamp)
		if t != nil && !t.IsZero() {
			return t.Format(time.RFC3339), nil
		}
	}
	return nil, nil
}
