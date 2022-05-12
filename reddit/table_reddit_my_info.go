package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableRedditMyInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_my_info",
		Description: "Information about the user making the query.",
		List: &plugin.ListConfig{
			Hydrate: listMyInfo,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user."},
			{Name: "created_utc", Type: proto.ColumnType_STRING, Transform: transform.FromField("Created").Transform(timeToRfc3339), Description: "Time when the user was created."},
			{Name: "link_karma", Type: proto.ColumnType_INT, Description: "Karma from links."},
			{Name: "comment_karma", Type: proto.ColumnType_INT, Description: "Karma from comments."},
			{Name: "is_friend", Type: proto.ColumnType_BOOL, Description: "True if the user is a friend."},
			{Name: "is_employee", Type: proto.ColumnType_BOOL, Description: "True if the user is an employee."},
			{Name: "has_verified_email", Type: proto.ColumnType_BOOL, Description: "True if the user email has been verified."},
			{Name: "over_18", Type: proto.ColumnType_BOOL, Description: "True if the user is over 18."},
			{Name: "is_suspended", Type: proto.ColumnType_BOOL, Description: "True if the user has been suspended."},
		},
	}
}

func listMyInfo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_post.listMyInfo", "connection_error", err)
		return nil, err
	}

	account, resp, err := conn.Account.Info(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_post.listMyInfo", "query_error", err, "resp", resp)
		return nil, err
	}
	d.StreamListItem(ctx, account)

	return nil, nil
}
