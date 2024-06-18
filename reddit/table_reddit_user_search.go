package reddit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func tableRedditUserSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_user_search",
		Description: "Search Reddit users.",
		// Allow for 0 counts
		DefaultTransform: transform.FromJSONTag(),
		List: &plugin.ListConfig{
			Hydrate: listUserSearch,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", CacheMatch: "exact"},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the user among the result rows, use for sorting."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Name"), Description: "Name of the user."},
			{Name: "link_karma", Type: proto.ColumnType_INT, Transform: transform.FromField("User.PostKarma"), Description: "Karma from links."},
			{Name: "comment_karma", Type: proto.ColumnType_INT, Transform: transform.FromField("User.CommentKarma"), Description: "Karma from comments."},
			// Other columns
			{Name: "created_utc", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Created").Transform(timeToRfc3339), Description: "Time when the user was created."},
			{Name: "has_verified_email", Type: proto.ColumnType_BOOL, Transform: transform.FromField("User.HasVerifiedEmail"), Description: "True if the user email has been verified."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "ID of the user."},
			{Name: "is_employee", Type: proto.ColumnType_BOOL, Transform: transform.FromField("User.IsEmployee"), Description: "True if the user is an employee."},
			{Name: "is_friend", Type: proto.ColumnType_BOOL, Transform: transform.FromField("User.IsFriend"), Description: "True if the user is a friend."},
			{Name: "is_suspended", Type: proto.ColumnType_BOOL, Transform: transform.FromField("User.IsSuspended"), Description: "True if the user has been suspended."},
			{Name: "over_18", Type: proto.ColumnType_BOOL, Transform: transform.FromField("User.NSFW"), Description: "True if the user is over 18."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Search query string."},
		}),
	}
}

type userSearchRow struct {
	User *reddit.User
	// Return rank to sort results
	Rank int `json:"rank"`
}

func listUserSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_post.listUserSearch", "connection_error", err)
		return nil, err
	}

	query := d.EqualsQuals["query"].GetStringValue()

	opts := &reddit.ListOptions{
		Limit: 100,
	}

	count := 0
	for {
		users, resp, err := conn.User.Search(ctx, query, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_my_post.listUserSearch", "query_error", err, "resp", resp, "query", query)
			return nil, err
		}
		for _, i := range users {
			count++
			row := userSearchRow{
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
