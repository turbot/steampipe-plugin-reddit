package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableRedditMyModeratedSubreddit(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_my_moderated_subreddit",
		Description: "Subreddits you are a moderator of.",
		List: &plugin.ListConfig{
			Hydrate: listMyModeratedSubreddit,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "display_name_prefixed", Type: proto.ColumnType_STRING, Description: "Prefixed name of the subreddit, e.g. /r/aws."},
			{Name: "subscribers", Type: proto.ColumnType_INT, Transform: transform.FromField("Subscribers"), Description: "Number of subscribers to the subreddit."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the subreddit."},
			// Other columns
			{Name: "active_user_count", Type: proto.ColumnType_INT, Description: "Active user count for the subreddit."},
			{Name: "created_utc", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(timeToRfc3339), Description: "Time when the subreddit was created."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Name of the subreddit, e.g. aws."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the subreddit."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name (full ID) of the subreddit."},
			{Name: "over18", Type: proto.ColumnType_BOOL, Description: "True if the post is not safe for work (over 18)."},
			{Name: "public_description", Type: proto.ColumnType_STRING, Description: "Public description of the subreddit."},
			{Name: "subreddit_type", Type: proto.ColumnType_STRING, Description: "Type of subreddit."},
			{Name: "suggested_comment_sort", Type: proto.ColumnType_STRING, Description: "Suggested sort order for comments."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the subreddit."},
			{Name: "user_has_favorited", Type: proto.ColumnType_BOOL, Description: "True if the caller has favorited the subreddit."},
			{Name: "user_is_moderator", Type: proto.ColumnType_BOOL, Description: "True if the caller is a moderator of the subreddit."},
			{Name: "user_is_subscriber", Type: proto.ColumnType_BOOL, Description: "True if the caller is a subscriber of the subreddit."},
		},
	}
}

func listMyModeratedSubreddit(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_moderated_subreddit.listMyModeratedSubreddit", "connection_error", err)
		return nil, err
	}

	opts := &reddit.ListSubredditOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
	}

	for {
		items, resp, err := conn.Subreddit.Moderated(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_my_moderated_subreddit.listMyModeratedSubreddit", "query_error", err, "resp", resp, "opts", opts)
			return nil, err
		}
		for _, i := range items {
			d.StreamListItem(ctx, i)
		}
		if resp.After == "" {
			break
		}
		opts.After = resp.After
	}

	return nil, nil
}
