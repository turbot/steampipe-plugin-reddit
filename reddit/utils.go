package reddit

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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

	var accessToken string
	if redditConfig.AccessToken != nil {
		accessToken = *redditConfig.AccessToken
	}

	var credentials reddit.Credentials
	if accessToken != "" {
		credentials = reddit.Credentials{AccessToken: accessToken}
	} else {
		if clientID == "" || clientSecret == "" || username == "" || password == "" {
			// Credentials not set
			return nil, errors.New("client_id, client_secret, username and password must be configured")
		}

		credentials = reddit.Credentials{ID: clientID, Secret: clientSecret, Username: username, Password: password}
	}

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

var getRedditAuthenticatedUserMemoize = plugin.HydrateFunc(getRedditAuthenticatedUserUncached).Memoize(memoize.WithCacheKeyFunction(getRedditAuthenticatedUserCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getRedditAuthenticatedUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getRedditAuthenticatedUserMemoize(ctx, d, h)
}

// Build a cache key for the call to getOrganizationIdCacheKey.
func getRedditAuthenticatedUserCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "reddit.authenticated_user"
	return key, nil
}

func getRedditAuthenticatedUserUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getRedditAuthenticatedUser", "connection_error", err)
		return nil, err
	}
	user, resp, err := conn.Account.Info(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {

		return user.Name, nil
	}

	return nil, nil
}
