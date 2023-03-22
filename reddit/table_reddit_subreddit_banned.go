package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRedditSubredditBanned(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_subreddit_banned",
		Description: "Users banned from the subreddit.",
		List: &plugin.ListConfig{
			Hydrate:    listSubredditBanned,
			KeyColumns: plugin.SingleColumn("subreddit"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Relationship.User"), Description: "Name of the banned user."},
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Relationship.Created").Transform(timeToRfc3339), Description: "Time when the subreddit was created."},
			{Name: "days_left", Type: proto.ColumnType_INT, Description: "Days left in the ban."},
			{Name: "note", Type: proto.ColumnType_INT, Description: "Note for the ban."},
			// Other columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Relationship.UserID"), Description: "ID of the banned user."},
			{Name: "rel_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Relationship.ID"), Description: "ID of the ban."},
			{Name: "subreddit", Type: proto.ColumnType_STRING, Transform: transform.FromQual("subreddit"), Description: "Subreddit for the ban."},
		},
	}
}

func listSubredditBanned(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_subreddit_banned.listSubredditBanned", "connection_error", err)
		return nil, err
	}

	subreddit := d.EqualsQuals["subreddit"].GetStringValue()

	opts := &reddit.ListOptions{
		Limit: 100,
	}

	for {
		items, resp, err := conn.Subreddit.Banned(ctx, subreddit, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_subreddit_banned.listSubredditBanned", "query_error", err, "resp", resp, "opts", opts)
			return nil, err
		}
		for _, i := range items {
			d.StreamListItem(ctx, i)
		}
		// Stop if we've reached the end, or the max target limit
		if resp.After == "" {
			break
		}
		// Set the page and continue
		opts.After = resp.After
	}

	return nil, nil
}
