package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func tableRedditSubredditSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_subreddit_search",
		Description: "Search subreddits.",
		// Allow for 0 counts
		DefaultTransform: transform.FromJSONTag(),
		List: &plugin.ListConfig{
			Hydrate: listSubredditSearch,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the subreddit among the result rows, use for sorting."},
			{Name: "display_name_prefixed", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.NamePrefixed"), Description: "Prefixed name of the subreddit, e.g. /r/aws."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Title"), Description: "Title of the subreddit."},
			{Name: "subscribers", Type: proto.ColumnType_INT, Transform: transform.FromField("Subreddit.Subscribers"), Description: "Number of subscribers to the subreddit."},
			// Other columns
			{Name: "active_user_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Subreddit.ActiveUserCount"), Description: "Active user count for the subreddit."},
			{Name: "created_utc", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Subreddit.Created").Transform(timeToRfc3339), Description: "Time when the subreddit was created."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Name"), Description: "Name of the subreddit, e.g. aws."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.ID"), Description: "ID of the subreddit."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.FullID"), Description: "Name (full ID) of the subreddit."},
			{Name: "over18", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.NSFW"), Description: "True if the post is not safe for work (over 18)."},
			{Name: "public_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Description"), Description: "Public description of the subreddit."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Search query string."},
			{Name: "subreddit_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.Type"), Description: "Type of subreddit."},
			{Name: "suggested_comment_sort", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.SuggestedCommentSort"), Description: "Suggested sort order for comments."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subreddit.URL"), Description: "URL of the subreddit."},
			{Name: "user_has_favorited", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.Favorite"), Description: "True if the caller has favorited the subreddit."},
			{Name: "user_is_moderator", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.UserIsMod"), Description: "True if the caller is a moderator of the subreddit."},
			{Name: "user_is_subscriber", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Subreddit.Subscribed"), Description: "True if the caller is a subscriber of the subreddit."},
		},
	}
}

type subredditSearchRow struct {
	Subreddit *reddit.Subreddit
	// Return rank to sort results
	Rank int `json:"rank"`
}

func listSubredditSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_post.listSubredditSearch", "connection_error", err)
		return nil, err
	}

	query := d.EqualsQuals["query"].GetStringValue()

	opts := &reddit.ListSubredditOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Sort: "relevance",
	}

	count := 0
	for {
		subreddits, resp, err := conn.Subreddit.Search(ctx, query, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_my_post.listSubredditSearch", "query_error", err, "resp", resp, "query", query)
			return nil, err
		}
		for _, i := range subreddits {
			count++
			row := subredditSearchRow{
				i,
				count,
			}
			d.StreamListItem(ctx, row)
		}
		if resp.After == "" {
			break
		}
		opts.After = resp.After
	}

	return nil, nil
}
