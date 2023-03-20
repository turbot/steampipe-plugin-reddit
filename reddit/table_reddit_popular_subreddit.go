package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRedditPopularSubreddit(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_popular_subreddit",
		Description: "Popular subreddits.",
		List: &plugin.ListConfig{
			Hydrate: listPopularSubreddit,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the subreddit among the result rows, use for sorting."},
			{Name: "display_name_prefixed", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.NamePrefixed"), Description: "Prefixed name of the subreddit, e.g. /r/aws."},
			{Name: "subscribers", Type: proto.ColumnType_INT, Transform: transform.FromField("Subreddit.Subscribers"), Description: "Number of subscribers to the subreddit."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.URL"), Description: "URL of the subreddit."},
			// Other columns
			{Name: "active_user_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Subreddit.ActiveUserCount"), Description: "Active user count for the subreddit."},
			{Name: "created_utc", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Subreddit.Created").Transform(timeToRfc3339), Description: "Time when the subreddit was created."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Name"), Description: "Name of the subreddit, e.g. aws."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.ID"), Description: "ID of the subreddit."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.FullID"), Description: "Name (full ID) of the subreddit."},
			{Name: "over18", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.NSFW"), Description: "True if the post is not safe for work (over 18)."},
			{Name: "public_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Description"), Description: "Public description of the subreddit."},
			{Name: "subreddit_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Type"), Description: "Type of subreddit."},
			{Name: "suggested_comment_sort", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.SuggestedCommentSort"), Description: "Suggested sort order for comments."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Title"), Description: "Title of the subreddit."},
			{Name: "user_has_favorited", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.Favorite"), Description: "True if the caller has favorited the subreddit."},
			{Name: "user_is_moderator", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.UserIsMod"), Description: "True if the caller is a moderator of the subreddit."},
			{Name: "user_is_subscriber", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.Subscribed"), Description: "True if the caller is a subscriber of the subreddit."},
		},
	}
}

type popularSubredditRow struct {
	Subreddit *reddit.Subreddit
	// Return rank to sort results by popularity
	Rank int `json:"rank"`
}

func listPopularSubreddit(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_subreddit_subscribed.listPopularSubreddit", "connection_error", err)
		return nil, err
	}

	opts := &reddit.ListSubredditOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
	}

	// TODO[NW] - As of v0.14.3 the limit seems to always be nil? It does work in the stripe plugin, so I'm unsure why.
	plugin.Logger(ctx).Debug("reddit_subreddit_subscribed.listPopularSubreddit", "d.QueryContext", d.QueryContext)
	limit := d.QueryContext.Limit
	if limit != nil {
		iLimit := int(*limit)
		if iLimit < opts.ListOptions.Limit {
			opts.ListOptions.Limit = iLimit
		}
	}

	count := 0
	for {
		plugin.Logger(ctx).Debug("reddit_subreddit_subscribed.listPopularSubreddit", "opts", opts, "limit", limit)
		items, resp, err := conn.Subreddit.Popular(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_subreddit_subscribed.listPopularSubreddit", "query_error", err, "resp", resp, "opts", opts)
			return nil, err
		}
		for _, i := range items {
			count++
			row := popularSubredditRow{
				i,
				count,
			}
			d.StreamListItem(ctx, row)
		}
		// Stop if we've reached the end, or the max target limit
		// TODO - make 500 a config option?
		if resp.After == "" || count >= 500 {
			break
		}
		if limit != nil {
			if int64(count) >= *limit {
				break
			}
		}
		// Set the page and continue
		opts.After = resp.After
	}

	return nil, nil
}
