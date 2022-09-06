package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableRedditSubredditModerator(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "reddit_subreddit_moderator",
		Description:      "Moderators of the subreddit.",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		List: &plugin.ListConfig{
			Hydrate:    listSubredditModerator,
			KeyColumns: plugin.SingleColumn("subreddit"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Relationship.User"), Description: "Name of the moderator user."},
			{Name: "mod_permissions", Type: proto.ColumnType_JSON, Description: "Moderation permissions granted to the user."},
			// Other columns
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Relationship.Created").Transform(timeToRfc3339), Description: "Time when the subreddit was created."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Relationship.UserID"), Description: "ID of the moderator user."},
			{Name: "rel_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Relationship.ID"), Description: "ID of the ban."},
			{Name: "subreddit", Type: proto.ColumnType_STRING, Transform: transform.FromQual("subreddit"), Description: "Subreddit for this moderator."},
		},
	}
}

func listSubredditModerator(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_subreddit_moderator.listSubredditModerator", "connection_error", err)
		return nil, err
	}

	subreddit := d.KeyColumnQuals["subreddit"].GetStringValue()

	items, resp, err := conn.Subreddit.Moderators(ctx, subreddit)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_subreddit_moderator.listSubredditModerator", "query_error", err, "resp", resp)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
