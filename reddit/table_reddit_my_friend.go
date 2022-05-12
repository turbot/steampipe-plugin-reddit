package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableRedditMyFriend(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_my_friend",
		Description: "Friends of the user account making the call.",
		List: &plugin.ListConfig{
			Hydrate: listMyFriend,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rel_id", Type: proto.ColumnType_STRING, Description: "ID of the relationship."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the friend."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the friend."},
			{Name: "date", Type: proto.ColumnType_STRING, Transform: transform.FromField("Created").Transform(timeToRfc3339), Description: "Time when the relationship was created."},
		},
	}
}

func listMyFriend(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_friend.listMyFriend", "connection_error", err)
		return nil, err
	}

	items, resp, err := conn.Account.Friends(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_friend.listMyFriend", "query_error", err, "resp", resp)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
