package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRedditMyBlocked(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_my_blocked",
		Description: "Users blocked by the user account making the call.",
		List: &plugin.ListConfig{
			Hydrate: listMyBlocked,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rel_id", Type: proto.ColumnType_STRING, Description: "ID of the relationship."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the blocked user."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the blocked user."},
			{Name: "date", Type: proto.ColumnType_STRING, Transform: transform.FromField("Created").Transform(timeToRfc3339), Description: "Time when the block was created."},
		},
	}
}

func listMyBlocked(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_blocked.listMyBlocked", "connection_error", err)
		return nil, err
	}

	items, resp, err := conn.Account.Blocked(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_blocked.listMyBlocked", "query_error", err, "resp", resp)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
