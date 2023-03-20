package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRedditMyPost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "reddit_my_post",
		Description: "Your posts.",
		List: &plugin.ListConfig{
			Hydrate: listMyPost,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the post among the result rows, use for sorting."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.ID"), Description: "ID of the post."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.FullID"), Description: "Slug (full ID) of the post."},
			{Name: "created_utc", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Post.Created").Transform(timeToRfc3339), Description: "Time when the post was created."},
			{Name: "edited", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Post.Edited").Transform(timeToRfc3339), Description: "Time when the post was edited."},
			{Name: "permalink", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Permalink"), Description: "Permalink (path only) to the post."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.URL"), Description: "URL the post links to, or of the post itself."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Title"), Description: "Title of the post."},
			{Name: "selftext", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Body"), Description: "Body of the post."},
			{Name: "likes", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Likes"), Description: "True if you've upvoted the post. False if you've downvoted it. Otherwise null."},
			{Name: "score", Type: proto.ColumnType_INT, Transform: transform.FromField("Post.Score"), Description: "Score of the post."},
			{Name: "upvote_ratio", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Post.UpvoteRatio"), Description: "Upvote ratio of the post."},
			{Name: "num_comments", Type: proto.ColumnType_INT, Transform: transform.FromField("Post.NumberOfComments"), Description: "Number of comments on the post."},
			{Name: "subreddit", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.SubredditName"), Description: "Name of the subreddit, e.g. aws."},
			{Name: "subreddit_name_prefixed", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.SubredditNamePrefixed"), Description: "Prefixed name of the subreddit, e.g. /r/aws."},
			{Name: "subreddit_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.SubredditID"), Description: "ID of the subreddit."},
			{Name: "subreddit_subscribers", Type: proto.ColumnType_INT, Transform: transform.FromField("Post.SubredditSubscribers"), Description: "Number of subscribers to the subreddit."},
			{Name: "author", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.Author"), Description: "Author of the post."},
			{Name: "author_fullname", Type: proto.ColumnType_STRING, Transform: transform.FromField("Post.AuthorID"), Description: "Full name of the author for the post."},
			{Name: "spoiler", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Spoiler"), Description: "True if the post is a spoiler."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Locked"), Description: "True if the post is locked."},
			{Name: "over_18", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.NSFW"), Description: "True if the post is not safe for work (over 18)."},
			{Name: "is_self", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.IsSelfPost"), Description: ""},
			{Name: "saved", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Saved"), Description: "True if the post has been saved."},
			{Name: "stickied", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Post.Stickied"), Description: "True if the post has been stickied."},
		},
	}
}

type myPostRow struct {
	Post *reddit.Post
	// Return rank to sort results
	Rank int `json:"rank"`
}

func listMyPost(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("reddit_my_post.listMyPost", "connection_error", err)
		return nil, err
	}

	authUserCached := plugin.HydrateFunc(getRedditAuthenticatedUser).WithCache()
	commonData, _ := authUserCached(ctx, d, h)
	authUser := commonData.(string)

	if authUser != "" {
		conn.Username = authUser
	}

	opts := &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		// One of: hot, new, top, controversial.
		Sort: "top",
		// One of: hour, day, week, month, year, all.
		Time: "all",
	}

	count := 0
	for {
		items, resp, err := conn.User.Posts(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("reddit_my_post.listMyPost", "query_error", err, "resp", resp, "opts", opts)
			return nil, err
		}
		for _, i := range items {
			count++
			row := myPostRow{
				i,
				count,
			}
			d.StreamListItem(ctx, row)
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
